package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 修改数组元素
func TestReplaceArray(t *testing.T) {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	for i := 0; i < len(arr); i++ {
		if arr[i] == "stupid" {
			arr[i] = "smart"
		} else if arr[i] == "weak" {
			arr[i] = "strong"
		}
	}

	fmt.Println(arr)
}

// 单生产者 单消费者
func TestProducerConsumer(t *testing.T) {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(1)

	producerTicker := time.NewTicker(1 * time.Second)
	defer producerTicker.Stop()
	go func(t *time.Ticker) {
		for {
			<- t.C
			c <- 1
		}
	}(producerTicker)

	consumerTicker := time.NewTicker(1 * time.Second)
	defer consumerTicker.Stop()
	go func(t *time.Ticker) {
		for {
			<- t.C
			i := <- c
			fmt.Printf("receive %d \n", i)
		}
	}(consumerTicker)

	wg.Wait()
}



















