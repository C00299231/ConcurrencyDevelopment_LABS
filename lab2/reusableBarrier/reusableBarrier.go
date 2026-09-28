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
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(index int, count *int, total int, inner *semaphore.Weighted, outer *semaphore.Weighted, mut *sync.Mutex, ctx *context.Context) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", index)

	mut.Lock()
	*count++
	var current = *count
	mut.Unlock()

	if current == total {
		outer.Acquire(*ctx, 1)
		inner.Release(1)
	} else {
		inner.Acquire(*ctx, 1)
		inner.Release(1)
	}

	fmt.Println("PartB", index)

	mut.Lock()
	*count--
	current = *count
	mut.Unlock()

	if current == 0 {
		inner.Acquire(*ctx, 1)
		outer.Release(1)
	} else {
		outer.Acquire(*ctx, 1)
		outer.Release(1)
	}
	fmt.Println("PartC", index)
	return true
}

func main() {
	fmt.Println("program did run")
	count := 0
	total := 3
	var ctx = context.TODO()
	var innerSem = semaphore.NewWeighted(int64(1))
	var outerSem = semaphore.NewWeighted(int64(1))
	innerSem.Acquire(ctx, 1)

	var mut sync.Mutex
	for i := range total {
		go doStuff(i, &count, total, innerSem, outerSem, &mut, &ctx)
	}
	time.Sleep(time.Second * 10)
}
