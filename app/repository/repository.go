package repository

type Repository interface {
	Save(clave string, valor string)
}

type MessageRepository struct {
}

var (
	messagesRepository map[string]string
)

func NewRepository() Repository {
	return &MessageRepository{}
}

func (repository *MessageRepository) Save(clave string, valor string) {
	messagesRepository[clave] = valor
}
