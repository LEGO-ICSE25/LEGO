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

func (airpurifier *ZhimiAirpurifier) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (airpurifier *ZhimiAirpurifier) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_power := *Update.Twin["power"].Expected.Value
	if new_power != airpurifier.power {
		airpurifier.SetPower(new_power)
	}
	new_mode := *Update.Twin["mode"].Expected.Value
	if new_mode != airpurifier.mode {
		airpurifier.SetMode(new_mode)
	}
	new_buzzer := *Update.Twin["buzzer"].Expected.Value
	if new_buzzer != airpurifier.buzzer {
		airpurifier.SetBuzzer(new_buzzer)
	}
	new_child_lock := *Update.Twin["child_lock"].Expected.Value
	if new_child_lock != airpurifier.child_lock {
		airpurifier.SetChild_lock(new_child_lock)
	}
	new_favorite_rpm := *Update.Twin["favorite_rpm"].Expected.Value
	if new_favorite_rpm != airpurifier.favorite_rpm {
		airpurifier.SetFavorite_rpm(new_favorite_rpm)
	}
	new_fan_level := *Update.Twin["fan_level"].Expected.Value
	if new_fan_level != airpurifier.fan_level {
		airpurifier.SetFan_level(new_fan_level)
	}
	new_led := *Update.Twin["led"].Expected.Value
	if new_led != airpurifier.led {
		airpurifier.SetLed(new_led)
	}
	new_led_brightness := *Update.Twin["led_brightness"].Expected.Value
	if new_led_brightness != airpurifier.led_brightness {
		airpurifier.SetLed_brightness(new_led_brightness)
	}
	new_buzzer_volume := *Update.Twin["buzzer_volume"].Expected.Value
	if new_buzzer_volume != airpurifier.buzzer_volume {
		airpurifier.SetBuzzer_volume(new_buzzer_volume)
	}
	new_favorite_level := *Update.Twin["favorite_level"].Expected.Value
	if new_favorite_level != airpurifier.favorite_level {
		airpurifier.SetFavorite_level(new_favorite_level)
	}
	new_led_brightness_level := *Update.Twin["led_brightness_level"].Expected.Value
	if new_led_brightness_level != airpurifier.led_brightness_level {
		airpurifier.SetLed_brightness_level(new_led_brightness_level)
	}
	new_anion := *Update.Twin["anion"].Expected.Value
	if new_anion != airpurifier.anion {
		airpurifier.SetAnion(new_anion)
	}
	new_gestures := *Update.Twin["gestures"].Expected.Value
	if new_gestures != airpurifier.gestures {
		airpurifier.SetGestures(new_gestures)
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
