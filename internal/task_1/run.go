package task_1

import (
	"fmt"
	"lab_4/internal/types"
	"sync"
	"time"
)

func Run() error {
	var countApplication int
	fmt.Print("Введите количество заявок: ")
	_, _ = fmt.Scanf("%d\n", &countApplication)

	var queueSize int
	fmt.Print("Введите размер очереди: ")
	_, _ = fmt.Scanf("%d\n", &queueSize)

	var producerSpeed int
	fmt.Print("Введите скорость работы производителя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &producerSpeed)

	var consumerSpeed int
	fmt.Print("Введите скорость работы потребителя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &consumerSpeed)
	fmt.Println()

	applicationChannel := make(chan *types.Application, queueSize)

	var wg sync.WaitGroup
	wg.Add(2)

	var unprocessedApplications []*types.Application

	// Горутина - производитель
	go func() {
		defer wg.Done()

		for i := 0; i < countApplication; i++ {
			time.Sleep(time.Duration(producerSpeed) * time.Millisecond)

			application := types.NewApplication(i + 1)

			if len(applicationChannel) == cap(applicationChannel) {
				fmt.Printf("Производитель: очередь заполнена, заявка №%d отклонена\n", i+1)
				unprocessedApplications = append(unprocessedApplications, application)
				continue
			}

			applicationChannel <- application

			fmt.Printf("Производитель: заявка №%d создана, в очереди %d\n", i+1, len(applicationChannel))
		}

		close(applicationChannel)
	}()

	// Горутина - потребитель
	go func() {
		defer wg.Done()

		for {
			time.Sleep(time.Duration(consumerSpeed) * time.Millisecond)

			application, ok := <-applicationChannel
			if !ok {
				return
			}

			if application == nil {
				fmt.Printf("Потребитель: очередь пустая\n")
				continue
			}

			fmt.Printf("Потребитель: заявка №%d обработана, в очереди %d\n", application.ID, len(applicationChannel))
		}
	}()

	wg.Wait()

	fmt.Println("\nПрограмма завершила работу")

	fmt.Printf("Необработанных заявок %d: ", len(unprocessedApplications))
	for _, application := range unprocessedApplications {
		fmt.Print(application.ID, " ")
	}

	return nil
}
