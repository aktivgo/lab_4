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

	var workerSpeed int
	fmt.Print("Введите скорость работы потока-рабочего(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &workerSpeed)

	var countWorkerIterations int
	fmt.Print("Введите количество итераций потока-рабочего: ")
	_, _ = fmt.Scanf("%d\n", &countWorkerIterations)

	var countReaders int
	fmt.Print("Введите количество потоков-читателей: ")
	_, _ = fmt.Scanf("%d\n", &countReaders)

	var readerSpeed int
	fmt.Print("Введите скорость работы потока-читателя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &readerSpeed)

	var countReaderIterations int
	fmt.Print("Введите количество итераций потока-читателя: ")
	_, _ = fmt.Scanf("%d\n", &countReaderIterations)
	fmt.Println()

	var wgWorker sync.WaitGroup
	wgWorker.Add(countWorkers)

	for i := 0; i < countWorkers; i++ {
		workerNumber := i + 1
		go func() {
			defer wgWorker.Done()

			for j := 0; j < countWorkerIterations; j++ {
				time.Sleep(time.Duration(workerSpeed) * time.Millisecond)

				index := rand.Intn(storageSize - 1)
				value := 1 + rand.Intn(8)

				if err := storage.Inc(index, value); err != nil {
					log.Printf("Рабочий №%d: %s\n", workerNumber, err)
				}

				log.Printf("Рабочий №%d: итерация %d, индекс %d, значение инкремента %d\n", j+1, workerNumber, index, value)
			}
		}()
	}

	var wgReader sync.WaitGroup
	wgReader.Add(countReaders)

	for i := 0; i < countReaders; i++ {
		readerNumber := i + 1
		go func() {
			defer wgReader.Done()

			for j := 0; j < countReaderIterations; j++ {
				time.Sleep(time.Duration(readerSpeed) * time.Millisecond)

				index := rand.Intn(storageSize - 1)

				var value int
				value, err := storage.Get(index)
				if err != nil {
					log.Printf("Читатель №%d: %s\n", readerNumber, err)
				}

				log.Printf("Читатель №%d: итерация %d, индекс = %d, значение = %d\n", j+1, readerNumber, index, value)
			}
		}()
	}

	wgWorker.Wait()
	wgReader.Wait()

	fmt.Println("\nПрограмма завершила работу")

	return nil
}
