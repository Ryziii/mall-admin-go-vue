package service

import (
	"encoding/json"
	"mall-admin-server-go/model"
)

type ProductSrv struct{}

func GetProducts(offset int, pageSize int) ([]model.Product, int64, error) {
	products, err := model.GetProductsByPagination(offset, pageSize)
	total := model.CountProduct()
	if err != nil {
		return nil, total, err
	}
	return products, total, nil
}

func GetProductById(productId int) (model.Product, error) {
	return model.GetProductById(productId)
}

func UpdateProduct(productId int, data interface{}) error {
	marshalData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var product model.Product
	err = json.Unmarshal(marshalData, &product)
	if err != nil {
		return err
	}
	product.ProductId = productId
	return model.UpdateProduct(productId, product)
}
