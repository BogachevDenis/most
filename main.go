package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"sort"
)


func main() {
	var writers, arrSize, iterCount int

    fmt.Print("Введите количество горутин: ")
	fmt.Fscan(os.Stdin, &writers) 
	
    fmt.Print("Введите размер массива: ")
	fmt.Fscan(os.Stdin, &arrSize)
	
	fmt.Print("Введите количестов итераций: ")
	fmt.Fscan(os.Stdin, &iterCount)

	fmt.Println("{writerGoroutineIdenitifier} {queueInsertionTime} {min} {median} {max}")

	queue := initQ()
	finish := make(chan bool) 
	go queue.work(finish, writers*iterCount)

	rand.Seed(time.Now().Unix())
	for i := 0; i < writers; i++ {
		go createArr(i, iterCount, arrSize, queue)
	}
	<- finish
}

func createArr(writer, iterCount, arrSize int, q *queue) {
	for i := 0; i < iterCount; i++ {
		arr := make([]int, 0)
		for k := 0; k < arrSize; k++ {
			arr = append(arr, rand.Intn(100-1) + 1) // (max - min) + min
		}
		t := time.Now().Format("15:04:05.000")
		q.push(data{writer, t, arr})
	}
}

type queue struct {
	ch	chan data
}

type data struct {
	id		int
	time	string
	arr		[]int
}

func initQ() *queue {
	ch := make(chan data) 
    return &queue{ch}
} 

func (q *queue) push(val data) {
	q.ch <- val
}

func (q *queue) work(ch chan bool, count int) {
	for i := 0; i < count; i++ {
		task := <- q.ch
		sort.Ints(task.arr)
		fmt.Println("goroutineId - ", task.id," ; time - ",  task.time, " ; min - ", task.arr[0], " ; median - ", task.arr[len(task.arr)/2], " ; max - ", task.arr[len(task.arr) - 1])
	}
	ch <- true
}
