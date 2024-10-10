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

func (light *BroadlinkLight) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (light *BroadlinkLight) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_power := *Update.Twin["power"].Expected.Value
	if new_power != light.power {
		light.SetPower(new_power)
	}
	new_brightness := *Update.Twin["brightness"].Expected.Value
	if new_brightness != light.brightness {
		light.SetBrightness(new_brightness)
	}
	new_red := *Update.Twin["red"].Expected.Value
	if new_red != light.red {
		light.SetRed(new_red)
	}
	new_blue := *Update.Twin["blue"].Expected.Value
	if new_blue != light.blue {
		light.SetBlue(new_blue)
	}
	new_green := *Update.Twin["green"].Expected.Value
	if new_green != light.green {
		light.SetGreen(new_green)
	}
	new_colortemp := *Update.Twin["colortemp"].Expected.Value
	if new_colortemp != light.colortemp {
		light.SetColortemp(new_colortemp)
	}
	new_hue := *Update.Twin["hue"].Expected.Value
	if new_hue != light.hue {
		light.SetHue(new_hue)
	}
	new_saturation := *Update.Twin["saturation"].Expected.Value
	if new_saturation != light.saturation {
		light.SetSaturation(new_saturation)
	}
	new_transitionduration := *Update.Twin["transitionduration"].Expected.Value
	if new_transitionduration != light.transitionduration {
		light.SetTransitionduration(new_transitionduration)
	}
	new_maxworktime := *Update.Twin["maxworktime"].Expected.Value
	if new_maxworktime != light.maxworktime {
		light.SetMaxworktime(new_maxworktime)
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

	light := NewDevice()
	klog.Infoln("New light establish.\n")

	go light.Initialize()

	light.Run()

	select {
	case <-stopchan:
		cli.Disconnect(250)
		klog.Infoln("Interrupt, exit.\n")
		break
	}
}
