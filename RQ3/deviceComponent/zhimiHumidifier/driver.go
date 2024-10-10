package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type ZhimiHumidifier struct {
	power string
	fault string
	mode string
	target_humidity string
	water_level string
	dry string
	use_time string
	button_pressed string
	speed_level string
	humidity string
	temperature string
	fahrenheit string
	buzzer string
	led_brightness string
	child_lock string
	actual_speed string
	power_time string
	clean_mode string
}

func (humidifier *ZhimiHumidifier) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "power")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "fault")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["fault"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["mode"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "target_humidity")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["target_humidity"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "water_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["water_level"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "dry")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["dry"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "use_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["use_time"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "button_pressed")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["button_pressed"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "speed_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["speed_level"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "humidity")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["humidity"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "temperature")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["temperature"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "fahrenheit")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["fahrenheit"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "buzzer")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["buzzer"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "led_brightness")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["led_brightness"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "child_lock")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["child_lock"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "actual_speed")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["actual_speed"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "power_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power_time"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "clean_mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["clean_mode"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (humidifier *ZhimiHumidifier) SetPower(value string) {

	command := exec.Command("python", "write.py", "power", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.power = value
}

func (humidifier *ZhimiHumidifier) SetMode(value string) {

	command := exec.Command("python", "write.py", "mode", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.mode = value
}

func (humidifier *ZhimiHumidifier) SetTarget_humidity(value string) {

	command := exec.Command("python", "write.py", "target_humidity", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.target_humidity = value
}

func (humidifier *ZhimiHumidifier) SetDry(value string) {

	command := exec.Command("python", "write.py", "dry", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.dry = value
}

func (humidifier *ZhimiHumidifier) SetSpeed_level(value string) {

	command := exec.Command("python", "write.py", "speed_level", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.speed_level = value
}

func (humidifier *ZhimiHumidifier) SetBuzzer(value string) {

	command := exec.Command("python", "write.py", "buzzer", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.buzzer = value
}

func (humidifier *ZhimiHumidifier) SetLed_brightness(value string) {

	command := exec.Command("python", "write.py", "led_brightness", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.led_brightness = value
}

func (humidifier *ZhimiHumidifier) SetChild_lock(value string) {

	command := exec.Command("python", "write.py", "child_lock", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.child_lock = value
}

func (humidifier *ZhimiHumidifier) SetClean_mode(value string) {

	command := exec.Command("python", "write.py", "clean_mode", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	humidifier.clean_mode = value
}

func NewDevice() *ZhimiHumidifier {
	humidifier := &ZhimiHumidifier{
		power: "default",
		fault: "default",
		mode: "default",
		target_humidity: "default",
		water_level: "default",
		dry: "default",
		use_time: "default",
		button_pressed: "default",
		speed_level: "default",
		humidity: "default",
		temperature: "default",
		fahrenheit: "default",
		buzzer: "default",	
		led_brightness: "default",
		child_lock: "default",
		actual_speed: "default",
		power_time: "default",
		clean_mode: "default",
	}
	return humidifier
}

