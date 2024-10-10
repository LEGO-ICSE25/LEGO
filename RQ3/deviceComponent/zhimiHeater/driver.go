package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type ZhimiHeater struct {
	power string
	target_temperature string
	temperature string
	countdown_time string
	led_brightness string
	relative_humidity string
	child_lock string
	buzzer string
}

func (heater *ZhimiHeater) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "power")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "target_temperature")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["target_temperature"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "temperature")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["temperature"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "countdown_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["countdown_time"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "led_brightness")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["led_brightness"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "relative_humidity")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["relative_humidity"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "child_lock")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["child_lock"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "buzzer")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["buzzer"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (heater *ZhimiHeater) SetPower(value string) {

	command := exec.Command("python", "write.py", "power", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	heater.power = value
}

func (heater *ZhimiHeater) SetTarget_temperature(value string) {

	command := exec.Command("python", "write.py", "target_temperature", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	heater.target_temperature = value
}

func (heater *ZhimiHeater) SetCountdown_time(value string) {

	command := exec.Command("python", "write.py", "countdown_time", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	heater.countdown_time = value
}

func (heater *ZhimiHeater) SetLed_brightness(value string) {

	command := exec.Command("python", "write.py", "led_brightness", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	heater.led_brightness = value
}

func (heater *ZhimiHeater) SetChild_lock(value string) {

	command := exec.Command("python", "write.py", "child_lock", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	heater.child_lock = value
}

func (heater *ZhimiHeater) SetBuzzer(value string) {

	command := exec.Command("python", "write.py", "buzzer", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	heater.buzzer = value
}

func NewDevice() *ZhimiHeater {
	heater := &ZhimiHeater{
		power: "default",
		target_temperature: "default",
		temperature: "default",
		countdown_time: "default",
		led_brightness: "default",
		relative_humidity: "default",
		child_lock: "default",
		buzzer: "default",
	}
	return heater
}

