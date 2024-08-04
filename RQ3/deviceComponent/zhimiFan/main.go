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

func (fan *ZhimiFan) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (fan *ZhimiFan) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_power := *Update.Twin["power"].Expected.Value
	if new_power != fan.power {
		fan.SetPower(new_power)
	}
	new_led := *Update.Twin["led"].Expected.Value
	if new_led != fan.led {
		fan.SetLed(new_led)
	}
	new_led_brightness := *Update.Twin["led_brightness"].Expected.Value
	if new_led_brightness != fan.led_brightness {
		fan.SetLed_brightness(new_led_brightness)
	}
	new_buzzer := *Update.Twin["buzzer"].Expected.Value
	if new_buzzer != fan.buzzer {
		fan.SetBuzzer(new_buzzer)
	}
	new_child_lock := *Update.Twin["child_lock"].Expected.Value
	if new_child_lock != fan.child_lock {
		fan.SetChild_lock(new_child_lock)
	}
	new_natural_level := *Update.Twin["natural_level"].Expected.Value
	if new_natural_level != fan.natural_level {
		fan.SetNatural_level(new_natural_level)
	}
	new_speed_level := *Update.Twin["speed_level"].Expected.Value
	if new_speed_level != fan.speed_level {
		fan.SetSpeed_level(new_speed_level)
	}
	new_angle_enable := *Update.Twin["angle_enable"].Expected.Value
	if new_angle_enable != fan.angle_enable {
		fan.SetAngle_enable(new_angle_enable)
	}
	new_poweroff_time := *Update.Twin["poweroff_time"].Expected.Value
	if new_poweroff_time != fan.poweroff_time {
		fan.SetPoweroff_time(new_poweroff_time)
	}
	new_angle := *Update.Twin["angle"].Expected.Value
	if new_angle != fan.angle {
		fan.SetAngle(new_angle)
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

	fan := NewDevice()
	klog.Infoln("New fan establish.\n")

	go fan.Initialize()

	fan.Run()

	select {
	case <-stopchan:
		cli.Disconnect(250)
		klog.Infoln("Interrupt, exit.\n")
		break
	}
}
