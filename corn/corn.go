package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	i := 0
	c := cron.New()
	spec := "0 */10 * * * *"
	_ = c.AddFunc(spec, func() {
		i++
		fmt.Println("cron running:", i,time.Now())

	})
	c.Start()

	select {}

}
