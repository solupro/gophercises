package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ex1
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

	// ex2
	//a := [3]string{"12", "#54", "#12"}
	//
	//t := reflect.TypeOf(a)
	//
	//fmt.Println(a, t.Kind())

	// ex3
	//addOne := closures(1)
	//fmt.Println(addOne())
	//fmt.Println(addOne())
	//fmt.Println(addOne())
	//
	//addTwo := closures(2)
	//fmt.Println(addTwo())
	//fmt.Println(addTwo())
	//fmt.Println(addTwo())
	//
	//fmt.Println(fact(7))

	// ex4
	//http.ListenAndServe(":8081", &Me{"/hello"})

	// ex5
	//piCh := make(chan string)
	//poCh := make(chan string)
	//
	//go ping(piCh, "hello world")
	//go pong(poCh, piCh)
	//fmt.Println(<-poCh)

	// ex6

}

func ping(piCh chan<- string, msg string) {
	piCh <- msg
}

func pong(poCh chan<- string, piCh <-chan string) {
	poCh <- <-piCh
}

type Me struct {
	path string
}

func (m *Me) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == m.path {
		fmt.Fprint(w, "hello 123")
	}
}

func closures(step int) func() int {
	var i int

	return func() int {
		i += step
		return i
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
