package main

import (
	"fmt"
	"regexp"

	"github.com/prometheus/client_golang/prometheus"
)

func newMasterCollector(httpClient *httpClient) prometheus.Collector {
	metrics := map[prometheus.Collector]func(metricMap, prometheus.Collector) error{
		// CPU/Disk/Mem resources in free/used
		gauge("master", "cpus", "Current CPU resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/cpus_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/cpus_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/cpus_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		gauge("master", "cpus_revocable", "Current revocable CPU resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/cpus_revocable_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/cpus_revocable_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/cpus_revocable_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		gauge("master", "gpus", "Current GPU resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/gpus_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/gpus_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/gpus_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		gauge("master", "gpus_revocable", "Current revocable GPU resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/gpus_revocable_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/gpus_revocable_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/gpus_revocable_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		gauge("master", "mem", "Current memory resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/mem_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/mem_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/mem_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		gauge("master", "mem_revocable", "Current revocable memory resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/mem_revocable_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/mem_revocable_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/mem_revocable_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		gauge("master", "disk", "Current disk resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/disk_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/disk_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/disk_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		gauge("master", "disk_revocable", "Current disk resources in cluster.", "type"): func(m metricMap, c prometheus.Collector) error {
			percent, ok := m["master/disk_revocable_percent"]
			if !ok {
				return errNotFoundInMap
			}
			total, ok := m["master/disk_revocable_total"]
			if !ok {
				return errNotFoundInMap
			}
			used, ok := m["master/disk_revocable_used"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("percent").Set(percent)
			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("free").Set(total - used)
			c.(*prometheus.GaugeVec).WithLabelValues("used").Set(used)
			return nil
		},
		// Master stats about uptime and election state
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "master",
			Name:      "elected",
			Help:      "1 if master is elected leader, 0 if not",
		}): func(m metricMap, c prometheus.Collector) error {
			elected, ok := m["master/elected"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(elected)
			return nil
		},
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "master",
			Name:      "uptime_seconds",
			Help:      "Number of seconds the master process is running.",
		}): func(m metricMap, c prometheus.Collector) error {
			uptime, ok := m["master/uptime_secs"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(uptime)
			return nil
		},
		// Master stats about agents
		counter("master", "slave_registration_events_total", "Total number of registration events on this master since it booted.", "event"): func(m metricMap, c prometheus.Collector) error {
			registrations, ok := m["master/slave_registrations"]
			if !ok {
				return errNotFoundInMap
			}
			reregistrations, ok := m["master/slave_reregistrations"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(registrations, "register")
			c.(*settableCounterVec).Set(reregistrations, "reregister")
			return nil
		},

		counter("master", "recovery_slave_removal_events_total", "Total number of recovery removal events on this master since it booted.", "event"): func(m metricMap, c prometheus.Collector) error {
			removals, ok := m["master/recovery_slave_removals"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(removals, "removal")
			return nil
		},

		counter("master", "slave_removal_events_total", "Total number of removal events on this master since it booted.", "event"): func(m metricMap, c prometheus.Collector) error {
			scheduled, ok := m["master/slave_shutdowns_scheduled"]
			if !ok {
				return errNotFoundInMap
			}
			canceled, ok := m["master/slave_shutdowns_canceled"]
			if !ok {
				return errNotFoundInMap
			}
			completed, ok := m["master/slave_shutdowns_completed"]
			if !ok {
				return errNotFoundInMap
			}
			removals, ok := m["master/slave_removals"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(scheduled, "scheduled")
			c.(*settableCounterVec).Set(canceled, "canceled")
			c.(*settableCounterVec).Set(completed, "completed")
			// set this explicitly to be more obvious
			died := removals - completed
			c.(*settableCounterVec).Set(died, "died")
			return nil
		},
		counter("master", "slave_removal_events_reasons", "Total number of slave removal events by reason on this master since it booted.", "reason"): func(m metricMap, c prometheus.Collector) error {
			re, err := regexp.Compile("master/slave_removals/reason_(.*?)$")
			if err != nil {
				return fmt.Errorf("could not compile slave_removals reasons regex: %s", err)
			}
			for metric, value := range m {
				matches := re.FindStringSubmatch(metric)
				if len(matches) != 2 {
					continue
				}
				reason := matches[1]
				c.(*settableCounterVec).Set(value, reason)
			}
			return nil
		},
		counter("master", "slave_unreachable_events_total", "Total number of slave unreachable events on this master since it booted.", "event"): func(m metricMap, c prometheus.Collector) error {
			canceled, ok := m["master/slave_unreachable_canceled"]
			if !ok {
				return errNotFoundInMap
			}
			completed, ok := m["master/slave_unreachable_completed"]
			if !ok {
				return errNotFoundInMap
			}
			scheduled, ok := m["master/slave_unreachable_scheduled"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(canceled, "canceled")
			c.(*settableCounterVec).Set(completed, "completed")
			c.(*settableCounterVec).Set(scheduled, "scheduled")
			return nil
		},

		gauge("master", "slaves_state", "Current number of slaves known to the master per connection and registration state.", "state"): func(m metricMap, c prometheus.Collector) error {
			active, ok := m["master/slaves_active"]
			if !ok {
				return errNotFoundInMap
			}
			inactive, ok := m["master/slaves_inactive"]
			if !ok {
				return errNotFoundInMap
			}
			disconnected, ok := m["master/slaves_disconnected"]
			if !ok {
				return errNotFoundInMap
			}
			unreachable, ok := m["master/slaves_unreachable"]
			if !ok {
				return errNotFoundInMap
			}

			// FIXME: Make sure those assumptions are right
			// Every "active" node is connected to the master
			c.(*prometheus.GaugeVec).WithLabelValues("connected_active").Set(active)
			// Every "inactive" node is connected but node sending offers
			c.(*prometheus.GaugeVec).WithLabelValues("connected_inactive").Set(inactive)
			// Every "disconnected" node is "inactive"
			c.(*prometheus.GaugeVec).WithLabelValues("disconnected_inactive").Set(disconnected)
			// Every "connected" node is either active or inactive
			c.(*prometheus.GaugeVec).WithLabelValues("unreachable").Set(unreachable)
			return nil
		},

		// Master stats about frameworks
		gauge("master", "frameworks_state", "Current number of frames known to the master per connection and registration state.", "state"): func(m metricMap, c prometheus.Collector) error {
			active, ok := m["master/frameworks_active"]
			if !ok {
				return errNotFoundInMap
			}
			inactive, ok := m["master/frameworks_inactive"]
			if !ok {
				return errNotFoundInMap
			}
			disconnected, ok := m["master/frameworks_disconnected"]
			if !ok {
				return errNotFoundInMap
			}
			// FIXME: Make sure those assumptions are right
			// Every "active" framework is connected to the master
			c.(*prometheus.GaugeVec).WithLabelValues("connected_active").Set(active)
			// Every "inactive" framework is connected but framework sending offers
			c.(*prometheus.GaugeVec).WithLabelValues("connected_inactive").Set(inactive)
			// Every "disconnected" framework is "inactive"
			c.(*prometheus.GaugeVec).WithLabelValues("disconnected_inactive").Set(disconnected)
			// Every "connected" framework is either active or inactive
			return nil
		},
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "master",
			Name:      "offers_pending",
			Help:      "Current number of offers made by the master which aren't yet accepted or declined by frameworks.",
		}): func(m metricMap, c prometheus.Collector) error {
			offers, ok := m["master/outstanding_offers"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(offers)
			return nil
		},

		// Master stats about tasks
		counter("master", "task_states_exit_total", "Total number of tasks processed by exit state.", "state"): func(m metricMap, c prometheus.Collector) error {
			dropped, ok := m["master/tasks_dropped"]
			if !ok {
				return errNotFoundInMap
			}
			errored, ok := m["master/tasks_error"]
			if !ok {
				return errNotFoundInMap
			}
			failed, ok := m["master/tasks_failed"]
			if !ok {
				return errNotFoundInMap
			}
			finished, ok := m["master/tasks_finished"]
			if !ok {
				return errNotFoundInMap
			}
			gone, ok := m["master/tasks_gone"]
			if !ok {
				return errNotFoundInMap
			}
			goneByOperator, ok := m["master/tasks_gone_by_operator"]
			if !ok {
				return errNotFoundInMap
			}
			killed, ok := m["master/tasks_killed"]
			if !ok {
				return errNotFoundInMap
			}
			killing, ok := m["master/tasks_killing"]
			if !ok {
				return errNotFoundInMap
			}
			lost, ok := m["master/tasks_lost"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(dropped, "dropped")
			c.(*settableCounterVec).Set(errored, "errored")
			c.(*settableCounterVec).Set(failed, "failed")
			c.(*settableCounterVec).Set(finished, "finished")
			c.(*settableCounterVec).Set(gone, "gone")
			c.(*settableCounterVec).Set(goneByOperator, "gone_by_operator")
			c.(*settableCounterVec).Set(killed, "killed")
			c.(*settableCounterVec).Set(killing, "killing")
			c.(*settableCounterVec).Set(lost, "lost")
			return nil
		},

		counter("master", "task_states_current", "Current number of tasks by state.", "state"): func(m metricMap, c prometheus.Collector) error {
			running, ok := m["master/tasks_running"]
			if !ok {
				return errNotFoundInMap
			}
			staging, ok := m["master/tasks_staging"]
			if !ok {
				return errNotFoundInMap
			}
			starting, ok := m["master/tasks_starting"]
			if !ok {
				return errNotFoundInMap
			}
			unreachable, ok := m["master/tasks_unreachable"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(running, "running")
			c.(*settableCounterVec).Set(staging, "staging")
			c.(*settableCounterVec).Set(starting, "starting")
			c.(*settableCounterVec).Set(unreachable, "unreachable")
			return nil
		},

		counter("master", "task_state_counts_by_source_reason", "Number of task states by source and reason", "state", "source", "reason"): func(m metricMap, c prometheus.Collector) error {
			re, err := regexp.Compile("master/task_(.*?)/source_(.*?)/reason_(.*?)$")
			if err != nil {
				return fmt.Errorf("could not compile task_state_counts regex: %s", err)
			}
			for metric, value := range m {
				matches := re.FindStringSubmatch(metric)
				if len(matches) != 4 {
					continue
				}
				state := matches[1]
				source := matches[2]
				reason := matches[3]
				c.(*settableCounterVec).Set(value, state, source, reason)
			}
			return nil
		},

		// Master stats about messages
		counter("master", "messages", "Number of messages by the master by state", "type"): func(m metricMap, c prometheus.Collector) error {
			droppedMessages, ok := m["master/dropped_messages"]
			if !ok {
				return errNotFoundInMap
			}
			authenticateMessages, ok := m["master/messages_authenticate"]
			if !ok {
				return errNotFoundInMap
			}
			deactivateFrameworkMessages, ok := m["master/messages_deactivate_framework"]
			if !ok {
				return errNotFoundInMap
			}
			declineOfferMessages, ok := m["master/messages_decline_offers"]
			if !ok {
				return errNotFoundInMap
			}
			executorToFrameworkMessages, ok := m["master/messages_executor_to_framework"]
			if !ok {
				return errNotFoundInMap
			}
			exitedExecutor, ok := m["master/messages_exited_executor"]
			if !ok {
				return errNotFoundInMap
			}
			frameworkToExecutor, ok := m["master/messages_framework_to_executor"]
			if !ok {
				return errNotFoundInMap
			}
			killTask, ok := m["master/messages_kill_task"]
			if !ok {
				return errNotFoundInMap
			}
			launchTasks, ok := m["master/messages_launch_tasks"]
			if !ok {
				return errNotFoundInMap
			}
			reconcileTasks, ok := m["master/messages_reconcile_tasks"]
			if !ok {
				return errNotFoundInMap
			}
			registerFramework, ok := m["master/messages_register_framework"]
			if !ok {
				return errNotFoundInMap
			}
			registerSlave, ok := m["master/messages_register_slave"]
			if !ok {
				return errNotFoundInMap
			}
			reregisterFramework, ok := m["master/messages_reregister_framework"]
			if !ok {
				return errNotFoundInMap
			}
			reregisterSlave, ok := m["master/messages_reregister_slave"]
			if !ok {
				return errNotFoundInMap
			}
			resourceRequest, ok := m["master/messages_resource_request"]
			if !ok {
				return errNotFoundInMap
			}
			reviveOffers, ok := m["master/messages_revive_offers"]
			if !ok {
				return errNotFoundInMap
			}
			statusUpdate, ok := m["master/messages_status_update"]
			if !ok {
				return errNotFoundInMap
			}
			statusUpdateAck, ok := m["master/messages_status_update_acknowledgement"]
			if !ok {
				return errNotFoundInMap
			}
			suppressOffers, ok := m["master/messages_suppress_offers"]
			if !ok {
				return errNotFoundInMap
			}
			unregisterFramework, ok := m["master/messages_unregister_framework"]
			if !ok {
				return errNotFoundInMap
			}
			unregisterSlave, ok := m["master/messages_unregister_slave"]
			if !ok {
				return errNotFoundInMap
			}
			updateSlave, ok := m["master/messages_update_slave"]
			if !ok {
				return errNotFoundInMap
			}

			c.(*settableCounterVec).Set(authenticateMessages, "authenticate_messages")
			c.(*settableCounterVec).Set(droppedMessages, "dropped_messages")
			c.(*settableCounterVec).Set(deactivateFrameworkMessages, "deactivate_framework")
			c.(*settableCounterVec).Set(declineOfferMessages, "decline_offers")
			c.(*settableCounterVec).Set(executorToFrameworkMessages, "executor_to_framework")
			c.(*settableCounterVec).Set(exitedExecutor, "exited_executor")
			c.(*settableCounterVec).Set(frameworkToExecutor, "framework_to_executor")
			c.(*settableCounterVec).Set(killTask, "kill_task")
			c.(*settableCounterVec).Set(launchTasks, "launch_tasks")
			c.(*settableCounterVec).Set(reconcileTasks, "reconcile_tasks")
			c.(*settableCounterVec).Set(registerFramework, "register_framework")
			c.(*settableCounterVec).Set(registerSlave, "register_slave")
			c.(*settableCounterVec).Set(reregisterFramework, "reregister_framework")
			c.(*settableCounterVec).Set(reregisterSlave, "reregister_slave")
			c.(*settableCounterVec).Set(resourceRequest, "resource_request")
			c.(*settableCounterVec).Set(reviveOffers, "revive_offers")
			c.(*settableCounterVec).Set(statusUpdate, "status_update")
			c.(*settableCounterVec).Set(statusUpdateAck, "status_update_acknowledgement")
			c.(*settableCounterVec).Set(suppressOffers, "suppress_offers")
			c.(*settableCounterVec).Set(unregisterFramework, "unregister_framework")
			c.(*settableCounterVec).Set(unregisterSlave, "unregister_slave")
			c.(*settableCounterVec).Set(updateSlave, "update_slave")
			return nil
		},

		counter("master", "messages_outcomes_total",
			"Total number of messages by outcome of operation and direction.",
			"source", "destination", "type", "outcome"): func(m metricMap, c prometheus.Collector) error {
			frameworkToExecutorValid, ok := m["master/valid_framework_to_executor_messages"]
			if !ok {
				return errNotFoundInMap
			}
			frameworkToExecutorInvalid, ok := m["master/invalid_framework_to_executor_messages"]
			if !ok {
				return errNotFoundInMap
			}
			executorToFrameworkValid, ok := m["master/valid_executor_to_framework_messages"]
			if !ok {
				return errNotFoundInMap
			}
			executorToFrameworkInvalid, ok := m["master/invalid_executor_to_framework_messages"]
			if !ok {
				return errNotFoundInMap
			}

			// status updates are sent from framework?(FIXME) to slave
			// status update acks are sent from slave to framework?
			statusUpdateAckValid, ok := m["master/valid_status_update_acknowledgements"]
			if !ok {
				return errNotFoundInMap
			}
			statusUpdateAckInvalid, ok := m["master/invalid_status_update_acknowledgements"]
			if !ok {
				return errNotFoundInMap
			}
			statusUpdateValid, ok := m["master/valid_status_updates"]
			if !ok {
				return errNotFoundInMap
			}
			statusUpdateInvalid, ok := m["master/invalid_status_updates"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(frameworkToExecutorValid, "framework", "executor", "", "valid")
			c.(*settableCounterVec).Set(frameworkToExecutorInvalid, "framework", "executor", "", "invalid")

			c.(*settableCounterVec).Set(executorToFrameworkValid, "executor", "framework", "", "valid")
			c.(*settableCounterVec).Set(executorToFrameworkInvalid, "executor", "framework", "", "invalid")

			// We consider a ack message simply as a message from slave to framework
			c.(*settableCounterVec).Set(statusUpdateValid, "framework", "slave", "status_update", "valid")
			c.(*settableCounterVec).Set(statusUpdateInvalid, "framework", "slave", "status_update", "invalid")
			c.(*settableCounterVec).Set(statusUpdateAckValid, "slave", "framework", "status_update", "valid")
			c.(*settableCounterVec).Set(statusUpdateAckInvalid, "slave", "framework", "status_update", "invalid")
			return nil
		},
		// Master stats about events
		gauge("master", "event_queue_length", "Current number of elements in event queue by type", "type"): func(m metricMap, c prometheus.Collector) error {
			dispatches, ok := m["master/event_queue_dispatches"]
			if !ok {
				return errNotFoundInMap
			}
			httpRequests, ok := m["master/event_queue_http_requests"]
			if !ok {
				return errNotFoundInMap
			}
			messages, ok := m["master/event_queue_messages"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("message").Set(messages)
			c.(*prometheus.GaugeVec).WithLabelValues("http_request").Set(httpRequests)
			c.(*prometheus.GaugeVec).WithLabelValues("dispatches").Set(dispatches)
			return nil
		},

		// Master stats about allocations
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "master",
			Name:      "allocator_event_queue_dispatches",
			Help:      "Number of dispatch events in the allocator event queue.",
		}): func(m metricMap, c prometheus.Collector) error {
			count := m["allocator/event_queue_dispatches"]
			c.(prometheus.Gauge).Set(count)
			return nil
		},

		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "master",
			Name:      "allocation_run_ms_count",
			Help:      "Number of allocation algorithm time measurements in the window",
		}): func(m metricMap, c prometheus.Collector) error {
			count := m["allocator/mesos/allocation_runs"]
			c.(prometheus.Gauge).Set(count)
			return nil
		},

		gauge("master", "allocation_run_ms", "Time spent in allocation algorithm in ms.", "type"): func(m metricMap, c prometheus.Collector) error {
			mean, ok := m["allocator/mesos/allocation_run_ms"]
			if !ok {
				return errNotFoundInMap
			}
			min, ok := m["allocator/mesos/allocation_run_ms/min"]
			if !ok {
				return errNotFoundInMap
			}
			max, ok := m["allocator/mesos/allocation_run_ms/max"]
			if !ok {
				return errNotFoundInMap
			}
			p50, ok := m["allocator/mesos/allocation_run_ms/p50"]
			if !ok {
				return errNotFoundInMap
			}
			p90, ok := m["allocator/mesos/allocation_run_ms/p90"]
			if !ok {
				return errNotFoundInMap
			}
			p95, ok := m["allocator/mesos/allocation_run_ms/p95"]
			if !ok {
				return errNotFoundInMap
			}
			p99, ok := m["allocator/mesos/allocation_run_ms/p99"]
			if !ok {
				return errNotFoundInMap
			}
			p999, ok := m["allocator/mesos/allocation_run_ms/p999"]
			if !ok {
				return errNotFoundInMap
			}
			p9999, ok := m["allocator/mesos/allocation_run_ms/p9999"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("mean").Set(mean)
			c.(*prometheus.GaugeVec).WithLabelValues("min").Set(min)
			c.(*prometheus.GaugeVec).WithLabelValues("max").Set(max)
			c.(*prometheus.GaugeVec).WithLabelValues("p50").Set(p50)
			c.(*prometheus.GaugeVec).WithLabelValues("p90").Set(p90)
			c.(*prometheus.GaugeVec).WithLabelValues("p95").Set(p95)
			c.(*prometheus.GaugeVec).WithLabelValues("p99").Set(p99)
			c.(*prometheus.GaugeVec).WithLabelValues("p999").Set(p999)
			c.(*prometheus.GaugeVec).WithLabelValues("p9999").Set(p9999)
			return nil
		},

		counter("master", "allocation_runs", "Number of times the allocation alorithm has run", "event"): func(m metricMap, c prometheus.Collector) error {
			runs, ok := m["allocator/mesos/allocation_runs"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(runs, "allocation")
			return nil
		},

		counter("master", "allocation_run_latency_ms_count", "Number of allocation batch latency measurements", "event"): func(m metricMap, c prometheus.Collector) error {
			count, ok := m["allocator/mesos/allocation_run_latency_ms/count"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*settableCounterVec).Set(count, "allocation")
			return nil
		},

		gauge("master", "allocation_run_latency_ms", "Allocation batch latency in ms.", "type"): func(m metricMap, c prometheus.Collector) error {
			mean, ok := m["allocator/mesos/allocation_run_latency_ms"]
			if !ok {
				return errNotFoundInMap
			}
			min, ok := m["allocator/mesos/allocation_run_latency_ms/min"]
			if !ok {
				return errNotFoundInMap
			}
			max, ok := m["allocator/mesos/allocation_run_latency_ms/max"]
			if !ok {
				return errNotFoundInMap
			}
			p50, ok := m["allocator/mesos/allocation_run_latency_ms/p50"]
			if !ok {
				return errNotFoundInMap
			}
			p90, ok := m["allocator/mesos/allocation_run_latency_ms/p90"]
			if !ok {
				return errNotFoundInMap
			}
			p95, ok := m["allocator/mesos/allocation_run_latency_ms/p95"]
			if !ok {
				return errNotFoundInMap
			}
			p99, ok := m["allocator/mesos/allocation_run_latency_ms/p99"]
			if !ok {
				return errNotFoundInMap
			}
			p999, ok := m["allocator/mesos/allocation_run_latency_ms/p999"]
			if !ok {
				return errNotFoundInMap
			}
			p9999, ok := m["allocator/mesos/allocation_run_latency_ms/p9999"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("mean").Set(mean)
			c.(*prometheus.GaugeVec).WithLabelValues("min").Set(min)
			c.(*prometheus.GaugeVec).WithLabelValues("max").Set(max)
			c.(*prometheus.GaugeVec).WithLabelValues("p50").Set(p50)
			c.(*prometheus.GaugeVec).WithLabelValues("p90").Set(p90)
			c.(*prometheus.GaugeVec).WithLabelValues("p95").Set(p95)
			c.(*prometheus.GaugeVec).WithLabelValues("p99").Set(p99)
			c.(*prometheus.GaugeVec).WithLabelValues("p999").Set(p999)
			c.(*prometheus.GaugeVec).WithLabelValues("p9999").Set(p9999)
			return nil
		},

		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "master",
			Name:      "event_queue_dispatches",
			Help:      "Number of dispatch events in the allocator mesos event queue.",
		}): func(m metricMap, c prometheus.Collector) error {
			count, ok := m["allocator/mesos/event_queue_dispatches"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(count)
			return nil
		},
		gauge("master", "allocator_offer_filters_active", "Number of active offer filters for all frameworks within the role", "role"): func(m metricMap, c prometheus.Collector) error {
			re, err := regexp.Compile("allocator/mesos/offer_filters/roles/(.*?)/active")
			if err != nil {
				return fmt.Errorf("could not compile allocator_offer_filters_active regex: %s", err)
			}
			for metric, value := range m {
				matches := re.FindStringSubmatch(metric)
				if len(matches) != 2 {
					continue
				}
				role := matches[1]
				c.(*prometheus.GaugeVec).WithLabelValues(role).Set(value)
			}
			return nil
		},

		gauge("master", "allocator_role_quota_offered_or_allocated", "Amount of resources considered offered or allocated towards a role's quota guarantee.", "role", "resource"): func(m metricMap, c prometheus.Collector) error {
			re, err := regexp.Compile("allocator/mesos/quota/roles/(.*?)/resources/(.*?)/offered_or_allocated")
			if err != nil {
				return fmt.Errorf("could not compile allocator_role_quota_offered_or_allocated regex: %s", err)
			}
			for metric, value := range m {
				matches := re.FindStringSubmatch(metric)
				if len(matches) != 3 {
					continue
				}
				role := matches[1]
				resource := matches[2]
				c.(*prometheus.GaugeVec).WithLabelValues(role, resource).Set(value)
			}
			return nil
		},

		gauge("master", "allocator_role_shares_dominant", "Dominance factor for a role", "role"): func(m metricMap, c prometheus.Collector) error {
			re, err := regexp.Compile("allocator/mesos/roles/(.*?)/shares/dominant")
			if err != nil {
				return fmt.Errorf("could not compile allocator_role_shares_dominant regex: %s", err)
			}
			for metric, value := range m {
				matches := re.FindStringSubmatch(metric)
				if len(matches) != 2 {
					continue
				}
				role := matches[1]
				c.(*prometheus.GaugeVec).WithLabelValues(role).Set(value)
			}
			return nil
		},

		gauge("master", "allocator_role_quota_guarantee", "Amount of resources guaranteed for a role via quota", "role", "resource"): func(m metricMap, c prometheus.Collector) error {
			re, err := regexp.Compile("allocator/mesos/quota/roles/(.*?)/resources/(.*?)/guarantee")
			if err != nil {
				return fmt.Errorf("could not compile allocator_role_quota_guarantee regex: %s", err)
			}
			for metric, value := range m {
				matches := re.FindStringSubmatch(metric)
				if len(matches) != 3 {
					continue
				}
				role := matches[1]
				resource := matches[2]
				c.(*prometheus.GaugeVec).WithLabelValues(role, resource).Set(value)
			}
			return nil
		},

		gauge("master", "allocator_resources_cpus", "Number of CPUs offered or allocated", "type"): func(m metricMap, c prometheus.Collector) error {
			total, ok := m["allocator/mesos/resources/cpus/total"]
			if !ok {
				return errNotFoundInMap
			}
			offeredOrAllocated, ok := m["allocator/mesos/resources/cpus/offered_or_allocated"]
			if !ok {
				return errNotFoundInMap
			}

			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("offered_or_allocated").Set(offeredOrAllocated)
			return nil
		},

		gauge("master", "allocator_resources_disk", "Allocated or offered disk space in MB", "type"): func(m metricMap, c prometheus.Collector) error {
			total, ok := m["allocator/mesos/resources/disk/total"]
			if !ok {
				return errNotFoundInMap
			}
			offeredOrAllocated, ok := m["allocator/mesos/resources/disk/offered_or_allocated"]
			if !ok {
				return errNotFoundInMap
			}

			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("offered_or_allocated").Set(offeredOrAllocated)
			return nil
		},

		gauge("master", "allocator_resources_mem", "Allocated or offered memory in MB", "type"): func(m metricMap, c prometheus.Collector) error {
			total, ok := m["allocator/mesos/resources/mem/total"]
			if !ok {
				return errNotFoundInMap
			}
			offeredOrAllocated, ok := m["allocator/mesos/resources/mem/offered_or_allocated"]
			if !ok {
				return errNotFoundInMap
			}

			c.(*prometheus.GaugeVec).WithLabelValues("total").Set(total)
			c.(*prometheus.GaugeVec).WithLabelValues("offered_or_allocated").Set(offeredOrAllocated)
			return nil
		},

		// Frameworks metrics
		counter("master", "frameworks_messages", "Messages passed around with the frameworks", "framework", "type"): func(m metricMap, c prometheus.Collector) error {
			re, err := regexp.Compile("frameworks/(.*?)/messages_(.*?)$")
			if err != nil {
				return fmt.Errorf("could not compile frameworks_messages regex: %s", err)
			}
			for metric, messageCount := range m {
				matches := re.FindStringSubmatch(metric)
				if len(matches) != 3 {
					continue
				}
				framework := matches[1]
				messageStatus := matches[2]
				if len(framework) > 0 && len(messageStatus) > 0 {
					c.(*settableCounterVec).Set(messageCount, framework, messageStatus)
				}
			}
			return nil
		},

		// Registrar stats
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "registrar",
			Name:      "registry_size_bytes",
			Help:      "Size of the registry in bytes",
		}): func(m metricMap, c prometheus.Collector) error {
			size, ok := m["registrar/registry_size_bytes"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(size)
			return nil
		},
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "registrar",
			Name:      "queued_operations",
			Help:      "Number of operations in the registry queue",
		}): func(m metricMap, c prometheus.Collector) error {
			ops, ok := m["registrar/queued_operations"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(ops)
			return nil
		},
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "registrar",
			Name:      "state_fetch_ms",
			Help:      "Duration of state JSON fetch in ms",
		}): func(m metricMap, c prometheus.Collector) error {
			ms, ok := m["registrar/state_fetch_ms"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(ms)
			return nil
		},
		gauge("registrar", "state_store_ms", "Duration of state json store in ms.", "type"): func(m metricMap, c prometheus.Collector) error {
			mean, ok := m["registrar/state_store_ms"]
			if !ok {
				return errNotFoundInMap
			}
			min, ok := m["registrar/state_store_ms/min"]
			if !ok {
				return errNotFoundInMap
			}
			max, ok := m["registrar/state_store_ms/max"]
			if !ok {
				return errNotFoundInMap
			}
			p50, ok := m["registrar/state_store_ms/p50"]
			if !ok {
				return errNotFoundInMap
			}
			p90, ok := m["registrar/state_store_ms/p90"]
			if !ok {
				return errNotFoundInMap
			}
			p95, ok := m["registrar/state_store_ms/p95"]
			if !ok {
				return errNotFoundInMap
			}
			p99, ok := m["registrar/state_store_ms/p99"]
			if !ok {
				return errNotFoundInMap
			}
			p999, ok := m["registrar/state_store_ms/p999"]
			if !ok {
				return errNotFoundInMap
			}
			p9999, ok := m["registrar/state_store_ms/p9999"]
			if !ok {
				return errNotFoundInMap
			}
			c.(*prometheus.GaugeVec).WithLabelValues("mean").Set(mean)
			c.(*prometheus.GaugeVec).WithLabelValues("min").Set(min)
			c.(*prometheus.GaugeVec).WithLabelValues("max").Set(max)
			c.(*prometheus.GaugeVec).WithLabelValues("p50").Set(p50)
			c.(*prometheus.GaugeVec).WithLabelValues("p90").Set(p90)
			c.(*prometheus.GaugeVec).WithLabelValues("p95").Set(p95)
			c.(*prometheus.GaugeVec).WithLabelValues("p99").Set(p99)
			c.(*prometheus.GaugeVec).WithLabelValues("p999").Set(p999)
			c.(*prometheus.GaugeVec).WithLabelValues("p9999").Set(p9999)
			return nil
		},
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "registrar",
			Name:      "log_recovered",
			Help:      "Recovered status of the registrar log",
		}): func(m metricMap, c prometheus.Collector) error {
			recovered, ok := m["registrar/log/recovered"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(recovered)
			return nil
		},
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "registrar",
			Name:      "log_ensemble_size",
			Help:      "Ensemble size of the registrar log",
		}): func(m metricMap, c prometheus.Collector) error {
			size, ok := m["registrar/log/ensemble_size"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(size)
			return nil
		},

		// Overlay log
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "overlay",
			Name:      "log_recovered",
			Help:      "Recovered status of the overlay log",
		}): func(m metricMap, c prometheus.Collector) error {
			recovered, ok := m["overlay/log/recovered"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(recovered)
			return nil
		},
		prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "mesos",
			Subsystem: "overlay",
			Name:      "log_ensemble_size",
			Help:      "Ensemble size of the overlay log",
		}): func(m metricMap, c prometheus.Collector) error {
			size, ok := m["overlay/log/ensemble_size"]
			if !ok {
				return errNotFoundInMap
			}
			c.(prometheus.Gauge).Set(size)
			return nil
		},

		// END
	}
	return newMetricCollector(httpClient, metrics)
}
