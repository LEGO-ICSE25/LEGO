package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type MystromLight struct {
	power string
	consumption string
	color string
	mode string
	firmware string
	transition_time string
	state string
	bulb_type string
}

func (light *MystromLight) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "power")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "consumption")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["consumption"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "color")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["color"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["mode"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "firmware")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["firmware"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "transition_time")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["transition_time"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "state")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["state"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "bulb_type")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["bulb_type"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (light *MystromLight) SetPower(value string) {

	command := exec.Command("python", "write.py", "power", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.power = value
}

func (light *MystromLight) SetConsumption(value string) {

	command := exec.Command("python", "write.py", "consumption", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.consumption = value
}

func (light *MystromLight) SetColor(value string) {

	command := exec.Command("python", "write.py", "color", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.color = value
}

func (light *MystromLight) SetMode(value string) {

	command := exec.Command("python", "write.py", "mode", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.mode = value
}

func (light *MystromLight) SetTransition_time(value string) {

	command := exec.Command("python", "write.py", "transition_time", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.transition_time = value
}

func (light *MystromLight) SetState(value string) {

	command := exec.Command("python", "write.py", "state", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.state = value
}

func NewDevice() *MystromLight {
	light := &MystromLight{
		power: "default",
		consumption: "default",
		color: "default",
		mode: "default",
		firmware: "default",
		transition_time: "default",
		state: "default",
		bulb_type: "default",
	}
	return light
}

