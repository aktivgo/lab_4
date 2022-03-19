package main

import (
	"lab_4/internal/task_3"
	"log"
)

func main() {
	if err := task_3.Run(); err != nil {
		log.Fatalln(err)
	}
}
