package main

import (
	"fmt"
	"log"
	"time"

	"github.com/drop-target-pinball/go-pinproc"
	"github.com/drop-target-pinball/go-pinproc/wpc"
)

func main() {
	log.SetFlags(0)

	pc, err := pinproc.NewWPC()
	if err != nil {
		log.Fatalf("unable to connect to P-ROC: %v", err)
	}
	pc.Reset(pinproc.ResetFlagUpdateDevice)

	var (
		leftMainOn,
		leftHoldOn,
		leftHoldOff,
		rightMainOn,
		rightHoldOn,
		rightHoldOff pinproc.DriverState
	)

	pc.DriverGetState(wpc.FLLM, &leftMainOn)
	pc.DriverGetState(wpc.FLLH, &leftHoldOn)
	pc.DriverGetState(wpc.FLLH, &leftHoldOff)
	pc.DriverGetState(wpc.FLRM, &rightMainOn)
	pc.DriverGetState(wpc.FLRH, &rightHoldOn)
	pc.DriverGetState(wpc.FLRH, &rightHoldOff)

	pinproc.DriverStatePulse(&leftMainOn, 25)
	pinproc.DriverStatePulse(&leftHoldOn, 0)
	pinproc.DriverStateDisable(&leftHoldOff)
	pinproc.DriverStatePulse(&rightMainOn, 25)
	pinproc.DriverStatePulse(&rightHoldOn, 0)
	pinproc.DriverStateDisable(&leftHoldOff)

	leftDriversOn := []pinproc.DriverState{leftMainOn, leftHoldOn}
	leftDriversOff := []pinproc.DriverState{leftHoldOff}
	rightDriversOn := []pinproc.DriverState{rightMainOn, rightHoldOn}
	rightDriversOff := []pinproc.DriverState{rightHoldOff}

	// Flipper buttons are optos so open == button pressed
	err = pc.SwitchUpdateRule(
		wpc.SF4,
		pinproc.EventTypeSwitchOpenNondebounced,
		pinproc.SwitchRule{},
		leftDriversOn,
		true)
	if err != nil {
		log.Fatal(err)
	}
	err = pc.SwitchUpdateRule(
		wpc.SF4,
		pinproc.EventTypeSwitchClosedNondebounced,
		pinproc.SwitchRule{},
		leftDriversOff,
		true)
	if err != nil {
		log.Fatal(err)
	}
	err = pc.SwitchUpdateRule(
		wpc.SF2,
		pinproc.EventTypeSwitchOpenNondebounced,
		pinproc.SwitchRule{},
		rightDriversOn,
		true)
	if err != nil {
		log.Fatal(err)
	}
	err = pc.SwitchUpdateRule(
		wpc.SF2,
		pinproc.EventTypeSwitchClosedNondebounced,
		pinproc.SwitchRule{},
		rightDriversOff,
		true)
	if err != nil {
		log.Fatal(err)
	}

	if err := pc.FlushWriteData(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("press flipper buttons")
	ticker := time.NewTicker(250 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			if err := pc.DriverWatchdogTickle(); err != nil {
				log.Fatal(err)
			}
			if err := pc.FlushWriteData(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
