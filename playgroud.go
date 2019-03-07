package main

import (
	"fmt"
)

func main() {
	//timer := time.NewTimer(time.Duration(3) * time.Second)
	//
	//ch := make(chan int)
	//go func() {
	//	time.Sleep(4 * time.Second)
	//	ch <- 1
	//}()
	//
	//select {
	//case <-timer.C:
	//	fmt.Println("timeout")
	//
	//case <-ch:
	//	fmt.Println("got an integer")
	//}
	//
	//fmt.Println("finish")
	//
	//a := [3]string{"12", "#54", "#12"}
	//
	//t := reflect.TypeOf(a)
	//
	//fmt.Println(a, t.Kind())
	
	addOne := closures(1)
	fmt.Println(addOne())
	fmt.Println(addOne())
	fmt.Println(addOne())

	addTwo := closures(2)
	fmt.Println(addTwo())
	fmt.Println(addTwo())
	fmt.Println(addTwo())

	fmt.Println(fact(7))

}

func closures(step int) func() int {
	var i int

	return func() int {
		i += step;
		return i;
	}
}

func fact(n int) int {
	return _fact(n, 1)
}

func _fact(n, result int) int {
	if 0 == n {
		return result
	}

	return _fact(n-1, result*n)
}