package productbatches_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/productbatchesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

var (
	expectedProductBatch = domain.ProductBatches{
		ID:                 1,
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}
	payload = productbatchesdto.CreateProductBatchesDTO{
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}
)

func TestCreate(t *testing.T) {
	// t.Run("CREATE - create_conflict", func(t *testing.T) {
	// 	// Criar um contexto
	// 	ctx := context.Background()

	// 	// Criar um objeto mock da Service
	// 	mockService := new(mocks.ProductBatchServiceMock)

	// 	// Definir o comportamento esperado do mockService
	// 	mockService.On("Save", &ctx, payload).Return(&domain.ProductBatches{}, productbatches.ErrConflict)

	// 	// Chamar a função Save da mockService
	// 	result, err := mockService.Save(&ctx, payload)

	// 	// Verificar se ocorreu um erro
	// 	if err == nil {
	// 		t.Error("Erro esperado, mas nenhum erro ocorreu")
	// 	} else if err != ErrConflict {
	// 		t.Errorf("Erro incorreto. Esperado: %s, Obtido: %s", ErrConflict.Error(), err.Error())
	// 	}

	// 	// Verificar se o resultado é nulo
	// 	if result != nil {
	// 		t.Error("O resultado deve ser nulo")
	// 	}

	// 	// Verificar se o método mockado foi chamado corretamente
	// 	mockService.AssertExpectations(t)

	// })
}
