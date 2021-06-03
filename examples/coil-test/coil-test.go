package main

import (
	"log"
	"os"

	"github.com/drop-target-pinball/go-pinproc"
	"github.com/drop-target-pinball/go-pinproc/wpc"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 2 {
		log.Fatal("expected one argument: coil identifier")
	}
	id := os.Args[1]
	device, ok := wpc.Devices[id]
	if !ok || !pinproc.IsCoil(id) {
		log.Fatal("invalid coil identifier")
	}

	pc, err := pinproc.NewWPC()
	if err != nil {
		log.Fatalf("unable to connect to P-ROC: %v", err)
	}
	pc.Reset(pinproc.ResetFlagUpdateDevice)

	if err := pc.DriverPulse(device, 20); err != nil {
		log.Fatal(err)
	}

	if err := pc.FlushWriteData(); err != nil {
		log.Fatal(err)
	}
}
