// tail 库以常驻模式进行日志追踪
package tail

import (
	"log"

	"github.com/hpcloud/tail"
)

var config = tail.Config{
	Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件那个位置开始读取
	MustExist: true,                                 // 文件必须存在，否则报错退出
	ReOpen:    true,                                 // 重新打开
	Follow:    true,                                 // 跟随文件
	Poll:      false,                                // 以轮训文件更改，如果 false 则使用 inotify
}

func Read() {

	t, err := tail.TailFile("my.log", config)
	if err != nil {
		log.Fatalln(err)
		return
	}

	// 关闭文件 t.Cleanup
	defer func() {
		t.Cleanup()
	}()

	// 停止追踪 t.Stop()，此处模拟 10秒 关闭追踪
	// ch := time.After(time.Second * 10)
	// go func(ch <-chan time.Time, t *tail.Tail) {
	// 	select {
	// 	case <-ch:
	// 		t.Stop()
	// 	}
	// }(ch, t)

	// 当有新的日志行写入时，写入 chan *tail.Line
	// Lines chan *tail.Line
	for line := range t.Lines {
		/*
			type Line struct {
				Text string 内容
				Time time.Time 写入时间
				Err  error // Error from tail
			}
		*/
		log.Println(line)
	}
}
