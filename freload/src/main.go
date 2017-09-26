package main

import (
	"fmt"
	"path/filepath"
	"time"

	"sunteng/commons/freload"
)

type TestWheel struct {
	callTime int
}

func (w *TestWheel) Name() string {
	return "Test"
}

func (w *TestWheel) FileSize(file string) int64 {
	return 0
}

func (w *TestWheel) Version(s string) string {
	return ""
}

func (w *TestWheel) Files() []string {
	file := "./1tmp_xxoo"
	file, _ = filepath.Abs(file)
	return []string{file}
}
func (w *TestWheel) Handler(file string) error {
	w.callTime += 1
	return nil
}

func main() {
	w := &TestWheel{}
	err := freload.StartWheel(w)
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(w.callTime)

	time.Sleep(25000 * time.Millisecond)
	fmt.Println(w.callTime)
}
