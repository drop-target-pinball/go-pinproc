// Package proc is the go interface to the Pinball Remote Operations Controller
// (P-ROC) C library.
//
// See: https://www.multimorphic.com/store/circuit-boards/p-roc/
//
// See: https://github.com/preble/libpinproc
package pinproc

// #cgo LDFLAGS: -lpinproc -lftdi1 -lstdc++
// #include "pinproc.h"
import "C"
import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

const (
	MaxEvents   = 256
	MaxSwitches = 256
)

type MachType int

type DMDConfig struct {
	NumRows            uint8
	NumColumns         uint16
	NumSubFrames       uint8
	NumFrameBuffers    uint8
	AutoIncBufferWrPtr bool
	EnableFrameEvents  bool
	Enable             bool
	RClkLowCycles      [8]uint8
	LatchHighCycles    [8]uint8
	DeHighCycles       [8]uint16
	DotClkHalfPeriod   [8]uint8
}

func (d *DMDConfig) toC(cd *C.PRDMDConfig) {
	cd.numRows = C.uint8_t(d.NumRows)
	cd.numColumns = C.uint16_t(d.NumColumns)
	cd.numSubFrames = C.uint8_t(d.NumSubFrames)
	cd.numFrameBuffers = C.uint8_t(d.NumFrameBuffers)
	cd.autoIncBufferWrPtr = bool_t(d.AutoIncBufferWrPtr)
	cd.enableFrameEvents = bool_t(d.EnableFrameEvents)
	cd.enable = bool_t(d.Enable)

	for i := 0; i < 8; i++ {
		cd.rclkLowCycles[i] = C.uint8_t(d.RClkLowCycles[i])
		cd.latchHighCycles[i] = C.uint8_t(d.LatchHighCycles[i])
		cd.deHighCycles[i] = C.uint16_t(d.DeHighCycles[i])
		cd.dotclkHalfPeriod[i] = C.uint8_t(d.DotClkHalfPeriod[i])
	}
}

type DriverState struct {
	DriverNum            uint16
	OutputDriveTime      uint8
	Polarity             bool
	State                bool
	WaitForFirstTimeSlot bool
	Timeslots            uint32
	PatterOnTime         uint8
	PatterOffTime        uint8
	PatterEnable         bool
	FutureEnable         bool
}

func (d *DriverState) fromC(cd *C.PRDriverState) {
	d.DriverNum = uint16(cd.driverNum)
	d.OutputDriveTime = uint8(cd.outputDriveTime)
	d.Polarity = goBool(cd.polarity)
	d.State = goBool(cd.state)
	d.WaitForFirstTimeSlot = goBool(cd.waitForFirstTimeSlot)
	d.Timeslots = uint32(cd.timeslots)
	d.PatterOnTime = uint8(cd.patterOnTime)
	d.PatterOffTime = uint8(cd.patterOffTime)
	d.PatterEnable = goBool(cd.patterEnable)
	d.FutureEnable = goBool(cd.futureEnable)
}

func (d *DriverState) toC(cd *C.PRDriverState) {
	cd.driverNum = C.uint16_t(d.DriverNum)
	cd.outputDriveTime = C.uint8_t(d.OutputDriveTime)
	cd.polarity = bool_t(d.Polarity)
	cd.state = bool_t(d.State)
	cd.waitForFirstTimeSlot = bool_t(d.WaitForFirstTimeSlot)
	cd.timeslots = C.uint32_t(d.Timeslots)
	cd.patterOnTime = C.uint8_t(d.PatterOnTime)
	cd.patterOffTime = C.uint8_t(d.PatterOffTime)
	cd.patterEnable = bool_t(d.PatterEnable)
	cd.futureEnable = bool_t(d.FutureEnable)
}

func (d DriverState) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "driver %02d: enable=%v", d.DriverNum, d.State)
	if d.OutputDriveTime > 0 {
		fmt.Fprintf(&sb, ", time=%v", d.OutputDriveTime)
	}
	if d.PatterEnable {
		fmt.Fprintf(&sb, "patter=(%v %v)", d.PatterOnTime, d.PatterOffTime)
	}
	return sb.String()
}

type Event struct {
	EventType EventType
	Value     uint32
	Time      uint32
}

func (e *Event) fromC(ce *C.PREvent) {
	e.EventType = EventType(ce._type)
	e.Value = uint32(ce.value)
	e.Time = uint32(ce.time)
}

type EventType uint32

const (
	EventTypeInvalid                  EventType = 0
	EventTypeSwitchClosedDebounced    EventType = 1
	EventTypeSwitchOpenDebounced      EventType = 2
	EventTypeSwitchClosedNondebounced EventType = 3
	EventTypeSwitchOpenNondebounced   EventType = 4
	EventTypeDMDFrameDisplayed        EventType = 5
)

