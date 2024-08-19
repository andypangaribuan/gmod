/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package server

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/andypangaribuan/gmod/gm"
	"github.com/go-co-op/gocron"
)

func (slf *stuServer) Cron(routes func(router RouterC)) {
	timezone, err := mrf2[string, error]("mrf-conf-val", "timezone")
	if err != nil {
		log.Fatalf("timezone not found when start the cron engine\n")
	}

	var (
		loc, _    = time.LoadLocation(timezone)
		scheduler = gocron.NewScheduler(loc)
		router    = &stuCronRouter{
			everyItems:  make([]*stuCronEveryItem, 0),
			everyNItems: make([]*stuCronNEveryItem, 0),
		}
	)

	routes(router)

	for _, item := range router.everyItems {
		slf.add(item.uid)
		slf.scheduleEveryItem(scheduler, item.uid, item.duration, item.startUpDelayed, item.allowParallel, item.fn)
	}

	for _, item := range router.everyNItems {
		slf.add(item.uid)
		slf.scheduleEveryItem(scheduler, item.uid, item.duration, item.startUpDelayed, item.allowParallel, item.fns...)
	}

	for _, item := range router.everyDayItems {
		slf.add(item.uid)
		slf.scheduleEveryDayItem(scheduler, item.uid, item.at, item.allowParallel, item.fn)
	}

	for _, item := range router.everyDayNItems {
		slf.add(item.uid)
		slf.scheduleEveryDayItem(scheduler, item.uid, item.at, item.allowParallel, item.fns...)
	}

	scheduler.StartAsync()
}

func (slf *stuServer) add(uid string) {
	cronMX[uid] = &sync.Mutex{}
	cronIsStartUp[uid] = false
}

func (slf *stuServer) scheduleEveryItem(scheduler *gocron.Scheduler, uid string, duration string, startUpDelayed *time.Duration, allowParallel bool, fns ...func()) {
	totalFns := len(fns)
	if totalFns == 0 {
		return
	}

	if startUpDelayed != nil {
		cronIsStartUp[uid] = true
	}

	_, err := scheduler.Every(duration).Do(func() {
		if cronIsStartUp[uid] {
			time.Sleep(*startUpDelayed)
			cronIsStartUp[uid] = false
		}

		if allowParallel {
			for _, fn := range fns {
				slf.executeFn(fn, nil)
			}
			return
		}

		cronMX[uid].Lock()
		defer cronMX[uid].Unlock()

		if totalFns == 1 {
			slf.executeFn(fns[0], nil)
		} else {
			wg := new(sync.WaitGroup)
			for i := 0; i < totalFns; i++ {
				wg.Add(1)
				go slf.executeFn(fns[i], wg)
			}

			wg.Wait()
		}
	})

	if err != nil {
		log.Fatalf("cron [every] error: %+v\n", err)
	}
}

func (slf *stuServer) scheduleEveryDayItem(scheduler *gocron.Scheduler, uid string, at string, allowParallel bool, fns ...func()) {
	totalFns := len(fns)
	if totalFns == 0 {
		return
	}

	_, err := scheduler.Every(1).Day().At(at).Do(func() {
		if allowParallel {
			for _, fn := range fns {
				slf.executeFn(fn, nil)
			}
			return
		}

		cronMX[uid].Lock()
		defer cronMX[uid].Unlock()

		if totalFns == 1 {
			slf.executeFn(fns[0], nil)
		} else {
			wg := new(sync.WaitGroup)
			for i := 0; i < totalFns; i++ {
				wg.Add(1)
				go slf.executeFn(fns[i], wg)
			}

			wg.Wait()
		}
	})

	if err != nil {
		log.Fatalf("cron [every day] error: %+v\n", err)
	}
}

func (slf *stuServer) executeFn(fn func(), wg *sync.WaitGroup) {
	err := gm.Util.PanicCatcher(fn)
	if err != nil {
		fmt.Printf("cron error: %+v\n", err)
	}

	if wg != nil {
		wg.Done()
	}
}
