package main

import (
	"fmt"
	"sync"
)

type Mapper interface {
	IsEven(int) bool
}

type Reducer interface {
	Reduce(int, int) int
}

/* add map functionality here */
func Map(input <-chan int, output1 chan<- int, output2 chan<- int, mapper Mapper, wg *sync.WaitGroup) {
  // Call done on the mapper waiting group
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
	input := make(chan int)
	inter1 := make(chan int)
	inter2 := make(chan int)
	output := make(chan int)

  // Create a waiting group for the two reducers
  var wg sync.WaitGroup
  wg.Add(2)

	/* channels and processes for shutdown here */
	go Map(input, inter1, inter2, EvenOddMapper{}, &wg)
	go Map(input, inter1, inter2, EvenOddMapper{}, &wg)

  // Routine that closes the channels after all mapping channels are done
  go func(){
    wg.Wait()
    close(inter1)
    close(inter2)
  }()

	go Reduce(inter1, output, EvenOddReducer{}, 0)
	go Reduce(inter2, output, EvenOddReducer{}, 0)
	input <- 1
	input <- 2
	input <- 3
	input <- 4
	input <- 5

  // Close the input channel
  close(input)

	res := <-output + <-output
	//should be 55
	fmt.Println(res)
	close(output)

}
