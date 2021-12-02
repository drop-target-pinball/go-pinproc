package main

import (
	"log"
	"time"

	"github.com/drop-target-pinball/go-pinproc"
	"github.com/drop-target-pinball/go-pinproc/wpc"
)

func main() {
	log.SetFlags(0)

	pc, err := pinproc.New(wpc.MachType)
	if err != nil {
		log.Fatalf("unable to connect to P-ROC: %v", err)
	}
	pc.Reset(pinproc.ResetFlagUpdateDevice)

	dmdConf := wpc.DMDConfigDefault
	if err := pc.DMDUpdateConfig(dmdConf); err != nil {
		log.Fatal(err)
	}

	subFrameSize := int(dmdConf.NumColumns) * int(dmdConf.NumRows) / 8
	frameSize := subFrameSize * int(dmdConf.NumSubFrames)

	dots := make([]uint8, frameSize, frameSize)
	c := uint8(0xff)
	i := 0

	ticker := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			dots[(0*subFrameSize)+i] = c
			dots[(1*subFrameSize)+i] = c
			dots[(2*subFrameSize)+i] = c
			dots[(3*subFrameSize)+i] = c
			i++
			if i >= subFrameSize {
				i = 0
				if c == 0xff {
					c = 0x00
				} else {
					c = 0xff
				}
			}
			if err := pc.DMDDraw(dots); err != nil {
				log.Fatal(err)
			}
			if err := pc.FlushWriteData(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
