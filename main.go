package main

import (
	"path/filepath"
	"fmt"
	"bytes"
	"encoding/json"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/client_golang/prometheus"
)

var errorCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: "mesos",
	Subsystem: "collector",
	Name:      "errors_total",
	Help:      "Total number of internal mesos-collector errors.",
})

func init() {
	prometheus.MustRegister(errorCounter)
}

func getX509CertPool(pemFiles []string) *x509.CertPool {
	pool := x509.NewCertPool()
	for _, f := range pemFiles {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}
		ok := pool.AppendCertsFromPEM(content)
		if !ok {
			log.Fatal("Error parsing .pem file %s", f)
		}
	}
	return pool
}

func mkHttpClient(url string, timeout time.Duration, auth authInfo, certPool *x509.CertPool) *httpClient {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: certPool, InsecureSkipVerify: auth.skipSSLVerify},
	}

	// HTTP Redirects are authenticated by Go (>=1.8), when redirecting to an identical domain or a subdomain.
	// -> Hijack redirect authentication, since hostnames rarely follow this logic.
	var redirectFunc func(req *http.Request, via []*http.Request) error
	if auth.username != "" && auth.password != "" {
		// Auth information is only available in the current context -> use lambda function
		redirectFunc = func(req *http.Request, via []*http.Request) error {
			req.SetBasicAuth(auth.username, auth.password)
			return nil
		}
	}

	client := &httpClient{
		http.Client{Timeout: timeout, Transport: transport, CheckRedirect: redirectFunc},
		url,
		auth,
	}

	if auth.strictMode {
		client.auth.token = authToken(client)
	}

	return client
}

func parsePrivateKey(httpClient *httpClient) []byte {
	if _, err := os.Stat(httpClient.auth.privateKey); os.IsNotExist(err) {
		buffer := bytes.NewBuffer([]byte(httpClient.auth.privateKey))
		var key mesosSecret
		if err := json.NewDecoder(buffer).Decode(&key); err != nil {
			log.Printf("Error decoding prviate key %s: %s", key, err)
			errorCounter.Inc()
			return []byte{}
		}
		httpClient.auth.username = key.UID
		httpClient.auth.loginUrl = key.LoginEndpoint
		return []byte(key.PrivateKey)
	} else {
		absPath, _ := filepath.Abs(httpClient.auth.privateKey)
		key, err := ioutil.ReadFile(absPath)
		if err != nil {
			log.Printf("Error reading private key %s: %s", absPath, err)
			errorCounter.Inc()
			return []byte{}
		}
		return key
	}
}

func loginAuthToken(httpClient *httpClient) string {
	key := parsePrivateKey(httpClient)
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Printf("Error parsing privateKey: %s", err)
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"uid": httpClient.auth.username,
	})
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Printf("Error creating login token %s: %s", token, err)
		return ""
	}
	return tokenString
}

func authToken(httpClient *httpClient) string {
	url := httpClient.auth.loginUrl
	loginToken := loginAuthToken(httpClient)
	body, err := json.Marshal(&tokenRequest{UID: httpClient.auth.username, Token: loginToken})
	if err != nil {
		log.Printf("Error creating JSON request: %s", err)
		return ""
	}
	buffer := bytes.NewBuffer(body)
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		log.Printf("Error creating HTTP request to %s: %s", url, err)
		return ""
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Error fetching %s: %s", url, err)
		errorCounter.Inc()
		return ""
	}
	defer res.Body.Close()

	var token tokenResponse
	if err := json.NewDecoder(res.Body).Decode(&token); err != nil {
		log.Printf("Error decoding response body from %s: %s", url, err)
		errorCounter.Inc()
		return ""
	}

	return fmt.Sprintf("token=%s", token.Token)
}

func csvInputToList(input string) []string {
	var entryList []string
	if input == "" {
		return entryList
	}
	sanitizedString := strings.Replace(input, " ", "", -1)
	entryList = strings.Split(sanitizedString, ",")
	return entryList
}

