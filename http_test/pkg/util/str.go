package util

import (
	"fmt"
	// "math/rand"
	"crypto/rand"
	// "time"
)

var (
	maxRoom = 16
	maxDot  = 16
)

func GetFlatStatus() string {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("some err")
	}
	s := fmt.Sprintf("%02X", b[0])
	return s
}

func GetStatus(n int) string {
	res := ""
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("some err")
	}
	for i := 0; i < n; i++ {
		s := fmt.Sprintf("%02X", b[i])
		res += s
	}
	// fmt.Printf("room: %s", res)
	return res
}

func GetStrFlat() string {
	// "src": "flat=02&room=21451632&dot=01"
	flat := GetStatus(1)
	room := GetStatus(maxRoom)
	dot := GetStatus(maxDot * maxRoom)
	res := fmt.Sprintf("flat=%s&room=%s&dot=%s", flat, room, dot)
	return res
}
