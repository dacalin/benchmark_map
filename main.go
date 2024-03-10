package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	size := 10000000
	printMemStats("Init")

	a := make(map[int]string)
	printMemStats("Map created")

	for i := 0; i < size; i++ {
		a[i] = ""
	}

	runtime.GC()
	printMemStats("After Map Add")

	for i := 0; i < size/5; i++ {
		delete(a, i)
	}

	runtime.GC()
	printMemStats("After Map Delete")

	a = nil

	runtime.GC()
	time.Sleep(1)

	printMemStats("After Nil Map")
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：memory = %vKB, GC Times = %v\n", mag, m.Alloc/1024, m.NumGC)
}
