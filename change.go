package main

import (
	"fmt"
)

var myres = make(map[int]int, 20)

func factorial(n int, ch chan int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
		ch <- res
	}
	close(ch)
}
func main() {
	var ch = make(chan int)
	for i := 1; i <= 20; i++ {
		go factorial(i, ch)

	}

	for q:=1;q<20;q++{
		myres[q]=<-ch
	}

	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}

}
