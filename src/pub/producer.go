package pub

import (
	"fmt"
	"time"
)

type Producer struct {
	queue     chan string
	done      chan int
	consumers []Consumer
}

func NewProducer(queue chan string, done chan int, consumers []Consumer) *Producer {
	return &Producer{
		queue:     queue,
		done:      done,
		consumers: consumers,
	}
}

func (p *Producer) Add(i, j int) int {
	return i + j
}

func (p *Producer) Enqueue(event string) {
	p.queue <- event
}

func (p *Producer) Process() {

	for {
		select {
		case val := <-p.queue:
			for _, con := range p.consumers {
				con.Process(val)
			}
		case <-p.done:
			fmt.Println(".... EXITING")
			break
		default:
			fmt.Println("............")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

type Consumer interface {
	Process(event string)
}

type Corporate struct{}

func (Corporate) Process(event string) {
	fmt.Println("Processing corporate event in other package: " + event)
}

type Personal struct{}

func (Personal) Process(event string) {
	fmt.Println("Processing personal event : " + event)
}
