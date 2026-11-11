//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

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

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Adam Noonan C00299231
// Issues: none
// Students Helped: Amelia Hamulewicz, Mark Lambert, Dorian Nowacki, Ariel Fajimiyo
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

var barrier = make(chan struct{}, 1)
var count int = 0
var total int = 10

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, theLock *sync.Mutex) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	theLock.Lock()
	count++
	if count == total {
		barrier <- struct{}{}
	}
	theLock.Unlock()

	// we wait here until everyone has completed part A
	<-barrier
	barrier <- struct{}{} // release cascades down

	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {

	var wg sync.WaitGroup
	wg.Add(total)

	// we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(total))
	sem.Acquire(ctx, 1)

	for i := range total { // create the go Routines here
		go doStuff(i, &wg, &theLock)
	}

	sem.Release(1)

	wg.Wait() // wait for everyone to finish before exiting
}
