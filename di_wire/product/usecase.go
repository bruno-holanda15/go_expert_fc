package product

type ProductUseCase struct {
	repository ProductRepositoryInterface
}

func NewProductUseCase(repository ProductRepositoryInterface) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

// deveria retornar um DTO ao invés da entidade
func (u *ProductUseCase) GetProductById(id int) (Product, error) {
	return u.repository.GetProductById(id)
}