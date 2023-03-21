package main

import (
	"fmt"
	"log"
	"time"

	"github.com/distatus/battery"
)

func main() {
	for {
		batteryInfo, err := battery.Get(0)
		if err != nil {
			log.Fatalf("Error getting battery info: %v", err)
		}

		printBatteryInfo(batteryInfo)
		time.Sleep(5 * time.Second)
	}
}

func printBatteryInfo(batteryInfo *battery.Battery) { // Updated the parameter type
	percentage := batteryInfo.Current / batteryInfo.Full * 100
	state := getBatteryState(batteryInfo.State)

	fmt.Printf("Battery Usage: %.2f%% (%s)\n", percentage, state)
}

func getBatteryState(state battery.State) string {
	switch state {
	case battery.Charging:
		return "Charging"
	case battery.Discharging:
		return "Discharging"
	case battery.Unknown: // Changed from battery.NotCharging to battery.Unknown
		return "Not Charging"
	case battery.Full:
		return "Full"
	default:
		return "Unknown"
	}
}
