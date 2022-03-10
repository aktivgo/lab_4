package app

import (
	"fmt"
	"lab_4/internal"
	"sync"
)

func Run() error {
	var countApplication int
	fmt.Print("Введите количество заявок: ")
	_, _ = fmt.Scanf("%d\n", &countApplication)

	var queueSize int
	fmt.Print("Введите размер очереди: ")
	_, _ = fmt.Scanf("%d\n", &queueSize)

	var writerSpeed int
	fmt.Print("Введите скорость работы производителя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &writerSpeed)

	var readerSpeed int
	fmt.Print("Введите скорость работы потребителя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &readerSpeed)
	fmt.Println()

	queue := internal.NewQueue(queueSize, countApplication)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		queue.Write(writerSpeed)
	}()

	go func() {
		defer wg.Done()
		queue.Read(readerSpeed)
	}()

	wg.Wait()

	fmt.Println("\nПрограмма завершила работу")

	return nil
}
