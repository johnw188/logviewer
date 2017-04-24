package main

import (
	"github.com/johnw188/logviewer"
	"time"
)

func main() {
	v := logviewer.NewLogViewer()
	v.AddLogFeed("Pod 1", 100)
	v.AddLogFeed("Pod 2", 100)
	v.AddLogFeed("Pod 3 really long", 100)

	t0 := time.NewTicker(time.Second * 1)
	t1 := time.NewTicker(time.Second * 3)
	t2 := time.NewTicker(time.Second * 5)

	go func(){
		for {
			select {
			case t := <-t0.C:
				v.AddLogLine(&logviewer.LogLine{Log: t.String() + ": log1"}, 0)
			case t := <-t1.C:
				v.AddLogLine(&logviewer.LogLine{Log: t.String() + ": log2"}, 1)
			case t := <-t2.C:
				v.AddLogLine(&logviewer.LogLine{Log: t.String() + ": log3"}, 2)
			}
		}
	}()

	v.Display()
}
