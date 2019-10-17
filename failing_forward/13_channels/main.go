/*
Channles and how they're used to synchronize data transmition between go routines

Channel basics
	create a channel with "make" command
		make(chan int)
			channel could be any time that receiver expects and receiver sends
			strongly typed
	send message into channel
		ch <- val
			position of the arrow kinda indicates wher the data will flow
			arrow is pointing into the channel
	receive message from channel
		val := <-ch
			the arrow is leading out of the channel
			the arrow is before the channel
	can have multiple senders and receivers
		very common for one channel to be distributed among multiple go routines
			multiple data senders and receivers

Restricting data flow
	bidirectional data flow
	Channel can be cast into send-only or receive only versions
		handle channel data in one direction; specify the direction
		send-only: chan <- int
		receive-only: <-chan int

Buffered channels
	Channels block sender side till receiver is available
	Block receiver side till message is available
	Can decouple sender and receiver with buffered channels
		make(chan int, 50)
	Use buffered channels when send and receiveer have assymetric loading

For...range loops with channels
	Use to monitor channel and process messages as they arrive
	Loop exists when channel is closed

Select statement
	work kinda like switch statements, but they work only in the context of channels
	Allows goroutine to monitor several channels at once
		blocks if all channels block
		if multiple channels receive value simultaneously, behavior is undefined
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// to help synchronize data flow
var wg = sync.WaitGroup{}

func main() {
	// 01. Channel Basics
	// channelBasics()
	// channelBasics2()
	// channelBasics3()

	// 02. Restricting data flow
	// restrictingChannel()

	// 03. Buffered Channels
	// bufferedChannels()

	// 04. Closing channels
	// closingChannels()

	// 05. For...range loops with channels
	// forChannels()

	// 06. Select statements
	selectStatements()
}

func channelBasics() {
	// channels always work in the context of go routines
	// channels help synchronize data transmition between multiple go routines
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch // receives data from channel
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i // sending data to channel
		i = 27  // this does update the variable; the value was already sent to the channel
		wg.Done()
	}()
	wg.Wait()
}

func channelBasics2() {
	ch := make(chan int)
	for j := 0; j < 5; j++ {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()
		go func() {
			i := 42
			ch <- i
			i = 27
			wg.Done()
		}()
	}
	wg.Wait()
}

func channelBasics3() {
	// both channels act as receivers and senders
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch // receives data from channel
		fmt.Println(i)
		ch <- 27 // sending data to channel
		wg.Done()
	}()
	go func() {
		ch <- 42          // sending data to channel
		fmt.Println(<-ch) // printing out received data from channel
		wg.Done()
	}()
	wg.Wait()
}

func restrictingChannel() {
	// send and receive only channels
	// this is much cleaner

	ch := make(chan int)
	wg.Add(2)

	// receive only channel
	// data is flowing out of the channel i.e. "<-chan"
	go func(ch <-chan int) {
		i := <-ch // receives data from channel
		fmt.Println(i)
		wg.Done()
	}(ch)

	// send only channel
	// data is sent out to the channel i.e. "chan->"
	go func(ch chan<- int) {
		ch <- 42 // sending data to channel
		wg.Done()
	}(ch)
	wg.Wait()
}

func bufferedChannels() {
	// adding a buffer of 50
	// tells go to create a channel that has an internal data store
	// that could store up to 50 integers
	ch := make(chan int, 50)

	wg.Add(2)

	// receive only channel
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// send only channel
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}

func closingChannels() {
	ch := make(chan int, 50)

	wg.Add(2)

	// receive only channel
	go func(ch <-chan int) {
		// monitor for continous massaging
		// continously monitor for massaging
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	// send only channel
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		// need to close channel so that the for
		// range loop knows to stop waiting on more messages
		// the close command gets sent to the receiver, the loop reads it in and exists
		close(ch)
		// ch <- 23 // can't send a message into a closed channel
		wg.Done()
	}(ch)
	wg.Wait()
}

func forChannels() {
	ch := make(chan int, 50)

	wg.Add(2)

	// receive only channel
	go func(ch <-chan int) {
		// functionally the same a the for range construct
		// explicitly seeing the "ok" i.e. close or not
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				// channel is closed; no more messages
				break
			}
		}
		wg.Done()
	}(ch)

	// send only channel
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func selectStatements() {
	go logger2()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 * time.Microsecond)

	// initialize empty struct
	doneCh <- struct{}{}
}

func logger() {
	for entry := range logCh {
		fmt.Printf("[%v]%v\n", entry.severity, entry.message)
		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
}

func logger2() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
