package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < size; i++ {
		data[i] = r.Intn(size) + 1
	}

	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxResults := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(chunkIndex, startIdx, endIdx int) {
			defer wg.Done()

			chunkMax := data[startIdx]
			for j := startIdx + 1; j < endIdx; j++ {
				if data[j] > chunkMax {
					chunkMax = data[j]
				}
			}
			maxResults[chunkIndex] = chunkMax
		}(i, start, end)
	}

	wg.Wait()

	return maximum(maxResults)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("\nИщем максимальное значение в один поток")
	startSingle := time.Now()
	max := maximum(data)
	elapsed := time.Since(startSingle)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d микросекунд\n", max, elapsed.Microseconds())

	fmt.Printf("\nИщем максимальное значение в %d потоков", CHUNKS)
	startMulti := time.Now()
	max = maxChunks(data)
	elapsed = time.Since(startMulti)
	fmt.Printf("\nМаксимальное значение элемента: %d\nВремя поиска: %d микросекунд\n", max, elapsed.Microseconds())
}
