package parallel

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestSimple(t *testing.T) {
	taskSlice := []func() error{
		func() error { fmt.Println("print a"); return nil },
		func() error { fmt.Println("print b"); return nil },
		func() error { fmt.Println("print c"); return nil },
	}

	assert.NoError(t, Run(taskSlice, 3, 3))
}

func TestSequential(t *testing.T) {
	var arr []int
	taskSlice := []func() error{
		func() error { arr = append(arr, 1); return nil },
		func() error { arr = append(arr, 2); return nil },
		func() error { arr = append(arr, 3); return nil },
	}

	Run(taskSlice, 1, 3)
	assert.Equal(t, []int{1, 2, 3}, arr)
}

func TestStopOnError(t *testing.T) {
	N := 3
	M := 1
	var completedTasks uint64
	taskSlice := []func() error{
		func() error {
			fmt.Println("error")
			atomic.AddUint64(&completedTasks, 1)
			return errors.New("Some error")
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
	}

	assert.Error(t, Run(taskSlice, N, M), "Limit of errors was exceeded")
	assert.LessOrEqual(t, completedTasks, uint64(N+M))
}

func TestStopOnErrorComplicated(t *testing.T) {
	N := 3
	M := 2
	var completedTasks uint64
	taskSlice := []func() error{
		func() error {
			fmt.Println("error")
			atomic.AddUint64(&completedTasks, 1)
			return errors.New("Some error")
		},
		func() error {
			fmt.Println("error")
			atomic.AddUint64(&completedTasks, 1)
			return errors.New("Some error")
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
		func() error {
			fmt.Println("good")
			time.Sleep(10 * time.Millisecond)
			atomic.AddUint64(&completedTasks, 1)
			return nil
		},
	}

	assert.Error(t, Run(taskSlice, N, M), "Limit of errors was exceeded")
	assert.LessOrEqual(t, completedTasks, uint64(N+M))
}

func TestEmpty(t *testing.T) {
	assert.NoError(t, Run([]func() error{}, 100, 100))
}

func TestRandomLoad(t *testing.T) {
	count := 1000
	rand.Seed(time.Now().UnixNano())
	maxErrors := rand.Intn(count)
	errorFunsN := rand.Intn(count)
	fmt.Println("Max errors: ", maxErrors, "Errors count: ", errorFunsN)

	goodFun := func() error {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return nil
	}

	errorFun := func() error {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return errors.New("Some error")
	}

	taskSlice := make([]func() error, count)

	for i := 0; i < errorFunsN; i++ {
		taskSlice[i] = errorFun
	}
	for i := errorFunsN; i < 1000; i++ {
		taskSlice[i] = goodFun
	}

	err := Run(taskSlice, 10, maxErrors)
	if errorFunsN < maxErrors {
		fmt.Println("success case")
		assert.Equal(t, nil, err)
	} else {
		fmt.Println("error case")
		assert.Equal(t, "Limit of errors was exceeded", err.Error())
	}
}
