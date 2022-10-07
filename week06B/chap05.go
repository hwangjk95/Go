// package main

// import "fmt"

//chap03 pointer
// import "fmt"

// func printPointer(myBoolPointer *bool) {
// 	fmt.Println(*myBoolPointer)
// }

// func main() {
// 	myBool := true
// 	printPointer(&myBool)
// }

// chap05 Array,Array literal
package main

// import "fmt"

// func main() {
// 	subjects := []string{
// 		"오픈소스프로그래밍",
// 		"운영체재",
// 		"시스템분석및설계",
// 		"자바프로그래밍",
// 	}

// for _, subject := range subjects {
// 	fmt.Println(subject)
// }

// for i, subject := range subjects {
// 	fmt.Println(i, subject)
// }

// for i := 0; i < len(subjects); i++ {
// 	fmt.Println(i, subjects[i])
// }

// for i := 0; i <= 3; i++ {
// 	fmt.Println(subjects[i])
// }
//}
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloat(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err

	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil
}

func main() {
	//파일 입출력 기본
	// file, err := os.Open("./data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// err = file.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if scanner.Err() != nil {
	// 	log.Fatal(scanner.Err())
	// }

	weeks, err := GetFloat("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, week := range weeks {
		sum = sum + week
	}
	fmt.Printf("평균 : %.2f 킬로그램\n", sum/float64(len(weeks)))
}

//배열의 길이가 고정되는 문제점 >> slice 활용!!
