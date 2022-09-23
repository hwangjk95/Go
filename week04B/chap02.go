// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	//var now time.Time = time.Now()
// 	now := time.Now()
// 	//var year int = now.Year()
// 	year := now.Year()
// 	fmt.Println(year)
// 	fmt.Println(now.Month())

// 	broken := "G# r#cks!"
// 	replacer := strings.NewReplacer("#", "o")
// 	fixed := replacer.Replace(broken)
// 	fmt.Println(fixed)

//챕터1

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	// fmt.Print("Enter a grade: ")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// input, _ := reader.ReadString('\n') //강제적으로 에러 무시 "_"=입력값 무시
// 	// fmt.Println(input)

// 	fmt.Print("Enter a grade: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, err := reader.ReadString('\n')
// 	//fmt.Println(err)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	input = strings.TrimSpace(input)
// 	score, err := strconv.ParseFloat(input, 64)
// 	if err != nil {
// 		log.Fatal(err) //2022/09/23 14:35:56 strconv.ParseFloat: parsing "50\r\n": invalid syntax
// 	}
// 	fmt.Println(input)
// 	var status string
// 	if score >= 60 {
// 		status = "합격!"
// 		//fmt.Println(status)
// 	} else {
// 		status = "불합격!"
// 		//fmt.Println(status)
// 	}
// 	fmt.Println(score, "점은", status, " 입니다.")

// }

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	//fmt.Println(target)
	fmt.Println("1에서 100사이의 수를 추측하세요~")

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("남은기회는 ", 10-guesses, " 번 입니다.")
		fmt.Print("점수입력: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			success = true
			fmt.Println("정답입니다. 축하해요")
			break
		} else if guess > target {
			fmt.Println("정답이 입력값보다 작습니다.")
		} else {
			fmt.Println("정답이 입력값보다 큽니다.")
		}
	}
	if !success {
		fmt.Println("죄송합니다. 당신은 기회를 모두 소진하였습니다. 정답:", target)
	}
}
