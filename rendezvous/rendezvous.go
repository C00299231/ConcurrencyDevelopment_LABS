//Rendezvous.go Template Code
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

// --------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Adam Noonan C00299231
// 18/09/2026
// Students Helped: Amelia Hamulewicz, Mark Lambert, Dorian Nowacki, Ariel Fajimiyo
// -------------------------------------------

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup,
	Num int,
	aArrived chan struct{},
	bArrived chan struct{}) bool {

	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	aArrived <- struct{}{} // send into A channel

	<-bArrived // get from B channel

	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	threadCount := 5

	var aArrived = make(chan struct{})
	var bArrived = make(chan struct{})

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, aArrived, bArrived)
	}

	for range threadCount { // wait for signal from each grt
		<-aArrived
	}

	// allow grts to do part B when all have done part A
	for range threadCount {
		bArrived <- struct{}{}
	}

	wg.Wait() //wait here until everyone (10 go routines) is done

}
