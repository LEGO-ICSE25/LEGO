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

func (humidifier *ZhimiHumidifier) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (humidifier *ZhimiHumidifier) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_power := *Update.Twin["power"].Expected.Value
	if new_power != humidifier.power {
		humidifier.SetPower(new_power)
	}
	new_led := *Update.Twin["mode"].Expected.Value
	if new_led != humidifier.mode {
		humidifier.SetMode(new_led)
	}
	new_target_humidity := *Update.Twin["target_humidity"].Expected.Value
	if new_target_humidity != humidifier.target_humidity {
		humidifier.SetTarget_humidity(new_target_humidity)
	}
	new_dry := *Update.Twin["dry"].Expected.Value
	if new_dry != humidifier.dry {
		humidifier.SetDry(new_dry)
	}
	new_speed_level := *Update.Twin["speed_level"].Expected.Value
	if new_speed_level != humidifier.speed_level {
		humidifier.SetSpeed_level(new_speed_level)
	}
	new_buzzer := *Update.Twin["buzzer"].Expected.Value
	if new_buzzer != humidifier.buzzer {
		humidifier.SetBuzzer(new_buzzer)
	}
	new_led_brightness := *Update.Twin["led_brightness"].Expected.Value
	if new_led_brightness != humidifier.led_brightness {
		humidifier.SetLed_brightness(new_led_brightness)
	}
	new_child_lock := *Update.Twin["child_lock"].Expected.Value
	if new_child_lock != humidifier.child_lock {
		humidifier.SetChild_lock(new_child_lock)
	}
	new_clean_mode := *Update.Twin["clean_mode"].Expected.Value
	if new_clean_mode != humidifier.clean_mode {
		humidifier.SetClean_mode(new_clean_mode)
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

	humidifier := NewDevice()
	klog.Infoln("New humidifier establish.\n")

	go humidifier.Initialize()

	humidifier.Run()

	select {
	case <-stopchan:
		cli.Disconnect(250)
		klog.Infoln("Interrupt, exit.\n")
		break
	}
}
