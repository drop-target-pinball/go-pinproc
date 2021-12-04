package main

import (
	"fmt"
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
	if err := pc.SwitchUpdateConfig(wpc.SwitchConfigDefault); err != nil {
		log.Fatal(err)
	}

	rule := pinproc.SwitchRule{NotifyHost: true}
	for id, device := range wpc.Devices {
		if !pinproc.IsSwitch(id) {
			continue
		}
		if err := pc.SwitchUpdateRule(device, pinproc.EventTypeSwitchClosedDebounced, rule, nil, false); err != nil {
			log.Fatal(err)
		}
		if err := pc.SwitchUpdateRule(device, pinproc.EventTypeSwitchOpenDebounced, rule, nil, false); err != nil {
			log.Fatal(err)
		}
	}
	if err := pc.FlushWriteData(); err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	events := make([]pinproc.Event, pinproc.MaxEvents)
	fmt.Println("listening for switch events")
	for {
		<-ticker.C
		n, err := pc.GetEvents(events)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < n; i++ {
			e := events[i]
			ident := wpc.SwitchNames[uint8(e.Value)]
			fmt.Printf("%v %v\n", ident, e.EventType)
		}
		if err := pc.DriverWatchdogTickle(); err != nil {
			log.Fatal(err)
		}
		if err := pc.FlushWriteData(); err != nil {
			log.Fatal(err)
		}
	}
}
