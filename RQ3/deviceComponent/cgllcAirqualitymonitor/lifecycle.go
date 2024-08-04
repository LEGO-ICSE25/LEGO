package main

import (
	"time"
	"k8s.io/klog/v2"
)

func (monitor *CgllcAirqualitymonitor) Initialize() {

	for {
		properties := monitor.GetStates()
		monitor.publishToMqtt(properties)
		time.Sleep(5 * time.Second)
	}
}

func (monitor *CgllcAirqualitymonitor) Run() {
	token := cli.Subscribe(topic+"/delta", 0, monitor.subcribeFromMqtt)

	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}
