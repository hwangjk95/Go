// package main

// import "gadget"

// func playList(device gadget.TapePlayer, songs []string) {
// 	for _, song := range songs {
// 		device.Play(song)
// 	}
// 	device.Stop()
// }

// func main() {
// 	player := gadget.TapePlayer{}
// 	mixtape := []string{"좋은날", "너랑나", "투명우산"}
// 	playList(player, mixtape)

// }

package main

import "gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mcrtape := []string{"혜성", "비밀번호486", "별의조각"}
	mixtape := []string{"좋은날", "너랑나", "투명우산"}
	playList(player, mixtape)
	playList(player, mcrtape)
}

//추상클래스와 인터페이스 개념(시험x 개념중요)
//덕타이핑
