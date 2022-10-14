package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// func GetFloat(fileName string) ([3]float64, error) {
func GetFloat(fileName string) ([]float64, error) {
	//var numbers [3]float64
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err //슬라이스는 nil값이 될 수 있어서 nil 반환 가능
	}
	//i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)

		number, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}
		//i++
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err

	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}

func main() {

	//p220 참고

	// a := []int{0, 55, 0, 0, 0}
	// c := append(a, 99)
	// d := append(c, 1, 2, 3)
	// e := append(d, 11)
	// fmt.Println(a, c, d, e)
	// e[1] = 30
	// fmt.Println(a, c, d, e)

	// z := []int{0, 55, 0, 0, 0}
	// z = append(z, 99)
	// z = append(z, 1, 2, 3)
	// z = append(z, 11)
	// z[1] = 30
	// fmt.Println(z)

	// var a []int         //슬라이스 변수
	// b := make([]int, 3) //메모리 할당 완료된 슬라이스 변수
	// var c [2]int        //배열 변수
	// d := [2]int{55, 99} //슬라이스 리터럴로 초기화
	// fmt.Printf("%#v %d\n", a, len(a))
	// fmt.Printf("%#v %d\n", b, len(b))
	// fmt.Printf("%#v\n", c)
	// fmt.Printf("%#v\n", d)

	weeks, err := GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, week := range weeks {
		fmt.Println(week)
		sum = sum + week
	}
	fmt.Printf("평균 : %.2f 킬로그램\n", sum/float64(len(weeks)))
	//6주차 내용 개선
}
