package main

import (
	"log"
	"time"

	"github.com/drop-target-pinball/go-pinproc"
	"github.com/drop-target-pinball/go-pinproc/wpc"
)

func main() {
	pc, err := pinproc.NewWPC()
	if err != nil {
		log.Fatalf("unable to connect to P-ROC: %v", err)
	}
	pc.Reset(pinproc.ResetFlagUpdateDevice)

	for id, driver := range wpc.Devices {
		if pinproc.IsLamp(id) {
			if err := pc.DriverPatter(driver, 127, 127, 0, true); err != nil {
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
