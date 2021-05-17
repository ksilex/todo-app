package todo

type TodoList struct {
	Id          int
	Title       string
	Description string
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
