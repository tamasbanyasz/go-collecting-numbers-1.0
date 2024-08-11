package main

/*
The code creates a determined sized array and selected randomly a name for this array from a list with letters("A" or "C" or "H").

Then the code randomly puts "1"'s on different inexes of these arrays. And there is a collector ("5") what will gathering these numbers of "1".

if the collected numbers lenght is odd the code will store the array name, the numbers, and the indexes where were these numbers into an another struct.
This process is run in GoRoutines

That what will you see at the end when the code is finished.


*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// #Create row for the matrix
type makingRow struct {
}

func (makingRow) row(some *[]int) []int {

	var store []int

	return append(*some, store...)
}

type createRow struct {
	a makingRow
}

//-------------------------------------------

// #Create matrix
type makingMatrix struct {
}

type indexesOfNums struct {
}

func (makingMatrix) matrix(size int) [][]int {

	// Create empty array with the lenght
	array := make([][]int, size)

	// Call the create struct
	create := createRow{}

	for k := 0; k != size; k++ {
		// A random row with elements
		randomRow := []int{}

		for j := 0; j != size; j++ {
			value := rand.Intn(2)
			var eh *[]int = &randomRow
			randomRow = append(*eh, value)
		}
		// Create rows by the random list with it's elements

		array[k] = create.a.row(&randomRow)
	}

	return array

}

func (indexesOfNums) indexesOfTheNumbers(arr [][]int) [][]int {

	var idxs [][]int

	for i := 0; i != len(arr); i++ {
		for j := 0; j != len(arr[i]); j++ {
			if arr[i][j] == 1 {
				idxs = append(idxs, []int{i, j})
			}
		}
	}

	return idxs
}

type Matrix struct {
	a   makingMatrix
	get indexesOfNums
}

//--------------------------------------------------------

type movement struct {
}

func (movement) movementOfTHeCollector(arr [][]int) []int {

	var collectorsRow int
	var collectorsColumn int

	var catch int

	var items []int

	myBool := true

	for myBool {

		// Print the matrix
		fmt.Println("The array:")
		for _, row := range arr {
			for _, val := range row {
				fmt.Print(val, "\t")
			}
			fmt.Println()
		}

		if collectorsColumn > 0 {
			arr[collectorsRow][collectorsColumn-1] = 0
		}

		for i := 0; i != len(arr); i++ {
			for j := 0; j != len(arr[i]); j++ {
				if arr[collectorsRow][collectorsColumn] == 1 {

					items = append(items, arr[collectorsRow][collectorsColumn])
					arr[collectorsRow][collectorsColumn] = 0
					catch++
				}
			}
		}

		arr[collectorsRow][collectorsColumn] = 5
		fmt.Println(items)
		if collectorsRow == len(arr)-1 && collectorsColumn == len(arr)-1 {
			myBool = false
		}

		if collectorsRow > 0 && collectorsColumn == 0 {
			x := collectorsRow
			arr[x-1][len(arr)-1] = 0
		}

		if collectorsColumn < len(arr) {
			collectorsColumn++

		}

		if collectorsColumn == len(arr) {
			collectorsColumn = 0
			collectorsRow++
		}

	}
	fmt.Println("The array:")
	for _, row := range arr {
		for _, val := range row {
			fmt.Print(val, "\t")
		}
		fmt.Println()
	}

	fmt.Println(catch)

	return items
}

type moving struct {
	move movement
}

type Arrays struct {
	Name     string
	elements []int
	indexes  [][]int
}

func dowork(name string, wg *sync.WaitGroup, finalArrays map[int]Arrays, randint int, lock *sync.Mutex) {
	fmt.Println("Start....")

	create := Matrix{}
	movementProcess := movement{}

	arr := create.a.matrix(randint)

	idxs := create.get.indexesOfTheNumbers(arr)

	fmt.Println(idxs)

	collectorsMovement := movementProcess.movementOfTHeCollector(arr)

	createFinalArray := Arrays{Name: name, elements: collectorsMovement, indexes: idxs}

	fmt.Println("End....")
	fmt.Println("============================================================================")
	fmt.Println("Odded arrays: ")

	if len(createFinalArray.elements)%2 == 0 && len(createFinalArray.elements) != 0 {

		lock.Lock()
		finalArrays[len(finalArrays)] = createFinalArray
		lock.Unlock()

	}
	wg.Done()
}

func main() {
	start := time.Now()
	names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "Z", "v", "W"}

	finalArrays := make(map[int]Arrays)

	var lock sync.Mutex
	wg := sync.WaitGroup{}

	wg.Add(10)
	for i := 0; i != 10; i++ {
		randomIndex := rand.Intn(len(names))
		pickName := names[randomIndex]

		go dowork(pickName, &wg, finalArrays, rand.Intn(6)+2, &lock)

	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("It was :  %s\n", elapsed)
	for _, k := range finalArrays {
		fmt.Println(k)
	}

}
