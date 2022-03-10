package app

import (
	"fmt"
	"lab_4/internal"
	"sync"
)

func Run() error {
	var countN int
	fmt.Print("Введите количество заявок: ")
	_, _ = fmt.Scanf("%d\n", &countN)

	var queueSize int
	fmt.Print("Введите размер очереди: ")
	_, _ = fmt.Scanf("%d\n", &queueSize)

	var writerSpeed int
	fmt.Print("Введите скорость работы производителя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &writerSpeed)

	var readerSpeed int
	fmt.Print("Введите скорость работы потребителя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &readerSpeed)

	queue := internal.NewQueue(queueSize, countN, 0)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		queue.Increment(writerSpeed)
	}()

	go func() {
		defer wg.Done()
		queue.Decrement(readerSpeed)
	}()

	wg.Wait()

	fmt.Println("Программа завершила работу")

	return nil
}