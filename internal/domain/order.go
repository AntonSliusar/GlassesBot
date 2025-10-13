package domain

import (
	"time"
)

const (
	STATUS_IN_WORK = "inWork"
	STATUS_PAUSER = "paused"
)

type Order struct {
	Frame         string
	Lenses        string
	Date          time.Time
	WorkingTime   time.Duration
	LastStartTime time.Time
	Status        string
}

func NewOrder() *Order {
	return &Order{
		Frame:   "",
		Lenses: "",
		Date:    time.Now(),
		WorkingTime: 0,
		LastStartTime: time.Now(),
		Status: "InWork",
	}
}	
func (o *Order) Pause() {
	o.Status = STATUS_PAUSER
	o.TotalWorkinTime()
}

func (o *Order) Resume() {
	o.Status = STATUS_IN_WORK
	o.LastStartTime = time.Now()
}

func (o *Order) TotalWorkinTime() {
	o.WorkingTime += time.Since(o.LastStartTime)
}