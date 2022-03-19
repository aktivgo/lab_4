package types

type Storage struct {
	Data []int
}

func NewStorage(data []int) *Storage {
	return &Storage{
		Data: data,
	}
}
