package Models

type Todo struct {
	ID          uint
	Title       string
	Description string
}

func (b *Todo) TableName() string {
	return "todo"
}
