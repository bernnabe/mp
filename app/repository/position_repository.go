package repository

type PositionRepositoryInterface interface {
	Add(clave string, valor float64)
	Get(clave string) (valor float64)
	GetAll() (data map[string]float64)
	Clear()
}

type PositionRepository struct {
}

var (
	DistancesParts = make(map[string]float64)
)

func NewPositionRepository() PositionRepositoryInterface {
	return &PositionRepository{}
}

func (repository *PositionRepository) Add(clave string, valor float64) {
	DistancesParts[clave] = valor
}

func (repository *PositionRepository) Get(clave string) (valor float64) {
	return DistancesParts[clave]
}

func (repository *PositionRepository) GetAll() (data map[string]float64) {
	return DistancesParts
}

func (repository *PositionRepository) Clear() {
	for k := range DistancesParts {
		delete(DistancesParts, k)
	}
}
