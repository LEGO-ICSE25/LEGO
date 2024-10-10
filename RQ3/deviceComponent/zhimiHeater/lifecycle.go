package main

import (
	"time"
	"k8s.io/klog/v2"
)

func (heater *ZhimiHeater) Initialize() {

	for {
		properties := heater.GetStates()
		heater.publishToMqtt(properties)
		time.Sleep(5 * time.Second)
	}
}

func (heater *ZhimiHeater) Run() {
	token := cli.Subscribe(topic+"/delta", 0, heater.subcribeFromMqtt)

	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}
