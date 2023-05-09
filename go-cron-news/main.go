package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func cronFunc() {
	fmt.Println("hello crons")
}
func main() {
	c := cron.New()
	c.AddFunc("@every 10s", cronFunc)
	c.Start()

	var a int
	fmt.Scan(&a)
	fmt.Println(a)
}
