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

func (camera *ChuangmiCamera) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (camera *ChuangmiCamera) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_power := *Update.Twin["power"].Expected.Value
	if new_power != camera.power {
		camera.SetPower(new_power)
	}
	new_motion_record := *Update.Twin["motion_record"].Expected.Value
	if new_motion_record != camera.motion_record {
		camera.SetMotion_record(new_motion_record)
	}
	new_light := *Update.Twin["light"].Expected.Value
	if new_light != camera.light {
		camera.SetLight(new_light)
	}
	new_full_color := *Update.Twin["full_color"].Expected.Value
	if new_full_color != camera.full_color {
		camera.SetFull_color(new_full_color)
	}
	new_flip := *Update.Twin["flip"].Expected.Value
	if new_flip != camera.flip {
		camera.SetFlip(new_flip)
	}
	new_improve_program := *Update.Twin["improve_program"].Expected.Value
	if new_improve_program != camera.improve_program {
		camera.SetImprove_program(new_improve_program)
	}
	new_wdr := *Update.Twin["wdr"].Expected.Value
	if new_wdr != camera.wdr {
		camera.SetWdr(new_wdr)
	}
	new_watermark := *Update.Twin["watermark"].Expected.Value
	if new_watermark != camera.watermark {
		camera.SetWatermark(new_watermark)
	}
	new_night_mode := *Update.Twin["night_mode"].Expected.Value
	if new_night_mode != camera.night_mode {
		camera.SetNight_mode(new_night_mode)
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

	camera := NewDevice()
	klog.Infoln("New camera establish.\n")

	go camera.Initialize()

	camera.Run()

	select {
	case <-stopchan:
		cli.Disconnect(250)
		klog.Infoln("Interrupt, exit.\n")
		break
	}
}
