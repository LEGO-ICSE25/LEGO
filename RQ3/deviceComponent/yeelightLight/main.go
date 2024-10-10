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

func (light *YeelightLight) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (light *YeelightLight) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

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
	new_color := *Update.Twin["color"].Expected.Value
	if new_color != light.color {
		light.SetColor(new_color)
	}
	new_mode := *Update.Twin["mode"].Expected.Value
	if new_mode != light.mode {
		light.SetMode(new_mode)
	}
	new_developer_mode := *Update.Twin["developer_mode"].Expected.Value
	if new_developer_mode != light.developer_mode {
		light.SetDeveloper_mode(new_developer_mode)
	}
	new_cfg_save_state := *Update.Twin["cfg_save_state"].Expected.Value
	if new_cfg_save_state != light.cfg_save_state {
		light.SetCfg_save_state(new_cfg_save_state)
	}
	new_name := *Update.Twin["name"].Expected.Value
	if new_name != light.name {
		light.SetName(new_name)
	}
	new_delay := *Update.Twin["delay"].Expected.Value
	if new_delay != light.delay {
		light.SetDelay(new_delay)
	}
	new_music_on := *Update.Twin["music_on"].Expected.Value
	if new_music_on != light.music_on {
		light.SetMusic_on(new_music_on)
	}
	new_active_mode := *Update.Twin["active_mode"].Expected.Value
	if new_active_mode != light.active_mode {
		light.SetActive_mode(new_active_mode)
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
