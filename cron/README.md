# Index

- 安装
- 规则说明
  - 表达格式
  - 特殊字符
  - 预定义的 Cron 时间表
- 基本使用

# 安装

- `go get github.com/robfig/cron/v3@v3.0.0`

# 规则说明

## 表达格式

| 字段                           | 是否必填 | 值范围          | 允许的特殊字符 |
| ------------------------------ | -------- | --------------- | -------------- |
| 秒（Seconds）                  | Yes      | 0-59            | \* / , -       |
| 分（Minutes）                  | Yes      | 0-59            | \* / , -       |
| 时（Hours）                    | Yes      | 0-23            | \* / , -       |
| 一个月中的某天（Day of month） | Yes      | 1-31            | \* / , - ?     |
| 月（Month）                    | Yes      | 1-12 or JAN-DEC | \* / , -       |
| 星期几（Day of week）          | Yes      | 0-6 or SUN-SAT  | \* / , - ?     |

## 特殊字符

| 字符         | 说明                                                                                                                                                                                   |
| ------------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 星号 ( \* )  | 星号表示将匹配字段的所有值                                                                                                                                                             |
| 斜线 ( / )   | 斜线用户 描述范围的增量，表现为 “N-MAX/x”，first-last/x 的形式，例如 3-59/15 表示此时的第三分钟和此后的每 15 分钟，到 59 分钟为止。即从 N 开始，使用增量直到该特定范围结束。它不会重复 |
| 逗号 ( , )   | 逗号用于分隔列表中的项目。例如，在 Day of week 使用“MON，WED，FRI”将意味着星期一，星期三和星期五                                                                                       |
| 连字符 ( - ) | 连字符用于定义范围。例如，9 - 17 表示从上午 9 点到下午 5 点的每个小时                                                                                                                  |
| 问号 ( ? )   | 不指定值，用于代替 “ \* ”，类似 “ \_ ” 的存在                                                                                                                                          |

## 预定义的 Cron 时间表

| 输入                   | 简述                                   | 相当于       |
| ---------------------- | -------------------------------------- | ------------ |
| @yearly (or @annually) | 1 月 1 日午夜运行一次                  | 0 0 0 1 1 \_ |
| @monthly               | 每个月的午夜，每个月的第一个月运行一次 | 0 0 0 1      |
| @weekly                | 每周一次，周日午夜运行一次             | 0 0 0 0      |
| @daily (or @midnight)  | 每天午夜运行一次                       | 0 0 0 \_     |
| @hourly                | 每小时运行一次                         | 0 0          |

# 基本使用

```
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
```
