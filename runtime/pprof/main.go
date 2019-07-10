package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

func f1()  {
	f2()
}
func f2()  {
	f3()
}
func f3()  {
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("f3")
	}
}
// Dump all goroutine stack traces
func main() {
	for _, profile := range pprof.Profiles() {
		fmt.Println(profile.Name())
	}
	//
	setupSigusr1Trap()
	f1()
}

func setupSigusr1Trap()  {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
			fmt.Println("heap")
			pprof.Lookup("heap").WriteTo(os.Stdout, 1)
			fmt.Println("allocs")
			pprof.Lookup("allocs").WriteTo(os.Stdout, 1)
			fmt.Println("threadcreate	")
			pprof.Lookup("threadcreate").WriteTo(os.Stdout, 1)
			fmt.Println("block")
			pprof.Lookup("block").WriteTo(os.Stdout, 1)
			fmt.Println("mutex")
			pprof.Lookup("mutex").WriteTo(os.Stdout, 1)

			//fmt.Println("eee")
			//buf := make([]byte, 16384)
			//buf = buf[:runtime.Stack(buf, true)]
			//fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
		}
	}()
}