func (e EventType) String() string {
	switch e {
	case EventTypeInvalid:
		return "invalid"
	case EventTypeSwitchOpenDebounced:
		return "open"
	case EventTypeSwitchClosedDebounced:
		return "closed"
	case EventTypeSwitchOpenNondebounced:
		return "open (non-debounced)"
	case EventTypeSwitchClosedNondebounced:
		return "closed (non-debounced)"
	case EventTypeDMDFrameDisplayed:
		return "dmd frame displayed"
	}
	return "???"
}

type LogLevel int

const (
	LogVerbose LogLevel = iota
	LogInfo
	LogWarning
	LogError
	LogNone
)

type SwitchConfig struct {
	Clear                    bool
	HostEventsEnable         bool
	UseColumn9               bool
	UseColumn8               bool
	DirectMatrixScanLoopTime uint8
	PulsesBeforeCheckingRX   uint8
	InactivePulsesAfterBurst uint8
	PulsesPerBurst           uint8
	PulseHalfPeriodTIme      uint8
}

func (s *SwitchConfig) toC(cs *C.PRSwitchConfig) {
	cs.clear = bool_t(s.Clear)
	cs.hostEventsEnable = bool_t(s.HostEventsEnable)
	cs.use_column_9 = bool_t(s.UseColumn9)
	cs.use_column_8 = bool_t(s.UseColumn8)
	cs.directMatrixScanLoopTime = C.uint8_t(s.DirectMatrixScanLoopTime)
	cs.pulsesBeforeCheckingRX = C.uint8_t(s.PulsesBeforeCheckingRX)
	cs.inactivePulsesAfterBurst = C.uint8_t(s.InactivePulsesAfterBurst)
	cs.pulsesPerBurst = C.uint8_t(s.PulsesPerBurst)
	cs.pulseHalfPeriodTime = C.uint8_t(s.PulseHalfPeriodTIme)
}

type SwitchRule struct {
	ReloadActive bool
	NotifyHost   bool
}

func (s *SwitchRule) toC(cs *C.PRSwitchRule) {
	cs.reloadActive = bool_t(s.ReloadActive)
	cs.notifyHost = bool_t(s.NotifyHost)
}

type PROC struct {
	h        C.PRHandle
	prEvents []C.PREvent
}

func New(machType MachType) (*PROC, error) {
	handle := C.PRCreate(C.PRMachineType(machType))
	if uintptr(handle) == C.kPRHandleInvalid {
		return nil, procError()
	}
	return &PROC{
		h:        handle,
		prEvents: make([]C.PREvent, MaxEvents),
	}, nil
}

func (p *PROC) DMDDraw(dots []uint8) error {
	dotsPtr := (*C.uint8_t)(unsafe.Pointer(&dots[0]))
	result := C.PRDMDDraw(p.h, dotsPtr)
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) DMDUpdateConfig(dmdConfig DMDConfig) error {
	var prDmdConfig C.PRDMDConfig
	dmdConfig.toC(&prDmdConfig)
	result := C.PRDMDUpdateConfig(p.h, &prDmdConfig)
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) DriverGetState(driverNum uint8, driverState *DriverState) {
	var pcDriverState C.PRDriverState
	C.PRDriverGetState(p.h, C.uint8_t(driverNum), &pcDriverState)
	driverState.fromC(&pcDriverState)
}

func (p *PROC) DriverEnable(driverNum uint8) error {
	return p.DriverPulse(driverNum, 0)
}

func (p *PROC) DriverDisable(driverNum uint8) error {
	result := C.PRDriverDisable(p.h, C.uint8_t(driverNum))
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) DriverPatter(driverNum uint8, millisecondsOn uint8, millisecondsOff uint8, originalOnTime uint8, now bool) error {
	result := C.PRDriverPatter(p.h, C.uint8_t(driverNum), C.uint8_t(millisecondsOn), C.uint8_t(millisecondsOff), C.uint8_t(originalOnTime), bool_t(now))
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) DriverPulse(driverNum uint8, milliseconds uint8) error {
	result := C.PRDriverPulse(p.h, C.uint8_t(driverNum), C.uint8_t(milliseconds))
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) DriverSchedule(driverNum uint8, schedule uint32, cycleSeconds uint8, now bool) error {
	result := C.PRDriverSchedule(p.h, C.uint8_t(driverNum), C.uint32_t(schedule), C.uint8_t(cycleSeconds), bool_t(now))
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) DriverUpdateState(driverState *DriverState) error {
	var prDriverState C.PRDriverState
	driverState.toC(&prDriverState)
	result := C.PRDriverUpdateState(p.h, &prDriverState)
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) DriverWatchdogTickle() error {
	result := C.PRDriverWatchdogTickle(p.h)
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) FlushWriteData() error {
	result := C.PRFlushWriteData(p.h)
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) GetEvents(events []Event) (int, error) {
	bufSize := len(events)
	if bufSize > MaxEvents {
		bufSize = MaxEvents
	}
	n := C.PRGetEvents(p.h, &p.prEvents[0], C.int(bufSize))
	for i := 0; i < int(n); i++ {
		events[i].fromC(&p.prEvents[i])
	}
	return int(n), nil
}

