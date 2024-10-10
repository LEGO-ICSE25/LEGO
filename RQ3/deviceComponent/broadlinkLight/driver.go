package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type BroadlinkLight struct {
	power string
	brightness string
	red string
	blue string
	green string
	colortemp string
	hue string
	saturation string
	transitionduration string
	maxworktime string
	bulb_colormode string
	bulb_scenes string
	bulb_scene string
}

func (light *BroadlinkLight) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "power")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "brightness")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["brightness"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "red")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["red"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "blue")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["blue"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "green")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["green"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "colortemp")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["colortemp"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "hue")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["hue"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "saturation")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["saturation"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "transitionduration")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["transitionduration"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "maxworktime")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["maxworktime"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "bulb_colormode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["bulb_colormode"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "bulb_scenes")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["bulb_scenes"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "bulb_scene")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["bulb_scene"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (light *BroadlinkLight) SetPower(value string) {

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

func (light *BroadlinkLight) SetBrightness(value string) {

	command := exec.Command("python", "write.py", "brightness", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.brightness = value
}

func (light *BroadlinkLight) SetRed(value string) {

	command := exec.Command("python", "write.py", "red", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.red = value
}

func (light *BroadlinkLight) SetBlue(value string) {

	command := exec.Command("python", "write.py", "blue", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.blue = value
}

func (light *BroadlinkLight) SetGreen(value string) {

	command := exec.Command("python", "write.py", "green", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.green = value
}

func (light *BroadlinkLight) SetColortemp(value string) {

	command := exec.Command("python", "write.py", "colortemp", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.colortemp = value
}

func (light *BroadlinkLight) SetHue(value string) {

	command := exec.Command("python", "write.py", "hue", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.hue = value
}

func (light *BroadlinkLight) SetSaturation(value string) {

	command := exec.Command("python", "write.py", "saturation", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.saturation = value
}

func (light *BroadlinkLight) SetTransitionduration(value string) {

	command := exec.Command("python", "write.py", "transitionduration", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.transitionduration = value
}

func (light *BroadlinkLight) SetMaxworktime(value string) {

	command := exec.Command("python", "write.py", "maxworktime", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.maxworktime = value
}

func NewDevice() *BroadlinkLight {
	light := &BroadlinkLight{
		power: "default",
		brightness: "default",
		red: "default",
		blue: "default",
		green: "default",
		colortemp: "default",
		hue: "default",
		saturation: "default",
		transitionduration: "default",
		maxworktime: "default",
		bulb_colormode: "default",
		bulb_scenes: "default",
		bulb_scene: "default",
	}
	return light
}

