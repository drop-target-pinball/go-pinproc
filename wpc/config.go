package wpc

import "github.com/drop-target-pinball/go-pinproc"

var DMDConfigDefault = pinproc.DMDConfig{
	NumRows:            32,
	NumColumns:         128,
	NumSubFrames:       4,
	NumFrameBuffers:    3,
	AutoIncBufferWrPtr: true,
	RClkLowCycles: [8]uint8{
		15, 15, 15, 15,
		0, 0, 0, 0,
	},
	LatchHighCycles: [8]uint8{
		15, 15, 15, 15,
		0, 0, 0, 0,
	},
	DotClkHalfPeriod: [8]uint8{
		1, 1, 1, 1,
		0, 0, 0, 0,
	},
	DeHighCycles: [8]uint16{
		90, 190, 50, 377,
		0, 0, 0, 0,
	},
}

var SwitchConfigDefault = pinproc.SwitchConfig{
	Clear:                    false,
	UseColumn8:               true,
	UseColumn9:               false,
	HostEventsEnable:         true,
	DirectMatrixScanLoopTime: 2,
	PulsesBeforeCheckingRX:   10,
	InactivePulsesAfterBurst: 12,
	PulsesPerBurst:           6,
	PulseHalfPeriodTIme:      13,
}
