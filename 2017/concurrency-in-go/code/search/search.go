package main

import (
	"math/rand"
	"time"
)

// START_DEFS OMIT
type Result []string

type Search interface {
	Search(string) Result
}

// END_DEFS OMIT

type (
	Web   struct{}
	Maps  struct{}
	Image struct{}
	Video struct{}
)

var web, ag, grep, onion Web
var maps, gps, glonass Maps
var image, memes, album2k16 Image
var video, cctv, yourtube Video

// START_IMPL OMIT
func (search Web) Search(query string) Result {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return Result{"..."}
}

// END_IMPL OMIT

func (search Maps) Search(query string) Result {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return nil
}

func (search Image) Search(query string) Result {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return nil
}

func (search Video) Search(query string) Result {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	return nil
}

// START_LINEAR OMIT
func LinearSearch(query string) []Result {
	var res []Result
	res = append(res, web.Search(query))
	res = append(res, maps.Search(query))
	res = append(res, image.Search(query), video.Search(query))
	return res
}

// END_LINEAR OMIT

// START_CONC OMIT
func ConcurrentSearch(query string) (res []Result) {
	ch := make(chan Result, 4)
	go func() { ch <- web.Search(query) }()
	go func() { ch <- maps.Search(query) }()
	go func() { ch <- image.Search(query) }()
	go func() { ch <- video.Search(query) }()

	for {
		value, ok := <-ch
		if ok {
			res = append(res, value)
		} else {
			return res
		}
	}
}

// END_CONC OMIT

// START_FAST OMIT
func FastSearch(query string) (res []Result) {
	ch := make(chan Result, 4)
	go func() { ch <- web.Search(query) }()
	go func() { ch <- maps.Search(query) }()
	go func() { ch <- image.Search(query) }()
	go func() { ch <- video.Search(query) }()

	timeout := time.After(78 * time.Millisecond)
	for i := 0; i < 4; i++ {
		select {
		case value := <-ch:
			res = append(res, value)
		case <-timeout:
			return res
		}
	}
	return
}

// END_FAST OMIT

// START_FIRST OMIT
func FirstOf(query string, servers ...Search) Result {
	ch := make(chan Result)
	for i := range servers {
		go func(i int) {
			ch <- servers[i].Search(query)
		}(i)
	}
	return <-ch
}

// END_FIRST OMIT

// START_SMART OMIT
func SmartSearch(query string) (res []Result) {
	ch := make(chan Result, 4)
	go func() { ch <- FirstOf(query, web, ag, grep, onion) }()
	go func() { ch <- FirstOf(query, gps, glonass) }()
	go func() { ch <- FirstOf(query, memes, album2k16) }()
	go func() { ch <- FirstOf(query, cctv, yourtube) }()

	timeout := time.After(78 * time.Millisecond)
	for i := 0; i < 4; i++ {
		select {
		case value := <-ch:
			res = append(res, value)
		case <-timeout:
			return res
		}
	}
	return
}

// END_SMART OMIT

func main() {
	LinearSearch("..")
}
