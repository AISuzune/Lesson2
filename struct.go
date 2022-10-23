package main

import (
	"fmt"
	"time"
)

type Question struct {
	Id         int64
	Mailbox    int64
	Heat       int64
	content    string
	Maintainer string
	State      bool
	HasAnswer  bool
	CreatedAt  time.Time
	DeletedAt  time.Time
	UpdateAt   time.Time
}

func main() {
	var q Question
	fmt.Printf("q=%#v\n", q)
}
