package random_practice

import (
	"fmt"
	"sync"
	"time"
)

// producer

func SimpleProducerConsumer() {
	// we need a shared channel to coordinate
	// and we need to run the 2 go-routines in async way
	var wg sync.WaitGroup

	var queue = make(chan string)

	wg.Add(1)
	go producer(queue, &wg)

	wg.Add(1)
	go consumer(queue, &wg)

	//waiting...
	wg.Wait()
}

func producer(queue chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	queue <- "Producing : Hi Mayank..."
	fmt.Println("Producing done ...")
}

func consumer(queue chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case msg := <-queue:
		time.Sleep(time.Second * 2)
		fmt.Print("printing msg -> ", msg)
	case <-time.After(time.Second * 5):
		fmt.Print("5 sec elapsed - exiting consumer")
	}
}

//////// Other way of doing it ..

func DriverProducerConsumer() {
	text := make(chan int)
	done := make(chan bool)

	prod := Producer{
		text: &text,
		done: &done,
	}

	consum := Consumer{
		text: &text,
	}

	go prod.produce(5)
	go consum.consume()

	// receiving signal on done...  which indicates production is done.
	<-done
	fmt.Println("Production is done.")

}

type Producer struct {
	text *chan int

	// indicator for done producing.
	done *chan bool
}

type Consumer struct {
	text *chan int
}

func (p *Producer) produce(max int) {

	for i := 0; i < max; i++ {
		*p.text <- i
	}

	// done producing
	*p.done <- true
}

func (c *Consumer) consume() {
	for {
		msg := <-*c.text
		fmt.Println("Message received by the consumer", msg)
	}

}
