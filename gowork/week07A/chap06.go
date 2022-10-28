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

// // func GetFloat(fn string) ([3]float64, error) {
// func GetFloat(fn string) ([]float64, error) {
// 	//var number [3]float64
// 	var numbers []float64
// 	f, err := os.Open(fn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	sc := bufio.NewScanner(f)
// 	//i := 0
// 	for sc.Scan() {
// 		//numbers[i], err = strconv.ParseFloat(sc.Text(), 64)
// 		number, err := strconv.ParseFloat(sc.Text(), 64)
// 		if err != nil {
// 			return nil, err
// 		}
// 		numbers = append(numbers, number)
// 		//i++
// 	}
// 	err = f.Close()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if sc.Err() != nil {
// 		return nil, sc.Err()
// 	}
// 	return numbers, nil
// }
// func main() {
// 	numbers, err := GetFloat("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var sum float64
// 	for _, number := range numbers {
// 		fmt.Println(number)
// 		sum = sum + number
// 	}
// 	fmt.Printf("평균 : %.2f\n", sum/float64(len(numbers)))
// }

// package main

// import "fmt"

// func main() {
// 	d := make([]int, 3)
// 	//var e []int
// 	// fmt.Println(d, len(d))
// 	// fmt.Println(e, len(e))
// 	// fmt.Printf("%#v %#v\n", d, e)
// 	d[1] = 7
// 	fmt.Println(d, len(d))
// 	d = append(d, -9, 19)
// 	//fmt.Println(d, len(d))
// 	d = append(d, 3, 9, 7)
// 	//fmt.Println(e, len(e))
// 	d = append(d, 3, 9, 7)
// 	//fmt.Println(f, len(f))
// 	d[2] = 999
// 	fmt.Println(d)
// 	// a := []int{1, 3, -9, 7}
// 	// b := a[1:4]
// 	// fmt.Println(b)
// 	// //c := a[1:5] // 파이썬과 다르기 런타임 panic 발생
// 	// //fmt.Println(c)
// 	// c := a[1:] // 파이썬과 동일
// 	// fmt.Println(c)
// }
