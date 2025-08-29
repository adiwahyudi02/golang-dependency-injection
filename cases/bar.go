package cases

type BarRepository struct {

}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService(repository *BarRepository) *BarService {
	return &BarService{repository}
}