package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
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
	//timer := time.NewTicker(time.Second)
	//
	//go func() {
	//	for t := range timer.C {
	//		fmt.Println("Tick at", t)
	//	}
	//}()
	//
	//time.Sleep(10 * time.Second)
	//timer.Stop()
	//fmt.Println("done")

	// ex7
	//jobs := make(chan int, 10)
	//result := make(chan *Result, 10)
	//
	//for i := 0; i < 3; i++ {
	//	go worker(i, jobs, result)
	//}
	//
	//for i := 0; i < 5; i++ {
	//	jobs <- i
	//}
	//close(jobs)
	//
	//for i := 0; i < 5; i++ {
	//	r := <- result
	//	fmt.Printf("worker-%d result=%d\n", r.Id, r.Num)
	//}
	//
	//fmt.Println("done")

	// ex8
	//state := make(map[int]int)
	//mutex := &sync.Mutex{}
	//
	//var r uint64
	//var w uint64
	//
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		for {
	//			key := rand.Intn(5)
	//			val := rand.Intn(100)
	//			mutex.Lock()
	//			state[key] = val
	//			mutex.Unlock()
	//
	//			atomic.AddUint64(&w, 1)
	//
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//for i := 0; i < 100; i++ {
	//	go func() {
	//		total := 0
	//		for {
	//			key := rand.Intn(5)
	//			mutex.Lock()
	//			total += state[key]
	//			mutex.Unlock()
	//
	//			atomic.AddUint64(&r, 1)
	//
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//time.Sleep(time.Second)
	//
	//readFinal := atomic.LoadUint64(&r)
	//writeFinal := atomic.LoadUint64(&w)
	//
	//fmt.Println("read ops", readFinal)
	//fmt.Println("write ops", writeFinal)
	//
	//fmt.Println(state)

	// ex9
	//var readOps uint64
	//var writeOps uint64
	//
	//reads := make(chan *readOp)
	//writes := make(chan *writeOp)
	//
	//go func() {
	//	var state = make(map[int]int)
	//	for {
	//		select {
	//		case read := <- reads:
	//			read.resp <- state[read.key]
	//		case write := <- writes:
	//			state[write.key] = write.val
	//			write.resp <- true
	//		}
	//	}
	//}()
	//
	//for r := 0; r < 100; r++ {
	//	go func() {
	//		for {
	//			read := &readOp{
	//				key: rand.Intn(5),
	//				resp: make(chan int)}
	//			reads <- read
	//			<- read.resp
	//
	//			atomic.AddUint64(&readOps, 1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//for w := 0; w < 10; w++ {
	//	go func() {
	//		for {
	//			write := &writeOp{
	//				key: rand.Intn(5),
	//				val: rand.Intn(100),
	//				resp: make(chan bool)}
	//			writes <- write
	//			<- write.resp
	//
	//			atomic.AddUint64(&writeOps, 1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//time.Sleep(time.Second)
	//
	//readOpsFinal := atomic.LoadUint64(&readOps)
	//fmt.Println("readOps:", readOpsFinal)
	//writeOpsFinal := atomic.LoadUint64(&writeOps)
	//fmt.Println("writeOps:", writeOpsFinal)

	// []byte{240, 159, 152, 134})

}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func worker(id int, jobs <-chan int, result chan<- *Result) {
	for job := range jobs {
		fmt.Printf("worker-%d processing job-%d\n", id, job)
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

		fmt.Printf("worker-%d finish job-%d\n", id, job)
		result <- &Result{Num: job * 2, Id: id}
	}
}

type Result struct {
	Num int
	Id  int
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
