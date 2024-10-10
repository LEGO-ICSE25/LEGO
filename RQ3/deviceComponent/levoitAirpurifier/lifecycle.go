package main

import (
	"time"
	"k8s.io/klog/v2"
)

func (airpurifier *LevoitAirpurifier) Initialize() {

	for {
		properties := airpurifier.GetStates()
		airpurifier.publishToMqtt(properties)
		time.Sleep(5 * time.Second)
	}
}

func (airpurifier *LevoitAirpurifier) Run() {
	token := cli.Subscribe(topic+"/delta", 0, airpurifier.subcribeFromMqtt)

	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}
