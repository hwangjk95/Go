package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	quest := [5]string{
		"이거 왜이래? 나 인하컴정 나온 여자야",
		"내가 빙다리 핫바지로 보이냐?",
		"늑대새끼가 어떻게 개 밑으로 들어갑니까?",
		"손은 눈보다 빠르다!",
		"내가 괜한 얘길 했나? 이 돈은 안받는걸로",
	}

	fmt.Println("<<< 타짜 >>>")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for success == false {

		seconds := time.Now().Unix()
		rand.Seed(seconds)
		target := rand.Intn(5)

		fmt.Println("Q : " + quest[target])
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		correct := [5]string{
			"정마담",
			"아귀",
			"고니",
			"평경장",
			"너구리",
		}
		input = strings.TrimSpace(input)

		if input == correct[target] {
			fmt.Println("정답입니다!")
			success = true
		} else {
			seconds2 := time.Now().Unix()
			rand.Seed(seconds2)
			wrongSeed := rand.Intn(3)

			wrong := [3]string{
				"너, 문제 맞출 수 있겠냐?",
				"또 못 맞추면 넌 변사체가 된다.",
				"어이 컴정 친구, 학생 답게 맞춰봐",
			}

			fmt.Println(wrong[wrongSeed])
			fmt.Println("===========================")
		}
	}
}
