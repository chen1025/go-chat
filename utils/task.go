package utils

import "time"

type TaskFunc func(interface{}) bool

// DoTask
// delay 首次执行延迟时间
// tick 间隔时间
// fun 执行方法
// 方法参数
func DoTask(delay, tick time.Duration, fun TaskFunc, params interface{}) {
	go func() {
		if fun == nil {
			return
		}
		t := time.NewTimer(delay)
		for {
			select {
			case <-t.C:
				if fun(params) == false {
					return
				}
				t.Reset(tick)
			}
		}
	}()
}
