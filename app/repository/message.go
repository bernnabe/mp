package repository

type MessageRepositoryInterface interface {
	Add(clave string, valor []string)
	Get(clave string) (valor []string)
	GetAll() (data map[string][]string)
	Clear()
}

type MessageRepository struct {
}

var (
	messagesParts = make(map[string][]string)
)

func NewMessageRepository() MessageRepositoryInterface {
	return &MessageRepository{}
}

func (repository *MessageRepository) Add(clave string, valor []string) {
	messagesParts[clave] = valor
}

func (repository *MessageRepository) Get(clave string) (valor []string) {
	return messagesParts[clave]
}

func (repository *MessageRepository) GetAll() (data map[string][]string) {
	return messagesParts
}

func (repository *MessageRepository) Clear() {
	for k := range messagesParts {
		delete(messagesParts, k)
	}
}
