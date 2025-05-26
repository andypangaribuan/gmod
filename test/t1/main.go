/*
 * Copyright (c) 2025.
 * Created by Andy Pangaribuan (iam.pangaribuan@gmail.com)
 * https://github.com/apangaribuan
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package main

import (
	"fmt"
	// "math"
	// "math/rand"
	"sync"
	"time"
)

var (
	mx       sync.Mutex
	syncMap  sync.Map
	latestId = 0
)

func main() {
	start := time.Now()
	exec()
	// exec()
	// exec()
	// exec()
	fmt.Println("-- FINAL DONE", time.Since(start))
}

func exec() {
	loop := 10
	totalJob := 200

	s := &stu{
		maxParallel: 10,
		needShow:    false,
		totalJob:    totalJob,
		job:         make(chan int, totalJob*10),
		waiter:      make(chan int, totalJob*10),
		// job:         make(chan int, math.MaxInt32),
		// waiter:      make(chan int, math.MaxInt32),
	}

	for range loop {
		start := time.Now()
		s.run()
		s.show()
		fmt.Println("--       DONE", time.Since(start))
	}

	s.prune()
}

func (slf *stu) run() {
	syncMap = sync.Map{}

	if !slf.hasInit {
		slf.hasInit = true
		for range slf.maxParallel {
			// go slf.worker(slf.job, slf.waiter)
			go slf.worker()
		}
	}

	// for range slf.maxParallel {
	// 	// go slf.worker(slf.job, slf.waiter)
	// 	go slf.worker()
	// }

	for i := range slf.totalJob {
		slf.job <- i
	}
	// close(slf.job)

	for range slf.totalJob {
		<-slf.waiter
	}
	// close(slf.waiter)

	// var ms = make(map[int]int, 0)

	// syncMap.Range(func(k, v any) bool {
	// 	key := v.(int)

	// 	value, ok := ms[key]
	// 	if ok {
	// 		ms[key] = value + 1
	// 	} else {
	// 		ms[key] = 1
	// 	}

	// 	return true
	// })

	// time.Sleep(time.Second * 3)
	// for i := range slf.maxParallel {
	// 	key := i + 1
	// 	value, ok := ms[key]
	// 	if !ok {
	// 		continue
	// 	}

	// 	fmt.Printf("%v: %v\n", key, value)
	// }
}

// func (slf *stu) worker(job <-chan int, waiter chan<- int) {
func (slf *stu) worker() {
	id := getId()

	// min := 9
	// max := 11

	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)

	for j := range slf.job {
		// randomNumber := r.Intn(max-min) + min
		// time.Sleep(time.Millisecond * time.Duration(randomNumber))
		time.Sleep(time.Millisecond * 10)
		if slf.needShow {
			syncMap.Store(j, id)
		}
		slf.waiter <- 0
	}
}

func (slf *stu) prune() {
	close(slf.job)
	close(slf.waiter)
}

func (slf *stu) show() {
	if !slf.needShow {
		return
	}

	var ms = make(map[int]int, 0)

	syncMap.Range(func(k, v any) bool {
		key := v.(int)

		value, ok := ms[key]
		if ok {
			ms[key] = value + 1
		} else {
			ms[key] = 1
		}

		return true
	})

	time.Sleep(time.Second * 3)
	for i := range slf.maxParallel {
		key := i + 1
		value, ok := ms[key]
		if !ok {
			continue
		}

		fmt.Printf("%v: %v\n", key, value)
	}
}

func getId() int {
	mx.Lock()
	defer mx.Unlock()
	latestId++
	return latestId
}

type stu struct {
	hasInit     bool
	needShow    bool
	maxParallel int
	// jobs        []int
	totalJob int
	job      chan int
	waiter   chan int
}
