package main

import (
	"fmt"
	"sync"
	"time"
)

// Mapper interface 
type Mapper interface {
	IsEven(int) bool
}

// Reducer interface
type Reducer interface {
	Reduce(int, int) int
}

/* add map functionality here */
func Map(input <-chan int, output1 chan<- int, output2 chan<- int, mapper Mapper, wg *sync.WaitGroup) {
	// Call done on the waiting group
	defer wg.Done()

	// Loop over all inputs from the channel until it is done
	for num := range input {
		isEven := mapper.IsEven(num)
		if isEven {
			output1 <- num * num
		} else {
			output2 <- num * num
		}
	}
}

func Reduce(input <-chan int, output chan<- int, reducer Reducer, init int /* more channels here, if required */) {
	sum := init
	for num := range input {
		sum = reducer.Reduce(sum, num)
	}

	// When there is no more input, send the result to the output channel
	output <- sum
}

/* add code for shutdown functionality here, if required */
type EvenOddMapper struct{}
type EvenOddReducer struct{}

/* implement your interfaces here */
func (e EvenOddMapper) IsEven(num int) bool {
	return num%2 == 0
}

func (e EvenOddReducer) Reduce(first_num int, second_num int) int {
	return first_num + second_num
}

func main() {
	// Create a variable for keeping track of how many numbers to add
	max_number_to_square := 5

	// Print the number
	fmt.Println("Max number to be squared: ", max_number_to_square)

	// Starting a timer to check the performance
	start := time.Now()

	// Input channels
	input := make(chan int)
	inter1 := make(chan int)
	inter2 := make(chan int)
	output := make(chan int)

	// Create a waiting group for the two mapper
	// Using the waiting group to close the channels to the two reducers when there is
	var wg sync.WaitGroup
	wg.Add(2)

	/* channels and processes for shutdown here */
	go Map(input, inter1, inter2, EvenOddMapper{}, &wg)
	go Map(input, inter1, inter2, EvenOddMapper{}, &wg)

	// Routine that closes the channels after all mapping channels are done
	// This is done after the input channels are done
	go func() {
		wg.Wait()
		close(inter1)
		close(inter2)
	}()

	go Reduce(inter1, output, EvenOddReducer{}, 0)
	go Reduce(inter2, output, EvenOddReducer{}, 0)

	// Send numbers in their own go routine
	total := 0
	for i := 1; i <= max_number_to_square; i++ {
		input <- i
		total += i * i
	}

	// Closing the input channel after we sent all the input numbers
	close(input)

	// Read from the output twice and sum
	// Will block the main thread until both values are read from the channel
	res := <-output + <-output

	// Closing the output channel
	close(output)

	// Print the result of both reducers summed
	fmt.Println("Result of reducers:", res)

	// Just as Oblig 1, check if we got the correct sum
	fmt.Println("Total - Res:", total-res)

	// Print the time it took
	elapsed_time_in_seconds := time.Since(start).Seconds()
	fmt.Printf("Execution took %f seconds\n", elapsed_time_in_seconds)
}
