package main

import (
	"log"
	"time"

	"github.com/drop-target-pinball/go-pinproc"
	"github.com/drop-target-pinball/go-pinproc/wpc"
)

func main() {
	pc, err := pinproc.New(wpc.MachType)
	if err != nil {
		log.Fatalf("unable to connect to P-ROC: %v", err)
	}
	pc.Reset(pinproc.ResetFlagUpdateDevice)

	for id, driver := range wpc.Devices {
		if pinproc.IsGI(id) {
			if err := pc.DriverPulse(driver, 0); err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := pc.FlushWriteData(); err != nil {
		log.Fatal(err)
	}

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
