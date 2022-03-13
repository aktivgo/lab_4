package main

import (
	"lab_4/internal/task_2"
	"log"
)

func main() {
	if err := task_2.Run(); err != nil {
		log.Fatal(err)
	}
}
