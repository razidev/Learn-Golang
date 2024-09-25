package helper

import "time"

func GiveMeResponse(channel chan string, name string) {
	time.Sleep(2 * time.Second)
	channel <- name
}
