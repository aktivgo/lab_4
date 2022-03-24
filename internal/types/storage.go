package types

import (
	"fmt"
	"sync"
)

type Storage struct {
	rw sync.RWMutex
	m sync.Mutex
	Data []int
}

func NewStorage(size int) *Storage {
	data := initData(size)

	return &Storage{
		Data: data,
	}
}

func initData(size int) []int {
	var data []int

	for i := 0; i < size; i++ {
		data = append(data, 0)
	}

	return data
}

func (s *Storage) Inc(index int, value int) error {
	s.m.Lock()
	defer s.m.Unlock()

	if err := s.validateIndex(index); err != nil {
		return err
	}

	s.Data[index] = value

	return nil
}

func (s *Storage) Get(index int) (int, error) {
	s.rw.Lock()
	defer s.rw.Unlock()

	if err := s.validateIndex(index); err != nil {
		return 0, err
	}

	return s.Data[index], nil
}

func (s *Storage) validateIndex(index int) error {
	if index < 0 || index >= len(s.Data) {
		return fmt.Errorf("%d: выход за пределы массива (%d; %d)", index, 0, len(s.Data))
	}

	return nil
}
