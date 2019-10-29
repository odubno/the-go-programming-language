package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency uses gorutines to calculate word frequency in text
func ConcurrentFrequency(strings []string) FreqMap {
	// wg waits for a group of goroutines to finish
	var wg = sync.WaitGroup{}
	ch := make(chan FreqMap)
	defer close(ch)

	// SENDER
	// loop over slice of strings, calulate frequency and send results to channel (ch)
	// Add() sets the number of goroutines to wait for
	wg.Add(len(strings))
	for i := 0; i < len(strings); i++ {
		data := strings[i]
		go func(s string) {
			// ch <- sends data to channel
			ch <- Frequency(data)
			// goroutine has completed
			defer wg.Done()
		}(data)
	}

	// RECEIVER
	// loop the same amount of times as the sender, request results from sender, copy results to new map
	result := FreqMap{}
	for i := 0; i < len(strings); i++ {
		// <-ch receives data from channel
		m := <-ch
		// copy results from m
		for k, v := range m {
			result[k] += v
		}

	}

	// Wait() blocks until all add()ed goroutines are done()
	// this ensures that function does not return the result before the goroutine is complete
	wg.Wait()
	return result
}
