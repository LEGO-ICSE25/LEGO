package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type ZhimiAirpurifier struct {
	power string
	aqi string
	mode string
	buzzer string
	child_lock string
	filter_life_remaining string
	filter_hours_used string
	motor_speed string
	favorite_rpm string
	average_aqi string
	humidity string
	tvoc string
	temperature string
	pm10_density string
	fan_level string
	led string
	led_brightness string
	buzzer_volume string
	favorite_level string
	use_time string
	purify_volume string
	filter_rfid_product_id string
	filter_rfid_tag string
	led_brightness_level string
	anion string
	filter_left_time string
	gestures string
}

func (airpurifier *ZhimiAirpurifier) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "power")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "aqi")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["aqi"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["mode"] = strings.TrimRight(string(output), "\n")

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

	command = exec.Command("python", "read.py", "filter_life_remaining")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["filter_life_remaining"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "filter_hours_used")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["filter_hours_used"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "motor_speed")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["motor_speed"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "favorite_rpm")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["favorite_rpm"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "average_aqi")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["average_aqi"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "humidity")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["humidity"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "tvoc")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["tvoc"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "temperature")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["temperature"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "pm10_density")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["pm10_density"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "fan_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["fan_level"] = strings.TrimRight(string(output), "\n")

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

	command = exec.Command("python", "read.py", "buzzer_volume")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["buzzer_volume"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "favorite_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["favorite_level"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "use_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["use_time"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "purify_volume")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["purify_volume"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "filter_rfid_product_id")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["filter_rfid_product_id"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "filter_rfid_tag")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["filter_rfid_tag"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "led_brightness_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["led_brightness_level"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "anion")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["anion"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "filter_left_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["filter_left_time"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "gestures")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["gestures"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (airpurifier *ZhimiAirpurifier) SetPower(value string) {

	command := exec.Command("python", "write.py", "power", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.power = value
}

func (airpurifier *ZhimiAirpurifier) SetMode(value string) {

	command := exec.Command("python", "write.py", "mode", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.mode = value
}

func (airpurifier *ZhimiAirpurifier) SetBuzzer(value string) {

	command := exec.Command("python", "write.py", "buzzer", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.buzzer = value
}

func (airpurifier *ZhimiAirpurifier) SetChild_lock(value string) {

	command := exec.Command("python", "write.py", "child_lock", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.child_lock = value
}

func (airpurifier *ZhimiAirpurifier) SetFavorite_rpm(value string) {

	command := exec.Command("python", "write.py", "favorite_rpm", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.favorite_rpm = value
}

func (airpurifier *ZhimiAirpurifier) SetFan_level(value string) {

	command := exec.Command("python", "write.py", "fan_level", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.fan_level = value
}

func (airpurifier *ZhimiAirpurifier) SetLed(value string) {

	command := exec.Command("python", "write.py", "led", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.led = value
}

func (airpurifier *ZhimiAirpurifier) SetLed_brightness(value string) {

	command := exec.Command("python", "write.py", "led_brightness", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.led_brightness = value
}

func (airpurifier *ZhimiAirpurifier) SetBuzzer_volume(value string) {

	command := exec.Command("python", "write.py", "buzzer_volume", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.buzzer_volume = value
}

func (airpurifier *ZhimiAirpurifier) SetFavorite_level(value string) {

	command := exec.Command("python", "write.py", "favorite_level", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.favorite_level = value
}

func (airpurifier *ZhimiAirpurifier) SetLed_brightness_level(value string) {

	command := exec.Command("python", "write.py", "led_brightness_level", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.led_brightness_level = value
}

func (airpurifier *ZhimiAirpurifier) SetAnion(value string) {

	command := exec.Command("python", "write.py", "anion", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.anion = value
}

func (airpurifier *ZhimiAirpurifier) SetGestures(value string) {

	command := exec.Command("python", "write.py", "gestures", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.gestures = value
}

func NewDevice() *ZhimiAirpurifier {
	airpurifier := &ZhimiAirpurifier{
		power: "default",
		aqi: "default",
		mode: "default",
		buzzer: "default",
		child_lock: "default",
		filter_life_remaining: "default",
		filter_hours_used: "default",
		motor_speed: "default",
		favorite_rpm: "default",
		average_aqi: "default",
		humidity: "default",
		tvoc: "default",
		temperature: "default",	
		pm10_density: "default",
		fan_level: "default",
		led: "default",
		led_brightness: "default",
		buzzer_volume: "default",
		favorite_level: "default",	
		use_time: "default",
		purify_volume: "default",
		filter_rfid_product_id: "default",
		filter_rfid_tag: "default",	
		led_brightness_level: "default",
		anion: "default",
		filter_left_time: "default",
		gestures: "default",
	}
	return airpurifier
}

