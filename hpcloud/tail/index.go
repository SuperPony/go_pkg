package main

import (
	"log"

	"github.com/hpcloud/tail"
)

func main() {
	config := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件那个位置开始读取
		MustExist: true,                                 // 文件必须存在，否则报错退出
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 跟随文件
		Poll:      true,                                 // 轮训文件更改
	}

	t, err := tail.TailFile("./my.log", config)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 关闭 chan
	// ch := time.After(time.Second * 10)
	// go func(ch <-chan time.Time, t *tail.Tail) {
	// 	select {
	// 	case <-ch:
	// 		t.Stop()
	// 	}
	// }(ch, t)

	for line := range t.Lines {
		/*
			type Line struct {
				Text string
				Time time.Time
				Err  error // Error from tail
			}
		*/
		log.Println(line)
	}
}
