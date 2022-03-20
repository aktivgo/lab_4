package task_3

import (
	"fmt"
	"lab_4/internal/types"
	"log"
	"math/rand"
	"sync"
	"time"
)

func Run() error {
	var storageSize int
	fmt.Print("Введите размер хранилища: ")
	_, _ = fmt.Scanf("%d\n", &storageSize)

	storage := types.NewStorage(storageSize)

	var countWorkers int
	fmt.Print("Введите количество потоков-рабочих: ")
	_, _ = fmt.Scanf("%d\n", &countWorkers)

	var countReaders int
	fmt.Print("Введите количество потоков-читателей: ")
	_, _ = fmt.Scanf("%d\n", &countReaders)

	var wgWorker sync.WaitGroup
	wgWorker.Add(countWorkers)

	for i := 0; i < countWorkers; i++ {
		workerNumber := i + 1
		go func() {
			defer wgWorker.Done()

			for {
				time.Sleep(time.Duration(rand.Intn(2000-1)) * time.Millisecond)

				index := rand.Intn(storageSize)
				value := 1 + rand.Intn(8)

				if err := storage.Inc(index, value); err != nil {
					log.Printf("Рабочий №%d: %s\n", workerNumber, err)
				}

				log.Printf("Рабочий №%d: индекс %d, значение %d\n", workerNumber, index, value)
			}
		}()
	}

	var wgReader sync.WaitGroup
	wgReader.Add(countReaders)

	for i := 0; i < countReaders; i++ {
		readerNumber := i + 1
		go func() {
			defer wgReader.Done()

			for {
				time.Sleep(time.Duration(rand.Intn(2000-1)) * time.Millisecond)

				index := rand.Intn(storageSize)

				var value int
				value, err := storage.Get(index)
				if err != nil {
					log.Printf("Читатель №%d: %s\n", readerNumber, err)
				}

				log.Printf("Читатель №%d: индекс = %d, значение = %d\n", readerNumber, index, value)
			}
		}()
	}

	wgWorker.Wait()
	wgReader.Wait()

	fmt.Println("\nПрограмма завершила работу")

	return nil
}
