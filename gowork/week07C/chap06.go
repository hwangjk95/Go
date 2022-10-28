package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// )

// func GetFloat(fn string) ([]float64, error) {
// 	//var numbers [3]float64
// 	var numbers []float64
// 	f, err := os.Open(fn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//i := 0
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		//numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
// 		number, err := strconv.ParseFloat(scanner.Text(), 64)
// 		if err != nil {
// 			return nil, err
// 		}
// 		//i++
// 		numbers = append(numbers, number)
// 	}
// 	err = f.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if scanner.Err() != nil {
// 		return nil, scanner.Err()
// 	}
// 	return numbers, nil
// }
// func main() {
// 	numbers, err := GetFloat("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var sum float64 = 0
// 	for _, number := range numbers {
// 		fmt.Println(number)
// 		sum = sum + number
// 	}
// 	fmt.Printf("평균 : %0.2f\n", sum/(float64(len(numbers))))
// }
