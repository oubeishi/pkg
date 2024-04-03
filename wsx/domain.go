package wsx

type Event struct {
	Event string `ws:"event" json:"event"`
}
