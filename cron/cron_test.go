package cron

import (
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	now := time.Now().Add(time.Second * 2)
	cronExpr := fmt.Sprintf("%d %d %d %d %d * %d", now.Second(), now.Minute(), now.Hour(), now.Day(), now.Month(), now.Year())
	c, err := New(cronExpr, func() {})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	go c.Run()
	<-c.Signals
	assert.Equal(t, c.Status, Finished)
	assert.Equal(t, c.RunCount, int64(1))
}

func TestRunWithInterruptSignal(t *testing.T) {
	cronExpression := "*/2 * * * * * *" // Runs at every 2s
	c, err := New(cronExpression, func() {})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		c.Run()
		wg.Done()
	}()

	// Simulate sending an interrupt signal after a short delay
	go func() {
		time.Sleep(5 * time.Second)
		c.Signals <- os.Interrupt
	}()

	// Wait for the cron job to finish or receive an interrupt
	wg.Wait()

	assert.Equal(t, c.Status, Finished)
	assert.Equal(t, c.RunCount, int64(2))
}

func TestGetNextExecutionTime(t *testing.T) {
	cronExpression := "*/2 * * * * * *" // Runs at every 2s
	c, err := New(cronExpression, func() {})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	now := time.Now()
	nextExecutionTime := c.GetNextExecutionTime(now)

	assert.True(t, nextExecutionTime.After(now))
}

func TestIsNextExecutionTime(t *testing.T) {
	cronExpression := "*/3 * * * * * *" // Runs at every 3s
	c, err := New(cronExpression, func() {})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	now := time.Now()
	nextExecutionTime := c.GetNextExecutionTime(now)
	isNextExecutionTime := c.isNextExecutionTime(nextExecutionTime)
	assert.True(t, isNextExecutionTime)
}

func TestHasNextExecutionTime(t *testing.T) {
	cronExpr := "*/4 * * * * * *" // Runs at every 4s
	c, err := New(cronExpr, func() {})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	hasNextExecutionTime := c.hasNextExecutionTime(time.Now())
	assert.True(t, hasNextExecutionTime)
}

func TestNew(t *testing.T) {
	cronExpr := "*/5 * * * * * *" // Runs at every 5s
	c, err := New(cronExpr, func() {})
	assert.NotNil(t, c)
	assert.Nil(t, err)

	assert.Equal(t, c.Status, Idle)
	assert.Equal(t, c.RunCount, int64(0))
}

func TestNewWithInvalidCronExpression(t *testing.T) {
	cronExpr := "* * * * * *"
	c, err := New(cronExpr, func() {})
	assert.Nil(t, c)
	assert.NotNil(t, err)
}

func BenchmarkRun(b *testing.B) {
	now := time.Now().Add(time.Second * 2)
	cronExpr := fmt.Sprintf("%d %d %d %d %d * %d", now.Second(), now.Minute(), now.Hour(), now.Day(), now.Month(), now.Year())
	c, _ := New(cronExpr, func() {})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go c.Run()
		<-c.Signals
	}
}

func BenchmarkGetNextExecutionTime(b *testing.B) {
	cronExpr := "*/3 * * * * * *" // Runs at every 3s
	c, _ := New(cronExpr, func() {})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.GetNextExecutionTime(time.Now())
	}
}

func BenchmarkIsNextExecutionTime(b *testing.B) {
	cronExpr := "*/2 * * * * * *" // Runs at every 2s
	c, _ := New(cronExpr, func() {})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.isNextExecutionTime(time.Now())
	}
}

func BenchmarkHasNextExecutionTime(b *testing.B) {
	cronExpr := "*/1 * * * * * *" // Runs at every 1s
	c, _ := New(cronExpr, func() {})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.hasNextExecutionTime(time.Now())
	}
}

func BenchmarkNew(b *testing.B) {
	cronExpr := "* * * * * * *" // Runs at every minute
	for i := 0; i < b.N; i++ {
		New(cronExpr, func() {})
	}
}
