package main

import (
    "fmt"
    "time"
    "sync"
    "math/rand"
)

// Go provides a simple syntax over two powerful mechanisms:
// goroutines and channels.
func goroutineDemo() {
    fmt.Println("start")
    go process()
    fmt.Println("ready to sleep for 10 milliseconds.")
    time.Sleep(time.Millisecond * 10)

    fmt.Println("complete the goroutine demo.")
}

// A goroutine is similar to a thread,
// but it is scheduled by Go, not the OS.
//
// Multiple goroutines will end up running
// on the same underlying OS thread. 
// This is often called an M:N threading model 
// because we have M application threads (goroutines)
// running on N OS threads. The result is that
// a goroutine has a fraction of overhead (a few KB)
// than OS threads. On modern hardware, 
// it’s possible to have millions of goroutines.
func process() {
    fmt.Println("processing in a goroutine..")
}

var (
    counter = 0
    lock sync.Mutex
)

func synchronizationDemo() {
    for i := 0; i < 20; i++ {
        go incr()
    }
    time.Sleep(time.Millisecond * 10)
}

func incr() {
    lock.Lock()
    defer lock.Unlock()
    counter++
    fmt.Println(counter)
}


func channelDemo() {
    c := make(chan int)
    for i := 0; i < 5; i++ {
        worker := &Worker{id: i}
        go worker.process(c)
    }

    for {
        c <- rand.Int()
        time.Sleep(time.Millisecond * 50)
    }
}

type Worker struct {
    id int
}

func (w *Worker) process(c chan int) {
    for {
        data := <- c
        fmt.Printf("worker %d got %d.\n", w.id, data)
    }
}

func main() {
    goroutineDemo()
    synchronizationDemo()
    channelDemo()

    fmt.Println("done.")
}
