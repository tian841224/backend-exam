package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Employee struct {
	ID int
}

type Item1 struct {
}

func (i *Item1) Process() {
	fmt.Println("開始處理 Item1")
	time.Sleep(1 * time.Second)
	fmt.Println("結束處理 Item1")
}

type Item2 struct {
}

func (i *Item2) Process() {
	fmt.Println("開始處理 Item2")
	time.Sleep(2 * time.Second)
	fmt.Println("結束處理 Item2")
}

type Item3 struct {
}

func (i *Item3) Process() {
	fmt.Println("開始處理 Item3")
	time.Sleep(3 * time.Second)
	fmt.Println("結束處理 Item3")
}

type Item interface {
	// Process 這是一個耗時操作
	Process()
}

func main() {

	const (
		workers    = 5
		perTypeNum = 10
		totalTasks = perTypeNum * 3
	)

	items := make([]Item, 0, totalTasks)

	// 準備物品：各 10 件
	for i := 0; i < perTypeNum; i++ {
		items = append(items, &Item1{})
		items = append(items, &Item2{})
		items = append(items, &Item3{})
	}

	// 隨機打亂物品順序
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(items), func(i, j int) { items[i], items[j] = items[j], items[i] })

	// 記錄開始時間
	startTime := time.Now()
	fmt.Printf("開始處理，共 %d 件物品\n", totalTasks)

	// 分配任務給員工，每個員工一次只能處理一種物品
	ch := make(chan Item, workers)
	done := make(chan bool, workers)
	workerCounts := make([]int, workers)

	for i := 0; i < workers; i++ {
		go func(workerID int) {
			count := 0
			for item := range ch {
				item.Process()
				count++
			}
			workerCounts[workerID] = count
			done <- true
		}(i)
	}

	// 分配任務給員工
	for _, item := range items {
		ch <- item
	}
	close(ch)

	// 等待所有員工完成任務
	for i := 0; i < workers; i++ {
		<-done
	}

	// 記錄結束時間並統計
	endTime := time.Now()
	totalDuration := endTime.Sub(startTime)

	fmt.Printf("結束處理，總處理時間: %v\n", totalDuration)
	fmt.Println("每個員工處理的物品數量:")
	for i, count := range workerCounts {
		fmt.Printf("  員工 %d: %d 件\n", i+1, count)
	}
}
