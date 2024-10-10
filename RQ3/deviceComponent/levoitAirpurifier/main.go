package main

import (
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"k8s.io/klog/v2"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var cli mqtt.Client

func createActualUpdateMessage(properties map[string]string) DeviceTwinUpdate {

	var deviceTwinUpdateMessage DeviceTwinUpdate
	actualMap := make(map[string]*MsgTwin)

	for key, value := range properties {
		temp := value
		actualMap[key] = &MsgTwin{Actual: &TwinValue{Value: &temp}, Metadata: &TypeMetadata{Type: "Updated"}}
	}

	deviceTwinUpdateMessage.Twin = actualMap
	return deviceTwinUpdateMessage
}

func (airpurifier *LevoitAirpurifier) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (airpurifier *LevoitAirpurifier) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_device_status := *Update.Twin["device_status"].Expected.Value
	if new_device_status != airpurifier.device_status {
		airpurifier.SetDevice_status(new_device_status)
	}
	new_screen_status := *Update.Twin["screen_status"].Expected.Value
	if new_screen_status != airpurifier.screen_status {
		airpurifier.SetScreen_status(new_screen_status)
	}
	new_mode := *Update.Twin["mode"].Expected.Value
	if new_mode != airpurifier.mode {
		airpurifier.SetMode(new_mode)
	}
	new_level := *Update.Twin["level"].Expected.Value
	if new_level != airpurifier.level {
		airpurifier.SetLevel(new_level)
	}
}

func connectToMqtt() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(mqttUrl)

	cli = mqtt.NewClient(opts)

	token := cli.Connect()
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}

	return cli
}


func main() {
	stopchan := make(chan os.Signal)
	signal.Notify(stopchan, syscall.SIGINT, syscall.SIGKILL)

	defer close(stopchan)

	cli = connectToMqtt()

	airpurifier := NewDevice()
	klog.Infoln("New airpurifier establish.\n")

	go airpurifier.Initialize()

	airpurifier.Run()

	select {
	case <-stopchan:
		cli.Disconnect(250)
		klog.Infoln("Interrupt, exit.\n")
		break
	}
}
