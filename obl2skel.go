package main

import "fmt"


type Mapper interface {
/* fill signatures here */
}

type Reducer interface {
/* fill signatures here */
}

func Map(input <-chan int, output1 chan<- int, output2 chan<- int, mapper Mapper /* more channels here, if required */) {
  /* add map functionality here */
}

func Reduce(input <-chan int, output chan<- int, reducer Reducer, init int /* more channels here, if required */){
  /* add reduce functionality here */
}


/* add code for shutdown functionality here, if required */


type EvenOddMapper struct {}
type EvenOddReducer struct {}

/* implement your interfaces here */


func main() {
	input := make(chan int)
	inter1 := make(chan int)
	inter2 := make(chan int)
  output := make(chan int)

/* channels and processes for shutdown here */

    go Map(input, inter1, inter2, EvenOddMapper{} /* parameters here, if needed */)
    go Map(input, inter1, inter2, EvenOddMapper{} /* parameters here, if needed */)
    go Reduce(inter1, output, EvenOddReducer{},0)
    go Reduce(inter2, output, EvenOddReducer{},0)
    input <- 1
    input <- 2
    input <- 3
    input <- 4
    input <- 5
/* initiate shutdown here */
    res := <-output + <-output
    fmt.Println(res) //should be 55 for example
    close(output)
}
