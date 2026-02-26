package main

type status int

const (
	todo status = iota
	inProgress
	done
)

type Task struct {
	Title       string
	Description string
	Status      status
}

func (t Task) GetTitle() string {
	return t.Title
}

func (t Task) GetDescription() string {
	return t.Description
}

func (t Task) GetStatus() status {
	return t.Status
}
