package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	t time.Time
}

func (t *Time) UnmarshalJSON(s []byte) error {
	var (
		year int
		mon  int
		mday int
		hour int
		min  int
		sec  int
	)
	if len(s) <= 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return fmt.Errorf("invalid time: %s", s)
	}
	fmt.Println(s)
	var str = string(s[1 : len(s)-1])
	if n, err := fmt.Sscanf(str, "%d-%02d-%02d %02d:%02d:%02d", &year, &mon, &mday, &hour, &min, &sec); err != nil {
		return fmt.Errorf("invalid string(%s): %s", err.Error(), s)
	} else if n != 6 {
		return fmt.Errorf("invalid time: %s", s)
	}
	t.t = time.Date(year, time.Month(mon), mday, hour, min, sec, 0, time.Local)
	return nil
}

func main() {
	var txt = `{"time": "2015-07-03 21:39:00"}`
	var t struct {
		Time Time `json:"time"`
	}
	if err := json.Unmarshal([]byte(txt), &t); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t.Time.t)
}
