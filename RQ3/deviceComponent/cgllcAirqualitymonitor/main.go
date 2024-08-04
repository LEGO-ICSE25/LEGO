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

func (monitor *CgllcAirqualitymonitor) publishToMqtt(properties map[string]string) {
	updateMessage := createActualUpdateMessage(properties)
	twinUpdateBody, _ := json.Marshal(updateMessage)

	token := cli.Publish(topic, 0, false, twinUpdateBody)
	if token.Wait() && token.Error() != nil {
		klog.Infoln(token.Error())
	}
}

func (monitor *CgllcAirqualitymonitor) subcribeFromMqtt(client mqtt.Client, msg mqtt.Message) {

	Update := &DeviceTwinDelta{}
	err := json.Unmarshal(msg.Payload(), Update)
	if err != nil {
		klog.Infoln("Unmarshal error: %v\n", err)
	}

	new_monitoring_frequency := *Update.Twin["monitoring_frequency"].Expected.Value
	if new_monitoring_frequency != monitor.monitoring_frequency {
		monitor.SetMonitoring_frequency(new_monitoring_frequency)
	}
	new_screen_off := *Update.Twin["screen_off"].Expected.Value
	if new_screen_off != monitor.screen_off {
		monitor.SetScreen_off(new_screen_off)
	}
	new_device_off := *Update.Twin["device_off"].Expected.Value
	if new_device_off != monitor.device_off {
		monitor.SetDevice_off(new_device_off)
	}
	new_temperature_unit := *Update.Twin["temperature_unit"].Expected.Value
	if new_temperature_unit != monitor.temperature_unit {
		monitor.SetTemperature_unit(new_temperature_unit)
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

	monitor := NewDevice()
	klog.Infoln("New monitor establish.\n")

	go monitor.Initialize()

	monitor.Run()

	select {
	case <-stopchan:
		cli.Disconnect(250)
		klog.Infoln("Interrupt, exit.\n")
		break
	}
}
