package inputs

type CreateTodo struct {
	Title  string
	Detail string
}

func NewCreateTodo(
	title string,
	detail string,
) CreateTodo {
	return CreateTodo{
		title,
		detail,
	}
}

type GetTodo struct {
	Id uint64
}

func NewGetTodo(
	id uint64,
) GetTodo {
	return GetTodo{
		id,
	}
}
