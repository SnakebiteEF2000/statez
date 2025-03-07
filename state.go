package statez

type ServiceState int32

const (
	StateIgnore ServiceState = iota - 1
	StateNotReady
	StateReady
)

var readyStateName = map[ServiceState]string{
	StateIgnore:   "ignored",
	StateNotReady: "not ready",
	StateReady:    "ready",
}

func (ss ServiceState) String() string {
	return readyStateName[ss]
}
