package repository

type DistanceRepositoryInterface interface {
	Add(clave string, valor float64)
	Get(clave string) (valor float64)
	GetAll() (data map[string]float64)
	Clear()
}

type DistanceRepository struct {
}

var (
	DistancesParts = make(map[string]float64)
)

func NewDistanceRepository() DistanceRepositoryInterface {
	return &DistanceRepository{}
}

func (repository *DistanceRepository) Add(clave string, valor float64) {
	DistancesParts[clave] = valor
}

func (repository *DistanceRepository) Get(clave string) (valor float64) {
	return DistancesParts[clave]
}

func (repository *DistanceRepository) GetAll() (data map[string]float64) {
	return DistancesParts
}

func (repository *DistanceRepository) Clear() {
	for k := range DistancesParts {
		delete(DistancesParts, k)
	}
}
