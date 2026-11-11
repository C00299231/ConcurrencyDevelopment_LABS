// Producer Consumer code
// Author: Adam Noonan
// Created: 05/10/2026

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

var numThreads = 10
var bufferTotal = 10

// create a worker to consume forever
func consumer(consumerIdx int, buffer chan int) {
	for {
		data := <-buffer // accept data from the channel (blocking)
		fmt.Println("consumer", consumerIdx, "consumed", data)
	}
}

// create a worker to produce forever
func producer(producerIdx int, buffer chan int) {
	for {
		data := rand.IntN(25) // create random number to send into the channel
		fmt.Println("producer", producerIdx, "produced", data)
		buffer <- data
	}
}

// main
func main() {

	buffer := make(chan int, bufferTotal)
	var wg sync.WaitGroup
	wg.Add(numThreads)

	// create many consumers
	for i := range numThreads {
		go consumer(i, buffer)
	}
	go producer(0, buffer) // create one producer (maybe can create more?)

	// producer & consumers run forever, so wg waits forever
	wg.Wait()
}
