package main

import (
	"fmt"
	"time"
)
//channel赋值
func Generate(ch chan<- int) {
	//赋值大于2的数字
	for i := 2; ; i++ {
		ch <- i
	}
}
//channel取值
//in 输入channel ，out 是输出channel》
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		//取值
		i := <-in
		//除以2没有余数才能手粗
		if i%prime != 0 {
			out <- i
		}
	}
}
func main() {
	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine. for i := 0; i < 15; i++ {
	for i := 0; i < 15; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

}
