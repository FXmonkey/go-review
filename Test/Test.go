package Test

import (
	"fmt"
	"strings"
	"time"
)

func Test_chan() {

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	done := make(chan bool)

	go func() {
		for {
			select {
			case ch1 <- "msg1 . . . . ":
				time.Sleep(800 * time.Millisecond)
			case <-done:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case ch2 <- "msg2 . . . . ":
				time.Sleep(1000 * time.Millisecond)
			case <-done:
				return
			}
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-done:
			return
		}
	}

}

func Test_string() {
	s := "hello"
	s = s + "world"
	build := strings.Builder{}
	build.WriteString(s)
	build.WriteString("!")
	build_conc := build.String()

	fmt.Println(build_conc)

	fmt.Println(s[0]) // 字节，非字符
	// error s[0] = 'h'
	byte_s := make([]byte, len(s))
	copy(byte_s, s)
	fmt.Println(string(byte_s))
	dst := strings.Clone(s)
	fmt.Println("dst:", dst)
}
