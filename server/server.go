package server

import (
	pb "../daysproto"
	"fmt"
	"golang.org/x/net/context"
)

//A  receiver
type A struct {
}

//GetWeekDay gets a day whether it's weekday or not
func (a *A) GetWeekDay(ctx context.Context, day *pb.Day) (*pb.DayResponse, error) {
	var dr = &pb.DayResponse{}
	var err error
	if day.Day == 0 || day.Day == 1 || day.Day == 2 || day.Day == 3 || day.Day == 4 {
		dr.IsWeekDay = true
	} else {
		if day.Day == 5 || day.Day == 6 {
			dr.IsWeekDay = false
		} else {
			dr.IsWeekDay = false
			err = fmt.Errorf("something's wrong with day: %v\n", *day)
		}
	}
	return dr, err
}

//NewServer starts a server
func NewServer() *A {
	a := new(A)
	return a
}
