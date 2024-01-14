package main

import (
	"log/slog"
	"time"
)

func main() {
	var count int
	for {
		time.Sleep(5 * time.Second)
		count++
		slog.Info("hello world", "iteration", count)
	}
}
