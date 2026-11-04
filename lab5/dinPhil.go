// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24

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

// Modified By: Adam Noonan
// Issues: none

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// sleep for random 1-5 seconds, return amount slept, in seconds
func waitRandomTime() int {
	var X time.Duration
	var randNum = rand.IntN(5)
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)
	return randNum
}

// wait with no forks acquired
func think(index int) {
	var time = waitRandomTime()
	fmt.Println("Philosopher", index, "thought for", time, "seconds")
}

// wait 2 forks acquired
func eat(index int) {
	var time = waitRandomTime()
	fmt.Println("Philosopher", index, "ate for", time, "seconds")
}

// try to acquire 2 forks, block if unavailable
func getForks(index int, forks map[int]chan bool) {
	if index == 0 { // first philosopher takes forks in different order, prevents deadlock
		forks[(index+1)%5] <- true
		forks[index] <- true
	} else {
		forks[index] <- true
		forks[(index+1)%5] <- true
	}
}

// return forks when done eating, order doesn't matter
func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
}

// create goroutine to loop: think, access forks, eat, and return forks
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for { // think and eat forever
		think(index)
		getForks(index, forks) // blocks if relevant forks arent available
		eat(index)
		putForks(index, forks) // return forks
	}
	wg.Done()
}

// main
func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1) // set up forks
	}
	for N := range philCount {
		go doPhilStuff(N, &wg, forks) // start philosophers
	}
	wg.Wait() // main thread wait here until everyone (5 go routines) is done (will never happen)

}
