package usecase

import (
	"fmt"

	"github.com/acs-fernando/domain/model"
)

type ProdutoUseCase struct {
	ProdutoKeyRepository model.ProdutoKeyRepositoryInterface
}

func (p *ProdutoUseCase) CreateProduct(name string, description string, price float32) (*model.Produtos, error) {
	fmt.Println("aqui")
	produtoKey, err := model.NewProdutoKey(name, description, price)
	if err != nil {
		return nil, err
	}

	p.ProdutoKeyRepository.CreateProduct(produtoKey)
	if produtoKey.ID == 0 {
		return nil, err
	}

	return produtoKey, nil
}

func (p *ProdutoUseCase) FindProducts() ([]model.Produtos, error) {
	produtoKey, err := p.ProdutoKeyRepository.FindProducts()
	if err != nil {
		return nil, err
	}
	return produtoKey, nil
}
