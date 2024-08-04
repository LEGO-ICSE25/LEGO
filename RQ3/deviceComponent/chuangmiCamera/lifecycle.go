package main

import (
	"time"
	"k8s.io/klog/v2"
)

func (camera *ChuangmiCamera) Initialize() {

	for {
		properties := camera.GetStates()
		camera.publishToMqtt(properties)
		time.Sleep(5 * time.Second)
	}
}

func (camera *ChuangmiCamera) Run() {
	token := cli.Subscribe(topic+"/delta", 0, camera.subcribeFromMqtt)

	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}
