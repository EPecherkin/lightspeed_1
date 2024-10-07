package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

const (
	monitorFile = "perfmon.csv"
	monitorEach = 1 * time.Second
)

func trackUsage(quit chan bool) {
	file, err := os.Create(monitorFile)
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	var mem runtime.MemStats
	started := time.Now()

	for {
		select {
		case <-quit:
			return
		default:
			runtime.ReadMemStats(&mem)

			seconds := time.Now().Sub(started).Seconds()
			alloc := mem.TotalAlloc / (1024 * 1024) // convert to MB
			malloc := mem.Mallocs / (1024 * 1024)   // convert to MB
			_, err := file.WriteString(fmt.Sprintf("%f,%d,%d,%d,%d\n", seconds, alloc, malloc, runtime.NumCPU(), runtime.NumGoroutine()))
			if err != nil {
				log.Printf("Error writing performance data: %v\n", err)
				return
			}

			time.Sleep(1 * time.Second)
		}
	}
}
