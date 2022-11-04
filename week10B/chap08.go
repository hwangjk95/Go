/*structure*/
package main

import (
	"fmt"

	"magazine"
)

func main() {
	var s1 magazine.Subscriber //subscriber 의 첫번째 글자가 대문자여야지 public
	s2 := magazine.Subscriber{Rate: 4.99}

	s2.Name = "최인하"
	addr := magazine.Address{Street: "안산천동로", City: "안산시", State: "경기도"}
	s2.Address = addr
	fmt.Printf("%s의 요금은 %.2f달러 입니다. 사는 동네는 %s\n", s2.Name, s2.Rate, s2.Address)

	s1.Name = "김인하"
	s1.Street = "인하로"
	s1.City = "용현동"
	s1.State = "인천시"
	s1.Rate = 5.99
	s1.Active = true

	fmt.Printf("%s의 요금은 %.2f 달러입니다. 사는 동네는 %s\n", s1.Name, s1.Rate, s1.City)
	fmt.Println(s2.Name, s2.Rate, s2.Active)

}
