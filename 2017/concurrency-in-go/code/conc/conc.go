package main

func main() {
    // (1) create a channel
    c := make(chan int32)

    // (2) anonymous function is invoked in the goroutine
    go func(ch chan int32) {

        // (3) wait for the data from the channel
        value := <-ch

        if value == 42 {
            println("\\ʕ◔ϖ◔ʔ/")
        }
    }(c)

    // (4) send value to the channel
    c <- 42
}
