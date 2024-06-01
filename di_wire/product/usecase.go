package product

type ProductUseCase struct {
	repository *ProductRepository
}

func NewProductUseCase(repository *ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

// deveria retornar um DTO ao invés da entidade
func (u *ProductUseCase) GetProductById(id int) (Product, error) {
	return u.repository.GetProductById(id)
}