func main() {
	fs := flag.NewFlagSet("mesos-exporter", flag.ExitOnError)
	addr := fs.String("addr", ":9110", "Address to listen on")
	masterURL := fs.String("master", "", "Expose metrics from master running on this URL")
	slaveURL := fs.String("slave", "", "Expose metrics from slave running on this URL")
	timeout := fs.Duration("timeout", 5*time.Second, "Master polling timeout")
	exportedTaskLabels := fs.String("exportedTaskLabels", "", "Comma-separated list of task labels to include in the corresponding metric")
	exportedSlaveAttributes := fs.String("exportedSlaveAttributes", "", "Comma-separated list of slave attributes to include in the corresponding metric")
	ignoreCompletedFrameworkTasks := fs.Bool("ignoreCompletedFrameworkTasks", false, "Don't export task_state_time metric")
	trustedCerts := fs.String("trustedCerts", "", "Comma-separated list of certificates (.pem files) trusted for requests to Mesos endpoints")
	strictMode := fs.Bool("strictMode", false, "Use strict mode authentication")
	username := fs.String("username", "", "Username for authentication")
	password := fs.String("password", "", "Password for authentication")
	loginUrl := fs.String("loginUrl", "https://leader.mesos/acs/api/v1/auth/login", "URLfor strict mode authentication")
	privateKey := fs.String("privateKey", "", "File path to certificate for strict mode authentication")
	skipSSLVerify := fs.Bool("skipSSLVerify", false, "Skip SSL certificate verification")

	fs.Parse(os.Args[1:])
	if *masterURL != "" && *slaveURL != "" {
		log.Fatal("Only -master or -slave can be given at a time")
	}

	auth := authInfo{
		strictMode: *strictMode,
		skipSSLVerify: *skipSSLVerify,
		loginUrl: *loginUrl,
	}

	if *strictMode && *privateKey != "" {
		auth.privateKey = *privateKey
	} else {
		auth.privateKey = os.Getenv("MESOS_EXPORTER_PRIVATE_KEY")
	}

	if *username != "" {
		auth.username = *username
	} else {
		auth.username = os.Getenv("MESOS_EXPORTER_USERNAME")
	}

	if *password != "" {
		auth.password = *password
	} else {
		auth.password = os.Getenv("MESOS_EXPORTER_PASSWORD")
	}

	var certPool *x509.CertPool = nil
	if *trustedCerts != "" {
		certPool = getX509CertPool(csvInputToList(*trustedCerts))
	}

	slaveAttributeLabels := csvInputToList(*exportedSlaveAttributes)
	slaveTaskLabels := csvInputToList(*exportedTaskLabels)

	switch {
	case *masterURL != "":
		for _, f := range []func(*httpClient) prometheus.Collector{
			newMasterCollector,
			func(c *httpClient) prometheus.Collector {
				return newMasterStateCollector(c, *ignoreCompletedFrameworkTasks, slaveAttributeLabels)
			},
		} {
			c := f(mkHttpClient(*masterURL, *timeout, auth, certPool))
			if err := prometheus.Register(c); err != nil {
				log.Fatal(err)
			}
		}
		log.Printf("Exposing master metrics on %s", *addr)

	case *slaveURL != "":
		slaveCollectors := []func(*httpClient) prometheus.Collector{
			func(c *httpClient) prometheus.Collector {
				return newSlaveCollector(c)
			},
			func(c *httpClient) prometheus.Collector {
				return newSlaveMonitorCollector(c)
			},
		}

		if len(slaveTaskLabels) > 0 || len(slaveAttributeLabels) > 0 {
			slaveCollectors = append(slaveCollectors, func(c *httpClient) prometheus.Collector {
				return newSlaveStateCollector(c, slaveTaskLabels, slaveAttributeLabels)
			})
		}

		for _, f := range slaveCollectors {
			c := f(mkHttpClient(*slaveURL, *timeout, auth, certPool))
			if err := prometheus.Register(c); err != nil {
				log.Fatal(err)
			}
		}
		log.Printf("Exposing slave metrics on %s", *addr)

	default:
		log.Fatal("Either -master or -slave is required")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
            <head><title>Mesos Exporter</title></head>
            <body>
            <h1>Mesos Exporter</h1>
            <p><a href="/metrics">Metrics</a></p>
            </body>
            </html>`))
	})
	http.Handle("/metrics", prometheus.Handler())
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal(err)
	}
}
