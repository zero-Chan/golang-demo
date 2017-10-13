package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var (
	N     *int
	Iouse *int
)

func init() {
	N = flag.Int("N", 40, "Fibonacci Count")
	Iouse = flag.Int("Iouse", 1, "go a io handler func.")
	flag.Parse()

}

func main() {

	runtime.GOMAXPROCS(1)
	stop := make(chan bool)

	if *Iouse == 1 {

		go func() {
			f, _ := os.Open("./hello.txt")
			buf := make([]byte, 10)

			for {
				time.Sleep(50 * time.Second)
				f.Read(buf)
			}
		}()
	}

	go func() {
		start := time.Now()

		Fibonacci(*N)

		end := time.Now()
		fmt.Printf("CPU Handler usage [%s]\n", end.Sub(start).String())
		stop <- true
	}()

	<-stop
}

func Fibonacci(n int) int64 {
	if n <= 0 {
		return 1
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
