package repository

import "time"

type Status struct {
	Code int
	Text string
}

type Data struct {
	RID         int64
	UnixTime    int
	LocalTime   time.Time //for view
	LogType     string
	LogSeverity int
	LogText     string

	Status Status
}

type Record interface {
	Get(id int64) (*Data, error)
	Insert(data Data) (int64, error)
	Tail(limit int64) (*[]Data, error)
}
