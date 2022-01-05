package application_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/thg021/fc2-arquitetura-hexagonal/application"
)

func TestProduct_Enable(t *testing.T) {

	product := application.Product{}
	product.Name = "Teste"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {

	product := application.Product{}
	product.Name = "Teste"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()

	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())

}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Teste"
	product.Price = 10
	product.ID = uuid.NewV4().String()
	product.Status = application.ENABLED

	product.Status = "INVALID"
	_, err := product.IsValid()
	require.Equal(t, "the status must be enable or disable", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()

	id := product.GetID()
	require.Equal(t, id, product.ID)

}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Teste"

	name := product.GetName()
	require.Equal(t, name, product.Name)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	price := product.GetPrice()
	require.Equal(t, price, product.Price)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = "DISABLED"

	status := product.GetStatus()
	require.Equal(t, status, product.Status)

}
