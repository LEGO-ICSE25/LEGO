package main

import (
	"time"
	"k8s.io/klog/v2"
)

func (fan *ZhimiFan) Initialize() {

	for {
		properties := fan.GetStates()
		fan.publishToMqtt(properties)
		time.Sleep(5 * time.Second)
	}
}

func (fan *ZhimiFan) Run() {
	token := cli.Subscribe(topic+"/delta", 0, fan.subcribeFromMqtt)

	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}
