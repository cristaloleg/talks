package main

import "fmt"
import "time"
import "sync"
import "sync/atomic"

// START_POOL OMIT
type WorkerPool struct {
    wg sync.WaitGroup
    sz int32
    ch <-chan int
    fn func(int)
}

// END_POOL OMIT

// START_NEW OMIT
func New(amount int32, jobs <-chan int, f func(int)) *WorkerPool {
    return &WorkerPool{
        wg: sync.WaitGroup{},
        sz: amount,
        ch: jobs,
        fn: f,
    }
}

// END_NEW OMIT

// START_START OMIT
func (w *WorkerPool) Start() {
    sz := w.sz
    w.sz = 0
    for i := int32(1); i <= sz; i++ {
        w.invoke(i)
    }
}

// END_START OMIT

// START_INVOKE OMIT
func (w *WorkerPool) invoke(id int32) {
    w.wg.Add(1)
    atomic.AddInt32(&w.sz, 1)
    go func() {
        for j := range w.ch {
            fmt.Println("worker", id, "started  job", j)
            w.fn(500)
            fmt.Println("worker", id, "finished job", j)

            if atomic.LoadInt32(&w.sz) < id {
                w.wg.Done()
                return
            }
        }
        w.wg.Done()
    }()
}

// END_INVOKE OMIT

// START_INC OMIT
func (w *WorkerPool) Inc(delta int32) {
    sz := atomic.LoadInt32(&w.sz)
    for i := int32(1); i < delta; i++ {
        w.invoke(i + sz)
    }
}

// END_INC OMIT

// START_DEC OMIT
func (w *WorkerPool) Dec(delta int32) {
    atomic.AddInt32(&w.sz, -delta)
}

// END_DEC OMIT

// START_WAIT OMIT
func (w *WorkerPool) Wait() {
    w.wg.Wait()
}

// END_WAIT OMIT

// START_MAIN OMIT
func main() {
    jobs := make(chan int)
    wait := func(i int) { time.Sleep(time.Duration(i) * time.Millisecond) }
    pool := New(3, jobs, wait)

    pool.Start()

    go func() {
        wait(1700)
        println("oh no...")
        pool.Dec(2)
    }()
    go func() {
        wait(10700)
        println("OH MYYYYY!!! \\ʕ◔ϖ◔ʔ/")
        pool.Inc(8)
    }()
    for i := 1; i <= 100; i++ {
        jobs <- i
    }
    close(jobs)

    pool.Wait()
}

// END_MAIN OMIT
