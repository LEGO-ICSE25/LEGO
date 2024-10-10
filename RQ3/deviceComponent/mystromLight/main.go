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

func (light *MystromLight) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (light *MystromLight) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_power := *Update.Twin["power"].Expected.Value
	if new_power != light.power {
		light.SetPower(new_power)
	}
	new_consumption := *Update.Twin["consumption"].Expected.Value
	if new_consumption != light.consumption {
		light.SetConsumption(new_consumption)
	}
	new_color := *Update.Twin["color"].Expected.Value
	if new_color != light.color {
		light.SetColor(new_color)
	}
	new_mode := *Update.Twin["mode"].Expected.Value
	if new_mode != light.mode {
		light.SetMode(new_mode)
	}
	new_transition_time := *Update.Twin["transition_time"].Expected.Value
	if new_transition_time != light.transition_time {
		light.SetTransition_time(new_transition_time)
	}
	new_state := *Update.Twin["state"].Expected.Value
	if new_state != light.state {
		light.SetState(new_state)
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
