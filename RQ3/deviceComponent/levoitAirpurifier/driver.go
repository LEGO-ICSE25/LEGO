package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type LevoitAirpurifier struct {
	device_status string
	connection_status string
	active_time string
	filter_life string
	screen_status string
	mode string
	level string
	air_quality string
}

func (airpurifier *LevoitAirpurifier) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "device_status")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["device_status"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "connection_status")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["connection_status"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "active_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["active_time"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "filter_life")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["filter_life"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "screen_status")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["screen_status"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["mode"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["level"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "air_quality")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["air_quality"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (airpurifier *LevoitAirpurifier) SetDevice_status(value string) {

	command := exec.Command("python", "write.py", "device_status", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.device_status = value
}

func (airpurifier *LevoitAirpurifier) SetScreen_status(value string) {

	command := exec.Command("python", "write.py", "screen_status", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.screen_status = value
}

func (airpurifier *LevoitAirpurifier) SetMode(value string) {

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

func (airpurifier *LevoitAirpurifier) SetLevel(value string) {

	command := exec.Command("python", "write.py", "level", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	airpurifier.level = value
}

func NewDevice() *LevoitAirpurifier {
	airpurifier := &LevoitAirpurifier{
		device_status: "default",
		connection_status: "default",
		active_time: "default",
		filter_life: "default",
		screen_status: "default",
		mode: "default",
		level: "default",
		air_quality: "default",
	}
	return airpurifier
}

