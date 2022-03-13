package types

type Application struct {
	ID int
}

func NewApplication(id int) *Application {
	return &Application{
		ID: id,
	}
}
