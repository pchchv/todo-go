package main

type TodoRepository interface {
	Create(todo *Todo)
	GetAll() []*Todo
	Get(id int) (t *Todo, err error)
	Update(*Todo) (err error)
	DeleteAll()
	Delete(id int) (err error)
}

type InMemoryTodoRepository struct {
	Todos  []*Todo
	nextId int
}

func NewInMemoryTodoRepository() TodoRepository {
	t := new(InMemoryTodoRepository)
	t.Todos = make([]*Todo, 0)
	t.nextId = 1
	return t
}
