package main

import (
	"fmt"
	"sort"
)


func saralash(sonlar []int, ch chan []int) {
	sort.Ints(sonlar) 
	ch <- sonlar      
}

func main() {
	sonlar := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	ch := make(chan []int) 

	go saralash(sonlar, ch) 


	saralganSonlar := <-ch

	fmt.Println("Berilgan sonlar:", sonlar)
	fmt.Println("Saralgan sonlar:", saralganSonlar)
}
