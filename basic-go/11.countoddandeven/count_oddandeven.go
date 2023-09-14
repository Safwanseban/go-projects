package main

import (
	"basic-go/utils"
	"fmt"
	"log"
	"strconv"
)

func main() {

	elements := make([]int, 0)
	var (
		min, max, mean int
	)
	for {

		data, err := utils.GetInput()
		if err != nil {
			log.Fatal("error reading data")
		}
		if data == "proceed" {
			break
		} else {
			toInt, err := strconv.ParseInt(data, 10, 64)
			if err != nil {
				log.Fatal("not a valid integar")
			}
			elements = append(elements, int(toInt))

		}

	}
	mean, min, max = findMinMaxandMean(elements)
	fmt.Println("enter which operation you want\nmin\nmax\ncount\nmean ")
	input, _ := utils.GetInput()

	switch input {

	case "mean":
		fmt.Println("mean of the elements is ", mean)
	case "min":
		fmt.Println("min of the element is ", min)
	case "max":
		fmt.Println("max of the element is ", max)
	case "count":
		fmt.Println("count of the element is ", len(elements))
	default:
		fmt.Println("entered operation is invalid", input)
	}
	fmt.Println("No of odds and evens available in the list")
	odd, even := CountEvenOdd(elements)
	fmt.Printf("no of odd numbers is %d\nno of even numbers is %d\n", odd, even)

}

// findMinMaxandMean finds min,max and mean in a list of elements
func findMinMaxandMean(elements []int) (int, int, int) {
	min, max, sum := elements[0], elements[0], 0

	for _, num := range []int(elements) {
		sum += num
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return sum / len(elements), min, max

}

// CountEvenOdd counts odd and even count of elements in a given list
func CountEvenOdd(nums []int) (odd int, even int) {

	for _, v := range []int(nums) {
		if v%2 != 0 {
			odd++
		} else {
			even++
		}

	}
	return odd, even

}
