package task_2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Run() error {
	var inputFilePath string
	/*fmt.Print("Введите путь до входного файла: ")
	_, _ = fmt.Scanf("%d\n", &inputFilePath)*/
	inputFilePath = "input/arrays.txt"

	var outputFilePath string
	/*fmt.Print("Введите путь до выходного файла: ")
	_, _ = fmt.Scanf("%d\n", &outputFilePath)*/
	outputFilePath = "output/result.txt"

	var readerSpeed int
	fmt.Print("Введите скорость работы потока-считывателя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &readerSpeed)

	var workerSpeed int
	fmt.Print("Введите скорость работы потока-исполнителя(в миллисекундах): ")
	_, _ = fmt.Scanf("%d\n", &workerSpeed)
	fmt.Println()

	inputChannel := make(chan []int, 64)
	outputChannel := make(chan int, 64)

	var wg sync.WaitGroup
	wg.Add(2)

	// Горутина, считывающая массивы из файла и записывающая результат в новый файл
	go func() {
		defer wg.Done()

		inputFile, err := os.Open(inputFilePath)
		if err != nil {
			log.Fatalln(err)
		}
		defer inputFile.Close()

		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			log.Fatalln(err)
		}
		defer outputFile.Close()

		scanner := bufio.NewScanner(inputFile)
		scanner.Scan()
		count, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Считано количество массивов m = %d\n", count)

		for i := 0; i < count; i++ {
			time.Sleep(time.Duration(readerSpeed) * time.Millisecond)

			if i < count {
				scanner.Scan()

				if err = scanner.Err(); err != nil {
					log.Fatalln(err)
				}

				line := scanner.Text()

				fmt.Printf("Считан массив №%d: %s\n", i+1, line)

				var curArray []int

				nums := strings.Split(line, " ")
				for _, num := range nums {
					n, err := strconv.Atoi(num)
					if err != nil {
						log.Fatalln(err)
					}
					curArray = append(curArray, n)
				}

				inputChannel <- curArray
			}

		}

		close(inputChannel)

		i := 0
		for {
			num, ok := <-outputChannel
			if !ok {
				return
			}

			cond := ""
			if i != count-1 {
				cond = "\n"
			}

			sum := strconv.Itoa(num)
			_, err = outputFile.WriteString(sum + cond)
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Printf("Записана сумма №%d: %s\n", i+1, sum)
			i++
		}
	}()

	// Горутина, суммирующая элементы массива
	go func() {
		defer wg.Done()

		i := 0
		for {
			time.Sleep(time.Duration(workerSpeed) * time.Millisecond)

			curArray, ok := <-inputChannel
			if !ok {
				break
			}

			if curArray == nil {
				continue
			}

			sum := 0
			for _, num := range curArray {
				sum += num
			}

			outputChannel <- sum

			fmt.Printf("Получена сумма №%d: %d\n", i+1, sum)
			i++
		}

		close(outputChannel)
	}()

	wg.Wait()

	fmt.Println("\nПрограмма завершила работу")

	return nil
}
