package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/sukhvir148/go_service/src/pub"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")

	scanner := bufio.NewScanner(os.Stdin)

	zero := 0
	acc := Account{
		money: &zero,
	}

	quote.Hello()

	wg := &sync.WaitGroup{}

	creditor := func(wg *sync.WaitGroup) {
		acc.credit()
		wg.Done()
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go creditor(wg)
	}

	wg.Wait()

	fmt.Printf("FINAL AMOUNT - %d\n", *acc.money)

	done := make(chan int)
	p := pub.NewProducer(make(chan string, 10), done, []pub.Consumer{pub.Corporate{}, pub.Personal{}})
	// p := Producer{
	// 	queue: make(chan string, 10),
	// 	done:  done,
	// 	consumers: []Consumer{
	// 		Corporate{},
	// 		Personal{},
	// 	},
	// }

	go p.Process()

	for {
		fmt.Print("Enter text: ")
		scanner.Scan()
		event := scanner.Text()

		p.Enqueue(event)

		if event == "stop" {
			done <- 100
			break
		}

		fmt.Println("You entered:", event)
	}

}
