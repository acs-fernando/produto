package repository

import (
	"fmt"

	"github.com/acs-fernando/domain/model"
	"github.com/jinzhu/gorm"
)

type ProdutoRepositoryDb struct {
	Db *gorm.DB
}

func (t ProdutoRepositoryDb) CreateProduct(produto *model.Produtos) error {
	err := t.Db.Create(produto).Error
	if err != nil {
		return err
	}
	return nil
}

func (t ProdutoRepositoryDb) FindProducts() ([]model.Produtos, error) {
	var produtos []model.Produtos
	t.Db.Find(&produtos)

	if produtos == nil {
		return nil, fmt.Errorf("no key was found")
	}
	return produtos, nil
}
