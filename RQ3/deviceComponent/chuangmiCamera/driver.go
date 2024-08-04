package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type ChuangmiCamera struct {
	power string
	motion_record string
	light string
	full_color string
	flip string
	improve_program string
	wdr string
	track string
	sdcard_status string
	watermark string
	max_client string
	night_mode string
	mini_level string
}

func (camera *ChuangmiCamera) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "power")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["power"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "motion_record")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["motion_record"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "light")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["light"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "full_color")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["full_color"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "flip")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["flip"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "improve_program")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["improve_program"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "wdr")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["wdr"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "track")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["track"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "sdcard_status")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["sdcard_status"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "watermark")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["watermark"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "max_client")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["max_client"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "night_mode")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["night_mode"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "mini_level")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["mini_level"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (camera *ChuangmiCamera) SetPower(value string) {

	command := exec.Command("python", "write.py", "power", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.power = value
}

func (camera *ChuangmiCamera) SetMotion_record(value string) {

	command := exec.Command("python", "write.py", "motion_record", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.motion_record = value
}

func (camera *ChuangmiCamera) SetLight(value string) {

	command := exec.Command("python", "write.py", "light", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.light = value
}

func (camera *ChuangmiCamera) SetFull_color(value string) {

	command := exec.Command("python", "write.py", "full_color", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.full_color = value
}

func (camera *ChuangmiCamera) SetFlip(value string) {

	command := exec.Command("python", "write.py", "flip", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.flip = value
}

func (camera *ChuangmiCamera) SetImprove_program(value string) {

	command := exec.Command("python", "write.py", "improve_program", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.improve_program = value
}

func (camera *ChuangmiCamera) SetWdr(value string) {

	command := exec.Command("python", "write.py", "wdr", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.wdr = value
}

func (camera *ChuangmiCamera) SetWatermark(value string) {

	command := exec.Command("python", "write.py", "watermark", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.watermark = value
}

func (camera *ChuangmiCamera) SetNight_mode(value string) {

	command := exec.Command("python", "write.py", "night_mode", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	camera.night_mode = value
}

func NewDevice() *ChuangmiCamera {
	camera := &ChuangmiCamera{
		power: "default",
		motion_record: "default",
		light: "default",
		full_color: "default",
		flip: "default",
		improve_program: "default",
		wdr: "default",
		track: "default",
		sdcard_status: "default",
		watermark: "default",
		max_client: "default",
		night_mode: "default",
		mini_level: "default",	
	}
	return camera
}
