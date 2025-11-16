package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
 */
func main() {
	var wg sync.WaitGroup
	taskProcessors := []func(){}
	taskProcessor1 := func() {
		fmt.Println("Task-1")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	}
	taskProcessor2 := func() {
		fmt.Println("Task-2")
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
	}
	taskProcessors = append(taskProcessors, taskProcessor1, taskProcessor2)
	tasks := []*Task{}
	for i, tp := range taskProcessors {
		wg.Add(1)
		task := &Task{Name: fmt.Sprintf("Task-%d", i+1), Processor: tp}
		tasks = append(tasks, task)
		go func(task *Task) {
			defer wg.Done()
			task.Execute()
		}(task)
	}
	wg.Wait()
	fmt.Println("All tasks completed")
	for _, task := range tasks {
		fmt.Printf("Task %v executed in %v\n", task.Name, task.Elapsed)
	}
}

type Task struct {
	Name      string
	StartTime time.Time
	Elapsed   time.Duration
	Processor func()
}

func (t *Task) Execute() {
	t.StartTime = time.Now()
	t.Processor()
	t.Elapsed = time.Since(t.StartTime)
	fmt.Printf("Task %v executed in %v\n", t.Name, t.Elapsed)
}
