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
// 	//var numbers [3]float64
// 	var numbers []float64
// 	f, err := os.Open(fn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//i := 0
// 	sc := bufio.NewScanner(f)
// 	for sc.Scan() {
// 		//numbers[i], err = strconv.ParseFloat(sc.Text(), 64)
// 		number, err := strconv.ParseFloat(sc.Text(), 64)
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
// 	if sc.Err() != nil {
// 		return nil, err
// 	}
// 	return numbers, nil
// }

// func main() {
// 	weeks, err := GetFloat("data.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var sum float64
// 	for _, week := range weeks {
// 		fmt.Println(week)
// 		sum = sum + week
// 	}
// 	fmt.Printf("평균 : %.2f 킬로그램\n", sum/float64(len(weeks)))
// }

// package main

// import "fmt"

// func main() {
// 	var a []int         // 슬라이스 변수, 메모리 할당 전
// 	b := make([]int, 3) // 메모리 할당 완료된 슬라이스 변수
// 	var c [2]int        // 배열 변수
// 	d := [2]int{55, 99} // 슬라이스 리터럴로 초기화
// 	fmt.Printf("%#v %d\n", a, len(a))
// 	fmt.Printf("%#v %d\n", b, len(b))
// 	fmt.Printf("%#v\n", c)
// 	fmt.Printf("%#v\n", d)
// }
