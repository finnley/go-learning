package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	arr := make([]string, 1, 10000000)
	arr[0] = "a"
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			log.Println(1)
		}
	}()

	http.ListenAndServe(":39090", nil)
}

/**
根据 go tool 工访问pprof特定端口路径下的使用情况
比如采集5秒内的CPU使用情况：
go tool pprof http://localhost:39090/debug/pprof/profile?secones=5s
➜  ~ go tool pprof http://localhost:39090/debug/pprof/profile\?secones\=5s
Fetching profile over HTTP from http://localhost:39090/debug/pprof/profile?secones=5s
Saved profile in /Users/finnley/pprof/pprof.samples.cpu.001.pb.gz
Type: cpu
Time: Aug 15, 2022 at 7:45am (CST)
Duration: 30s, Total samples = 10ms (0.033%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)

使用top查看从高到底的使用情况
(pprof) top
Showing nodes accounting for 10ms, 100% of 10ms total
      flat  flat%   sum%        cum   cum%
      10ms   100%   100%       10ms   100%  runtime.pthread_cond_wait
         0     0%   100%       10ms   100%  runtime.findrunnable
         0     0%   100%       10ms   100%  runtime.mPark (inline)
         0     0%   100%       10ms   100%  runtime.mcall
         0     0%   100%       10ms   100%  runtime.notesleep
         0     0%   100%       10ms   100%  runtime.park_m
         0     0%   100%       10ms   100%  runtime.schedule
         0     0%   100%       10ms   100%  runtime.semasleep
         0     0%   100%       10ms   100%  runtime.stopm
(pprof)

查看内存的使用情况
go tool pprof http://localhost:39090/debug/pprof/heap
➜  ~ go tool pprof http://localhost:39090/debug/pprof/heap
Fetching profile over HTTP from http://localhost:39090/debug/pprof/heap
Saved profile in /Users/finnley/pprof/pprof.alloc_objects.alloc_space.inuse_objects.inuse_space.001.pb.gz
Type: inuse_space
Time: Aug 15, 2022 at 7:50am (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 5126.25kB, 100% of 5126.25kB total
Showing top 10 nodes out of 28
      flat  flat%   sum%        cum   cum%
 3075.38kB 59.99% 59.99%  3075.38kB 59.99%  runtime.allocm
  513.31kB 10.01% 70.01%   513.31kB 10.01%  regexp/syntax.(*compiler).inst (inline)
  513.31kB 10.01% 80.02%   513.31kB 10.01%  vendor/golang.org/x/net/http2/hpack.(*headerFieldTable).addEntry (inline)
  512.20kB  9.99% 90.01%   512.20kB  9.99%  runtime.malg
  512.05kB  9.99%   100%  1538.67kB 30.02%  runtime.main
         0     0%   100%   513.31kB 10.01%  internal/profile.init
         0     0%   100%   513.31kB 10.01%  regexp.Compile (inline)
         0     0%   100%   513.31kB 10.01%  regexp.MustCompile
         0     0%   100%   513.31kB 10.01%  regexp.compile
         0     0%   100%   513.31kB 10.01%  regexp/syntax.(*compiler).cap (inline)
(pprof)
*/
