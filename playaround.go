package main

import (
	"fmt"
	"math/rand"
	"time"
)


func producer(c chan int) {
  for {
   rand.Seed( time.Now().UnixNano() )
   number := rand.Intn( 13375 )

   c <- number
   time.Sleep(2 * time.Second)
  }
}

func consumer(c_in chan int, c_out chan int) {
  for {
    foo := <- c_in
    c_out <- foo
  }
}

func logger(c chan int) {
  for {
    foobar := <- c
    fmt.Println(foobar)
  }
}

func myChannelExperiments() {
	channel_producer := make(chan int)
        channel_forworder := make(chan int)

	go producer(channel_producer)
	go consumer(channel_producer, channel_forworder)
	go logger(channel_forworder)

}

func main() {
  myChannelExperiments()
  // Hang forever.
  select {} // join
}
