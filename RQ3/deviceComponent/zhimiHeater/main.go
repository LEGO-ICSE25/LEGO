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

func (heater *ZhimiHeater) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (heater *ZhimiHeater) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_power := *Update.Twin["power"].Expected.Value
	if new_power != heater.power {
		heater.SetPower(new_power)
	}
	new_target_temperature := *Update.Twin["target_temperature"].Expected.Value
	if new_target_temperature != heater.target_temperature {
		heater.SetTarget_temperature(new_target_temperature)
	}
	new_countdown_time := *Update.Twin["countdown_time"].Expected.Value
	if new_countdown_time != heater.countdown_time {
		heater.SetCountdown_time(new_countdown_time)
	}
	new_led_brightness := *Update.Twin["led_brightness"].Expected.Value
	if new_led_brightness != heater.led_brightness {
		heater.SetLed_brightness(new_led_brightness)
	}
	new_child_lock := *Update.Twin["child_lock"].Expected.Value
	if new_child_lock != heater.child_lock {
		heater.SetChild_lock(new_child_lock)
	}
	new_buzzer := *Update.Twin["buzzer"].Expected.Value
	if new_buzzer != heater.buzzer {
		heater.SetBuzzer(new_buzzer)
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

	heater := NewDevice()
	klog.Infoln("New heater establish.\n")

	go heater.Initialize()

	heater.Run()

	select {
	case <-stopchan:
		cli.Disconnect(250)
		klog.Infoln("Interrupt, exit.\n")
		break
	}
}