func (p *PROC) Reset(resetFlags ResetFlag) {
	C.PRReset(p.h, C.uint32_t(resetFlags))
}

func (p *PROC) SwitchUpdateConfig(switchConfig SwitchConfig) error {
	var prSwitchConfig C.PRSwitchConfig
	switchConfig.toC(&prSwitchConfig)
	result := C.PRSwitchUpdateConfig(p.h, &prSwitchConfig)
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) SwitchUpdateRule(switchNum uint8, eventType EventType, rule SwitchRule, linkedDrivers []DriverState, driverOutputsNow bool) error {
	var prRule C.PRSwitchRule
	rule.toC(&prRule)

	if linkedDrivers == nil || len(linkedDrivers) == 0 {
		result := C.PRSwitchUpdateRule(p.h, C.uint8_t(switchNum), C.PREventType(eventType), &prRule, nil, 0, bool_t(driverOutputsNow))
		if int(result) != C.kPRSuccess {
			return procError()
		}
		return nil
	}

	linkedDriversLen := len(linkedDrivers)
	prLinkedDrivers := make([]C.PRDriverState, linkedDriversLen, linkedDriversLen)
	for i := 0; i < linkedDriversLen; i++ {
		linkedDrivers[i].toC(&prLinkedDrivers[i])
	}
	result := C.PRSwitchUpdateRule(p.h, C.uint8_t(switchNum), C.PREventType(eventType), &prRule, &prLinkedDrivers[0], C.int(linkedDriversLen), bool_t(driverOutputsNow))
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

func (p *PROC) GetSwitchStates(states []EventType) error {
	statesPtr := (*C.PREventType)(unsafe.Pointer(&states[0]))
	result := C.PRSwitchGetStates(p.h, statesPtr, C.ushort(len(states)))
	if int(result) != C.kPRSuccess {
		return procError()
	}
	return nil
}

type ResetFlag int

const (
	ResetFlagDefault ResetFlag = iota
	ResetFlagUpdateDevice
)

func DriverStateDisable(driverState *DriverState) {
	driverState.State = false
	driverState.Timeslots = 0
	driverState.WaitForFirstTimeSlot = false
	driverState.OutputDriveTime = 0
	driverState.PatterOnTime = 0
	driverState.PatterOffTime = 0
	driverState.PatterEnable = false
	driverState.FutureEnable = false
}

func DriverStatePulse(driverState *DriverState, milliseconds uint8) {
	driverState.State = true
	driverState.Timeslots = 0
	driverState.WaitForFirstTimeSlot = false
	driverState.OutputDriveTime = milliseconds
	driverState.PatterOnTime = 0
	driverState.PatterOffTime = 0
	driverState.PatterEnable = false
	driverState.FutureEnable = false
}

func DriverStateFuturePulse(driverState *DriverState, milliseconds uint8, futureTime uint32) {
	driverState.State = true
	driverState.Timeslots = futureTime
	driverState.WaitForFirstTimeSlot = false
	driverState.OutputDriveTime = milliseconds
	driverState.PatterOnTime = 0
	driverState.PatterOffTime = 0
	driverState.PatterEnable = false
	driverState.FutureEnable = true
}

func DriverStatePatter(driverState *DriverState, millisecondsOn uint8, millisecondsOff uint8, originalOnTime uint8, now bool) {
	driverState.State = true
	driverState.Timeslots = 0
	driverState.WaitForFirstTimeSlot = !now
	driverState.OutputDriveTime = originalOnTime
	driverState.PatterOnTime = millisecondsOn
	driverState.PatterOffTime = millisecondsOff
	driverState.PatterEnable = true
	driverState.FutureEnable = false
}

func DriverStatePulsedPatter(driverState *DriverState, millisecondsOn uint8, millisecondsOff uint8, patterTime uint8, now bool) {
	driverState.State = false
	driverState.Timeslots = 0
	driverState.WaitForFirstTimeSlot = !now
	driverState.OutputDriveTime = patterTime
	driverState.PatterOnTime = millisecondsOn
	driverState.PatterOffTime = millisecondsOff
	driverState.PatterEnable = true
	driverState.FutureEnable = false
}

func IsCoil(id string) bool {
	return id[0] == 'C'
}

func IsGI(id string) bool {
	return id[0] == 'G'
}

func IsLamp(id string) bool {
	return id[0] == 'L'
}

func IsSwitch(id string) bool {
	return id[0] == 'S'
}

func LogSetLevel(level LogLevel) {
	C.PRLogSetLevel(C.PRLogLevel(level))
}

func bool_t(b bool) C.int32_t {
	if b {
		return 1
	}
	return 0
}

func goBool(i C.int) bool {
	if i == 0 {
		return false
	}
	return true
}

func procError() error {
	msg := C.GoString(C.PRGetLastErrorText())
	return errors.New(msg)
}

func init() {
	LogSetLevel(LogNone)
}
