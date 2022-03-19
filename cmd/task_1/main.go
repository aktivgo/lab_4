package main

import (
	"lab_4/internal/task_1"
	"log"
)

func main() {
	if err := task_1.Run(); err != nil {
		log.Fatalln(err)
	}
}
