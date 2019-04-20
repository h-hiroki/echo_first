package models

type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type TodoCollection struct {
	Todos []Todo `json:"items"`
}

func GetTodo() (tc TodoCollection) {
	tc = TodoCollection {
		[]Todo {
			{1, "todo"},
			{2, "echo"},
		},
	}
	return tc
}