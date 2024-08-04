package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type ZhimiFan struct {
	power string
	humidity string
	temperature string
	led string
	led_brightness string
	buzzer string
	child_lock string
	natural_level string
	speed_level string
	angle_enable string
	battery string
	bat_charge string
	bat_state string
	ac_power string
	poweroff_time string
	speed string
	angle string
	use_time string
	button_pressed string
}

func (fan *ZhimiFan) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "power")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power"] = strings.TrimRight(string(output), "\n")

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

	command = exec.Command("python", "read.py", "led")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["led"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "led_brightness")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["led_brightness"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "buzzer")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["buzzer"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "child_lock")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["child_lock"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "natural_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["natural_level"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "speed_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["speed_level"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "angle_enable")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["angle_enable"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "battery")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["battery"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "bat_charge")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["bat_charge"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "bat_state")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["bat_state"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "ac_power")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["ac_power"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "poweroff_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["poweroff_time"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "speed")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["speed"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "angle")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["angle"] = strings.TrimRight(string(output), "\n")

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

	return properties
}

func (fan *ZhimiFan) SetPower(value string) {

	command := exec.Command("python", "write.py", "power", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.power = value
}

func (fan *ZhimiFan) SetLed(value string) {

	command := exec.Command("python", "write.py", "led", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.led = value
}

func (fan *ZhimiFan) SetLed_brightness(value string) {

	command := exec.Command("python", "write.py", "led_brightness", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.led_brightness = value
}

func (fan *ZhimiFan) SetBuzzer(value string) {

	command := exec.Command("python", "write.py", "buzzer", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.buzzer = value
}

func (fan *ZhimiFan) SetChild_lock(value string) {

	command := exec.Command("python", "write.py", "child_lock", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.child_lock = value
}

func (fan *ZhimiFan) SetNatural_level(value string) {

	command := exec.Command("python", "write.py", "natural_level", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.natural_level = value
}

func (fan *ZhimiFan) SetSpeed_level(value string) {

	command := exec.Command("python", "write.py", "speed_level", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.speed_level = value
}

func (fan *ZhimiFan) SetAngle_enable(value string) {

	command := exec.Command("python", "write.py", "angle_enable", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.angle_enable = value
}

func (fan *ZhimiFan) SetPoweroff_time(value string) {

	command := exec.Command("python", "write.py", "poweroff_time", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.poweroff_time = value
}

func (fan *ZhimiFan) SetAngle(value string) {

	command := exec.Command("python", "write.py", "angle", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	fan.angle = value
}

func NewDevice() *ZhimiFan {
	fan := &ZhimiFan{
		power: "default",
		humidity: "default",
		temperature: "default",
		led: "default",
		led_brightness: "default",
		buzzer: "default",
		child_lock: "default",
		natural_level: "default",
		speed_level: "default",
		angle_enable: "default",
		battery: "default",
		bat_charge: "default",
		bat_state: "default",	
		ac_power: "default",
		poweroff_time: "default",
		speed: "default",
		angle: "default",
		use_time: "default",
		button_pressed: "default",	
	}
	return fan
}

