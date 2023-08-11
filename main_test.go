package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func BenchmarkMain(b *testing.B) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()
	for i := 0; i < b.N; i++ {
		now := time.Now().Add(time.Second * 2)
		cronExpression := fmt.Sprintf("%d %d %d %d %d * %d", now.Second(), now.Minute(), now.Hour(), now.Day(), now.Month(), now.Year())
		os.Args = []string{"cronTool", "-s", cronExpression}
		main()
	}
}
