package internal

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	sync.Mutex
	Size int
	CountN int
	CurrentN int
}

func NewQueue(size int, countN int, currentN int) *Queue {
	return &Queue{
		Size: size,
		CountN: countN,
		CurrentN: currentN,
	}
}

func (c *Queue) Increment(speed int) {
	var i = 0
	for {
		c.Lock()

		if i >= c.CountN {
			return
		}

		if c.CurrentN == c.Size {
			fmt.Printf("Производитель: очередь заполнена, данные %d отклонены\n", i + 1)
			c.CountN -= 1
			c.Unlock()
			time.Sleep(time.Duration(speed) * time.Millisecond)
			i++
			continue
		}

		c.CurrentN++
		fmt.Printf("Производитель: данные %d, значение %d\n", i + 1, c.CurrentN)
		i++

		c.Unlock()

		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}

func (c *Queue) Decrement(speed int) {
	var i = 0
	for {
		c.Lock()

		if i >= c.CountN {
			return
		}

		if c.CurrentN == 0 {
			fmt.Printf("Потребитель: очередь пустая\n")
			c.Unlock()
			time.Sleep(time.Duration(speed) * time.Millisecond)
			i++
			continue
		}

		c.CurrentN--
		fmt.Printf("Потребитель: данные %d, значение %d\n", i + 1, c.CurrentN)
		i++

		c.Unlock()

		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}