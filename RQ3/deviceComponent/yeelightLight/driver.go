package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type YeelightLight struct {
	power string
	brightness string
	color string
	mode string
	developer_mode string
	cfg_save_state string
	name string
	delay string
	music_on string
	active_mode string
}

func (light *YeelightLight) GetStates() map[string]string {

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

	command = exec.Command("python", "read.py", "developer_mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["developer_mode"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "cfg_save_state")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["cfg_save_state"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "name")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["name"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "delay")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["delay"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "music_on")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["music_on"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "active_mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["active_mode"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (light *YeelightLight) SetPower(value string) {

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

func (light *YeelightLight) SetBrightness(value string) {

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

func (light *YeelightLight) SetColor(value string) {

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

func (light *YeelightLight) SetMode(value string) {

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

func (light *YeelightLight) SetDeveloper_mode(value string) {

	command := exec.Command("python", "write.py", "developer_mode", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.developer_mode = value
}

func (light *YeelightLight) SetCfg_save_state(value string) {

	command := exec.Command("python", "write.py", "cfg_save_state", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.cfg_save_state = value
}

func (light *YeelightLight) SetName(value string) {

	command := exec.Command("python", "write.py", "name", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.name = value
}

func (light *YeelightLight) SetDelay(value string) {

	command := exec.Command("python", "write.py", "delay", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.delay = value
}

func (light *YeelightLight) SetMusic_on(value string) {

	command := exec.Command("python", "write.py", "music_on", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.music_on = value
}

func (light *YeelightLight) SetActive_mode(value string) {

	command := exec.Command("python", "write.py", "active_mode", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	light.active_mode = value
}

func NewDevice() *YeelightLight {
	light := &YeelightLight{
		power: "default",
		brightness: "default",
		color: "default",
		mode: "default",
		developer_mode: "default",
		cfg_save_state: "default",
		name: "default",
		delay: "default",
		music_on: "default",
		active_mode: "default",
	}
	return light
}

