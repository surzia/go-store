package main

import (
	"fmt"
	"sync"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := MapReduce(data, Mapper, Reducer)

	fmt.Println("Result:", result)
}

func Mapper(item int) int {
	return item * 2
}

func Reducer(result []int) int {
	sum := 0
	for _, item := range result {
		sum += item
	}
	return sum
}

func MapReduce(data []int, mapper func(int) int, reducer func([]int) int) int {
	numWorkers := 4
	var wg sync.WaitGroup
	dataChannel := make(chan int)
	resultChannel := make(chan int)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for item := range dataChannel {
				mapped := mapper(item)
				resultChannel <- mapped
			}
		}()
	}

	go func() {
		defer close(resultChannel)
		results := []int{}
		for mapped := range resultChannel {
			results = append(results, mapped)
		}
		result := reducer(results)
		resultChannel <- result
	}()

	go func() {
		for _, item := range data {
			dataChannel <- item
		}
		close(dataChannel)
	}()

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	result := <-resultChannel
	return result
}
