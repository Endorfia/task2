package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := make(chan int)

	var wg sync.WaitGroup

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			sum <- n
		}(n)
	}

	go func() {
		wg.Wait()
		close(sum)
	}()

	total := 0
	for n := range sum {
		total += n
	}

	fmt.Println(total)
}
