package structs

import "time"

type Data struct {
	Time time.Time `json:"time"`
	Msg  string    `json:"msg"`
}
