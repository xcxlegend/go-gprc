package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Vec struct {
	X int32
	Y int32
}

func (v *Vec) Distance(v2 *Vec) float64 {
	return math.Sqrt(math.Pow(float64(v2.X-v.X), 2) + math.Pow(float64(v2.Y-v.Y), 2))
}

func main() {
	var fireC = make(chan interface{})
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var timeout = time.After(30 * time.Second)
		for {
			select {
			case f := <-fireC:
				fmt.Println(f)
			case <-timeout:
				fmt.Println("timeout")
				return
			}
		}
	}()
	var speed = 20
	var vec1 = &Vec{10, 10}
	var vec2 = &Vec{40, 50}
	var dis = vec1.Distance(vec2)
	var t = dis / float64(speed) // t->second
	fmt.Printf("walk to distance %v use time: %v\n", dis, t)
	var resetTime = make(chan float64)
	var startTime = time.Now()
	go func() {
		defer func() {
			fmt.Println("walk done")
			t := time.Now().Sub(startTime)
			fmt.Printf("total use: %v\n", t)
		}()
		var timer = time.NewTimer(time.Duration(t*1000) * time.Millisecond)
		//var after = time.After(time.Duration(t*1000) * time.Millisecond)
		for {
			select {
			case <-timer.C:
				go func() {
					fireC <- "fire"
				}()
				return
			case rt := <-resetTime:
				if rt > 0 {
					//after = time.After(time.Duration(rt*1000) * time.Millisecond)
					ok := timer.Reset(time.Duration(rt*1000) * time.Millisecond)
					fmt.Println("reset", ok)
				}
			}
		}
	}()
	// 返回一个剩余时间
	var jiasu = func() float64 {
		useTime := time.Now().Sub(startTime).Milliseconds()
		useDis := float64(int64(speed) * useTime / 1000)
		if dis > useDis {
			dis -= useDis
			newSpeed := 25
			var t = dis / float64(newSpeed) // t->second
			fmt.Printf("new walk to distance %v use time: %v\n", dis, t)
			return t
		}
		return 0
	}
	_ = jiasu
	time.Sleep(1000 * time.Millisecond)
	resetTime <- jiasu()
	wg.Wait()
}
