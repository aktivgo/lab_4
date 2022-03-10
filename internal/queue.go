package internal

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	sync.Mutex
	Size                    int
	CountApplication        int
	Counter                 int
	CountProcessApplication int
}

func NewQueue(queueSize int, countApplication int) *Queue {
	return &Queue{
		Size:                    queueSize,
		CountApplication:        countApplication,
		Counter:                 0,
		CountProcessApplication: countApplication,
	}
}

func (c *Queue) Write(speed int) {
	for i := 0; i < c.CountApplication; i++ {
		time.Sleep(time.Duration(speed) * time.Millisecond)

		c.Lock()

		if c.Counter >= c.Size {
			c.CountProcessApplication--
			fmt.Printf("Производитель: очередь заполнена, данные %d отклонены\n", i+1)

			c.Unlock()

			continue
		}

		c.Counter++
		fmt.Printf("Производитель: данные %d, в очереди %d\n", i+1, c.Counter)

		c.Unlock()
	}
}

func (c *Queue) Read(speed int) {
	for i := 0; i < c.CountProcessApplication; {
		time.Sleep(time.Duration(speed) * time.Millisecond)

		c.Lock()

		if c.Counter <= 0 {
			fmt.Printf("Потребитель: очередь пустая\n")

			c.Unlock()

			continue
		}

		c.Counter--
		fmt.Printf("Потребитель: данные %d, в очереди %d\n", i+1, c.Counter)
		i++

		c.Unlock()
	}
}
