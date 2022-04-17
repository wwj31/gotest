package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	t1 := time.Now().Unix()
	defer func() {
		fmt.Println(time.Now().Unix() - t1)
	}()

	ti := time.NewTimer(1 * time.Second)
	time.Sleep(2000 * time.Millisecond)
	if !ti.Stop() {
		println("stop false")
		_, ok := <-ti.C
		println(ok)
	}
	ti.Reset(10 * time.Second)

	<-ti.C
}

// 永久等待timer.C例子1
func TestErrTimer1(t *testing.T) {
	ti := time.NewTimer(1000 * time.Millisecond)
	ti.Stop()
	<-ti.C
	// stop会把timer从堆里移除，但是不会close time.C 结果导致永远不会触发计时器
}

// 永久等待timer.C例子2
func TestErrTimer2(t *testing.T) {
	println(time.Now().Unix())
	ti := time.NewTimer(time.Second)
	time.Sleep(2 * time.Second)

	ti.Reset(1 * time.Second)
	time.Sleep(2 * time.Second)

	tt := <-ti.C
	println("c1", tt.Unix())
	tt = <-ti.C
	println("c2", tt.Unix())
	// 直接调用reset导致里面的chan没有消费干净，所以第一次能直接取出c1
	// 而c2需要等到reset的10秒才会触发,从而导致非预期结果
}

// 永久等待timer.C例子2
func TestErrTimer3(t *testing.T) {
	ti := time.NewTimer(time.Second)

	go func() {
		ti.Stop()
	}()
	time.Sleep(1 * time.Second)

	if !ti.Stop() {
		println("<-C")
		<-ti.C
		println("bbb")
	}
	ti.Reset(1 * time.Second)
	println("ok")
}
