package main

import (
	"time"
	"k8s.io/klog/v2"
)

func (light *MystromLight) Initialize() {

	for {
		properties := light.GetStates()
		light.publishToMqtt(properties)
		time.Sleep(5 * time.Second)
	}
}

func (light *MystromLight) Run() {
	token := cli.Subscribe(topic+"/delta", 0, light.subcribeFromMqtt)

	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}
