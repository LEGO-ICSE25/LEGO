package main

import (
	"os"
	"os/exec"
	"strings"
	"k8s.io/klog/v2"
)

type CgllcAirqualitymonitor struct {
	humidity string
	pm25 string
	pm10 string
	temperature string
	co2 string
	battery string
	charging_state string
	monitoring_frequency string
	screen_off string
	device_off string
	temperature_unit string
}

func (monitor *CgllcAirqualitymonitor) GetStates() map[string]string {

	properties := make(map[string]string)

	command := exec.Command("python", "read.py", "humidity")
	output, err := command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["humidity"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "pm25")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["pm25"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "pm10")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["pm10"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "temperature")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["temperature"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "co2")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["co2"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "battery")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["battery"] = strings.TrimRight(string(output), "\n")
	
	command = exec.Command("python", "read.py", "charging_state")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["charging_state"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "monitoring_frequency")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["monitoring_frequency"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "screen_off")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["screen_off"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "device_off")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["device_off"] = strings.TrimRight(string(output), "\n")

	command = exec.Command("python", "read.py", "temperature_unit")
	output, err = command.Output()
	if err != nil {
		klog.Infoln(err)
	}
	properties["temperature_unit"] = strings.TrimRight(string(output), "\n")

	return properties
}

func (monitor *CgllcAirqualitymonitor) SetMonitoring_frequency(value string) {

	command := exec.Command("python", "write.py", "monitoring_frequency", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	monitor.monitoring_frequency = value
}

func (monitor *CgllcAirqualitymonitor) SetScreen_off(value string) {

	command := exec.Command("python", "write.py", "screen_off", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	monitor.screen_off = value
}

func (monitor *CgllcAirqualitymonitor) SetDevice_off(value string) {

	command := exec.Command("python", "write.py", "device_off", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	monitor.device_off = value
}

func (monitor *CgllcAirqualitymonitor) SetTemperature_unit(value string) {

	command := exec.Command("python", "write.py", "temperature_unit", value)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		klog.Infoln("Fail to execute script.\n")
		return
	}

	monitor.temperature_unit = value
}

func NewDevice() *CgllcAirqualitymonitor {
	monitor := &CgllcAirqualitymonitor{
		humidity: "default",
		pm25: "default",
		pm10: "default",
		temperature: "default",
		co2: "default",
		battery: "default",
		charging_state: "default",
		monitoring_frequency: "default",
		screen_off: "default",
		device_off: "default",
		temperature_unit: "default",
	}
	return monitor
}
