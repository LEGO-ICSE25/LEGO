package main

import (
	"time"
	"k8s.io/klog/v2"
)

func (humidifier *ZhimiHumidifier) Initialize() {

	for {
		properties := humidifier.GetStates()
		humidifier.publishToMqtt(properties)
		time.Sleep(5 * time.Second)
	}
}

func (humidifier *ZhimiHumidifier) Run() {
	token := cli.Subscribe(topic+"/delta", 0, humidifier.subcribeFromMqtt)

	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}
