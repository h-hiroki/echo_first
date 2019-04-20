package models

type Todo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
}

type TodoCollection struct {
	Todos []Todo `json:"items"`
}

func GetTodos() (tc TodoCollection) {
	tc = TodoCollection {
		[]Todo {
			{1, "todo", "done"},
			{2, "echo", ""},
			{3, "vue", ""},
			{4, "docker","done"},
			{5, "laravel", "done"},
		},
	}
	return tc
}