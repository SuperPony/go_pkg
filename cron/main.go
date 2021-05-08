package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Job struct {
}

func (j *Job) Run() {
	fmt.Println("Run")
}

func main() {
	example()
	// jobExample()
}

// Cron.AddJob(spec string, cmd cron.Job) 更适合一些复杂的场景， Job 是一个接口，包含 Run 方法
func jobExample() {
	c := cron.New()
	if _, err := c.AddJob("* * * * ?", &Job{}); err != nil {
		log.Fatalln(err)
	}
	c.Start()
	defer c.Stop()

	timer := time.After(time.Minute)
	<-timer
	fmt.Println("is over")
}

// 基础示范
func example() {
	c := cron.New()

	task := func() {
		fmt.Println("task run")
	}

	if _, err := c.AddFunc("* * * * ?", task); err != nil {
		log.Fatalln(err)
	}

	// 另起一个 goroutine 开始执行
	c.Start()
	// 停止正在运行的任务，否则不做任何操作
	defer c.Stop()

	// 让程序进入阻塞
	select {}
}
