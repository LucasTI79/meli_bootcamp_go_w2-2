package services

import (
	"context"
	"encoding/json"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/purchase_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories/mocks"
	buyer_mock "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func Test_purchaseOrderService_GetAll(t *testing.T) {
	type args struct {
		ctx                  *context.Context
		expectedGetAllResult []entities.PurchaseOrder
		expectedGetAllError  error
		expectedGetAllCalls  int
	}

	ctx := context.TODO()

	var expectedPurchaseOrders []entities.PurchaseOrder
	expectedPurchaseOrdersSerialized, _ := os.ReadFile("../../../test/resources/valid_purchase_orders.json")
	if err := json.Unmarshal(expectedPurchaseOrdersSerialized, &expectedPurchaseOrders); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    []entities.PurchaseOrder
		wantErr error
	}{
		{
			name: "Successfully get all",
			args: args{
				ctx:                  &ctx,
				expectedGetAllResult: expectedPurchaseOrders,
				expectedGetAllError:  nil,
				expectedGetAllCalls:  1,
			},
			want:    expectedPurchaseOrders,
			wantErr: nil,
		},
		{
			name: "Error getting all",
			args: args{
				ctx:                  &ctx,
				expectedGetAllResult: []entities.PurchaseOrder{},
				expectedGetAllError:  assert.AnError,
				expectedGetAllCalls:  1,
			},
			want:    []entities.PurchaseOrder{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sellerRepositoryMock := mocks.NewMockPurchaseOrderRepository(t)
			sellerRepositoryMock.On("GetAll", *tt.args.ctx).Return(tt.args.expectedGetAllResult, tt.args.expectedGetAllError)

			buyerRepositoryMock := buyer_mock.NewBuyerRepositoryMock()

			service := NewPurchaseOrderService(sellerRepositoryMock, buyerRepositoryMock)
			got, err := service.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			sellerRepositoryMock.AssertNumberOfCalls(t, "GetAll", tt.args.expectedGetAllCalls)
		})
	}
}

func Test_purchaseOrder_Get(t *testing.T) {
	type args struct {
		ctx               *context.Context
		id                int
		expectedGetResult entities.PurchaseOrder
		expectedGetError  error
		expectedGetCalls  int
	}

	ctx := context.TODO()

	var expectedPurchaseOrder entities.PurchaseOrder
	expectedPurchaseOrderSerialized, _ := os.ReadFile("../../../test/resources/valid_purchase_order.json")
	if err := json.Unmarshal(expectedPurchaseOrderSerialized, &expectedPurchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    entities.PurchaseOrder
		wantErr error
	}{
		{
			name: "Successfully get purchaseOrder from db",
			args: args{
				ctx:               &ctx,
				id:                expectedPurchaseOrder.ID,
				expectedGetResult: expectedPurchaseOrder,
				expectedGetError:  nil,
				expectedGetCalls:  1,
			},
			want:    expectedPurchaseOrder,
			wantErr: nil,
		},
		{
			name: "PurchaseOrder not found in db",
			args: args{
				ctx:               &ctx,
				id:                expectedPurchaseOrder.ID,
				expectedGetResult: entities.PurchaseOrder{},
				expectedGetError:  ErrNotFound,
				expectedGetCalls:  1,
			},
			want:    entities.PurchaseOrder{},
			wantErr: ErrNotFound,
		},
		{
			name: "Error connecting db",
			args: args{
				ctx:               &ctx,
				id:                expectedPurchaseOrder.ID,
				expectedGetResult: entities.PurchaseOrder{},
				expectedGetError:  assert.AnError,
				expectedGetCalls:  1,
			},
			want:    entities.PurchaseOrder{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			purchaseOrderRepositoryMock := mocks.NewMockPurchaseOrderRepository(t)
			purchaseOrderRepositoryMock.On("Get", *tt.args.ctx, mock.AnythingOfType("int")).Return(tt.args.expectedGetResult, tt.args.expectedGetError)

			buyerRepositoryMock := buyer_mock.NewBuyerRepositoryMock()
			service := NewPurchaseOrderService(purchaseOrderRepositoryMock, buyerRepositoryMock)
			got, err := service.Get(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			purchaseOrderRepositoryMock.AssertNumberOfCalls(t, "Get", tt.args.expectedGetCalls)

		})
	}
}

func Test_purchaseOrder_Create(t *testing.T) {
	type args struct {
		ctx                  *context.Context
		purchaseOrder        entities.PurchaseOrder
		expectedExistsResult bool
		expectedExistsCalls  int
		expectedSaveResult   int
		expectedSaveError    error
		expectedSaveCalls    int
	}

	ctx := context.TODO()

	var expectedPurchaseOrder entities.PurchaseOrder
	expectedPurchaseOrderSerialized, _ := os.ReadFile("../../../test/resources/valid_purchase_order.json")
	if err := json.Unmarshal(expectedPurchaseOrderSerialized, &expectedPurchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    entities.PurchaseOrder
		wantErr error
	}{
		{
			name: "Successfully create purchaseOrder",
			args: args{
				ctx:                  &ctx,
				purchaseOrder:        expectedPurchaseOrder,
				expectedExistsResult: false,
				expectedExistsCalls:  1,
				expectedSaveResult:   expectedPurchaseOrder.ID,
				expectedSaveError:    nil,
				expectedSaveCalls:    1,
			},
			want:    expectedPurchaseOrder,
			wantErr: nil,
		},
		{
			name: "Error duplicated purchaseOrder",
			args: args{
				ctx:                  &ctx,
				purchaseOrder:        expectedPurchaseOrder,
				expectedExistsResult: true,
				expectedExistsCalls:  1,
				expectedSaveResult:   0,
				expectedSaveError:    nil,
				expectedSaveCalls:    0,
			},
			want:    entities.PurchaseOrder{},
			wantErr: ErrConflict,
		},
		{
			name: "Error saving purchaseOrder",
			args: args{
				ctx:                  &ctx,
				purchaseOrder:        expectedPurchaseOrder,
				expectedExistsResult: false,
				expectedExistsCalls:  1,
				expectedSaveError:    assert.AnError,
				expectedSaveCalls:    1,
			},
			want:    entities.PurchaseOrder{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			purchaseOrderRepositoryMock := mocks.NewMockPurchaseOrderRepository(t)

			purchaseOrderRepositoryMock.On(
				"Exists",
				*tt.args.ctx,
				mock.AnythingOfType("int"),
			).Return(
				tt.args.expectedExistsResult,
			)

			purchaseOrderRepositoryMock.On(
				"Save",
				*tt.args.ctx,
				mock.AnythingOfType("entities.PurchaseOrder"),
			).Return(
				tt.args.expectedSaveResult,
				tt.args.expectedSaveError,
			)

			buyerRepositoryMock := buyer_mock.NewBuyerRepositoryMock()
			service := NewPurchaseOrderService(purchaseOrderRepositoryMock, buyerRepositoryMock)
			got, err := service.Create(tt.args.ctx, tt.args.purchaseOrder)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			purchaseOrderRepositoryMock.AssertNumberOfCalls(t, "Exists", tt.args.expectedExistsCalls)
			purchaseOrderRepositoryMock.AssertNumberOfCalls(t, "Save", tt.args.expectedSaveCalls)
		})
	}
}

func Test_purchaseOrder_Update(t *testing.T) {
	type args struct {
		ctx                        *context.Context
		id                         int
		updatePurchaseOrderRequest dtos.UpdatePurchaseOrderRequestDTO
		expectedGetResult          entities.PurchaseOrder
		expectedGetError           error
		expectedGetCalls           int
		expectedExistsResult       bool
		expectedExistsCalls        int
		expectedUpdateError        error
		expectedUpdateCalls        int
	}

	ctx := context.TODO()

	var originalPurchaseOrder entities.PurchaseOrder
	originalPurchaseOrderSerialized, _ := os.ReadFile("../../../test/resources/valid_purchase_order.json")
	err := json.Unmarshal(originalPurchaseOrderSerialized, &originalPurchaseOrder)
	if err != nil {
		t.Fatal(err)
	}

	newOrderNumber := "2"
	newOrderDate := "2021-04-05"
	newTrackingCode := "abcdf124"
	newBuyerID := 2
	newCarrierID := 2
	newOrderStatusID := 2
	newWarehouseID := 2
	newProductRecordID := 2

	updatePurchaseOrderRequest := dtos.UpdatePurchaseOrderRequestDTO{
		OrderNumber:     &newOrderNumber,
		OrderDate:       &newOrderDate,
		TrackingCode:    &newTrackingCode,
		BuyerID:         &newBuyerID,
		CarrierID:       &newCarrierID,
		OrderStatusID:   &newOrderStatusID,
		WarehouseID:     &newWarehouseID,
		ProductRecordID: &newProductRecordID,
	}

	updatedPurchaseOrder := entities.PurchaseOrder{
		ID:              originalPurchaseOrder.ID,
		OrderNumber:     newOrderNumber,
		OrderDate:       newOrderDate,
		TrackingCode:    newTrackingCode,
		BuyerID:         newBuyerID,
		CarrierID:       newCarrierID,
		OrderStatusID:   newOrderStatusID,
		WarehouseID:     newWarehouseID,
		ProductRecordID: newProductRecordID,
	}

	tests := []struct {
		name    string
		args    args
		want    entities.PurchaseOrder
		wantErr error
	}{
		{
			name: "Successfully updating all fields",
			args: args{
				ctx:                        &ctx,
				id:                         originalPurchaseOrder.ID,
				updatePurchaseOrderRequest: updatePurchaseOrderRequest,
				expectedGetResult:          originalPurchaseOrder,
				expectedGetError:           nil,
				expectedGetCalls:           1,
				expectedExistsResult:       false,
				expectedExistsCalls:        1,
				expectedUpdateError:        nil,
				expectedUpdateCalls:        1,
			},
			want:    updatedPurchaseOrder,
			wantErr: nil,
		},
		{
			name: "Error purchaseOrder doesn't exists",
			args: args{
				ctx:                        &ctx,
				id:                         originalPurchaseOrder.ID,
				updatePurchaseOrderRequest: updatePurchaseOrderRequest,
				expectedGetResult:          entities.PurchaseOrder{},
				expectedGetError:           ErrNotFound,
				expectedGetCalls:           1,
			},
			want:    entities.PurchaseOrder{},
			wantErr: ErrNotFound,
		},
		{
			name: "Error duplicated purchaseOrder_id",
			args: args{
				ctx:                        &ctx,
				id:                         originalPurchaseOrder.ID,
				updatePurchaseOrderRequest: updatePurchaseOrderRequest,
				expectedGetResult:          originalPurchaseOrder,
				expectedGetError:           nil,
				expectedGetCalls:           1,
				expectedExistsResult:       true,
				expectedExistsCalls:        1,
			},
			want:    entities.PurchaseOrder{},
			wantErr: ErrConflict,
		},
		{
			name: "Error updating purchaseOrder",
			args: args{
				ctx:                        &ctx,
				id:                         originalPurchaseOrder.ID,
				updatePurchaseOrderRequest: updatePurchaseOrderRequest,
				expectedGetResult:          originalPurchaseOrder,
				expectedGetError:           nil,
				expectedGetCalls:           1,
				expectedExistsResult:       false,
				expectedExistsCalls:        1,
				expectedUpdateError:        assert.AnError,
				expectedUpdateCalls:        1,
			},
			want:    entities.PurchaseOrder{},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			purchaseOrderRepositoryMock := mocks.NewMockPurchaseOrderRepository(t)
			buyerRepositoryMock := buyer_mock.NewBuyerRepositoryMock()

			service := NewPurchaseOrderService(purchaseOrderRepositoryMock, buyerRepositoryMock)

			purchaseOrderRepositoryMock.On("Get", *tt.args.ctx, mock.AnythingOfType("int")).Return(tt.args.expectedGetResult, tt.args.expectedGetError)
			purchaseOrderRepositoryMock.On("Exists", *tt.args.ctx, mock.AnythingOfType("int")).Return(tt.args.expectedExistsResult)
			purchaseOrderRepositoryMock.On("Update", *tt.args.ctx, mock.AnythingOfType("entities.PurchaseOrder")).Return(tt.args.expectedUpdateError)

			newPurchaseOrder, err := service.Update(tt.args.ctx, tt.args.id, tt.args.updatePurchaseOrderRequest)

			assert.Equal(t, tt.want, newPurchaseOrder)
			assert.Equal(t, tt.wantErr, err)

			purchaseOrderRepositoryMock.AssertNumberOfCalls(t, "Get", tt.args.expectedGetCalls)
			purchaseOrderRepositoryMock.AssertNumberOfCalls(t, "Exists", tt.args.expectedExistsCalls)
			purchaseOrderRepositoryMock.AssertNumberOfCalls(t, "Update", tt.args.expectedUpdateCalls)

		})
	}
}

func Test_purchaseOrder_Delete(t *testing.T) {
	type args struct {
		ctx                 *context.Context
		id                  int
		expectedGetResult   entities.PurchaseOrder
		expectedGetError    error
		expectedGetCalls    int
		expectedDeleteError error
		expectedDeleteCalls int
	}

	ctx := context.TODO()

	var expectedPurchaseOrder entities.PurchaseOrder
	expectedPurchaseOrderSerialized, _ := os.ReadFile("../../../test/resources/valid_purchase_order.json")
	if err := json.Unmarshal(expectedPurchaseOrderSerialized, &expectedPurchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Successfully deleting purchaseOrder",
			args: args{
				ctx:                 &ctx,
				id:                  1,
				expectedGetResult:   expectedPurchaseOrder,
				expectedGetError:    nil,
				expectedGetCalls:    1,
				expectedDeleteError: nil,
				expectedDeleteCalls: 1,
			},
			wantErr: nil,
		},
		{
			name: "Error getting purchaseOrder",
			args: args{
				ctx:                 &ctx,
				id:                  1,
				expectedGetResult:   entities.PurchaseOrder{},
				expectedGetError:    assert.AnError,
				expectedGetCalls:    1,
				expectedDeleteError: nil,
				expectedDeleteCalls: 0,
			},
			wantErr: assert.AnError,
		},
		{
			name: "Error deleting purchaseOrder",
			args: args{
				ctx:                 &ctx,
				id:                  1,
				expectedGetError:    nil,
				expectedGetCalls:    1,
				expectedDeleteError: assert.AnError,
				expectedDeleteCalls: 0,
			},
			wantErr: assert.AnError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			purchaseOrderRepositoryMock := mocks.NewMockPurchaseOrderRepository(t)
			purchaseOrderRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(tt.args.expectedGetResult, tt.args.expectedGetError)
			purchaseOrderRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(tt.args.expectedDeleteError)

			buyerRepositoryMock := buyer_mock.NewBuyerRepositoryMock()
			service := NewPurchaseOrderService(purchaseOrderRepositoryMock, buyerRepositoryMock)

			err := service.Delete(&ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err)
			purchaseOrderRepositoryMock.On("Get", tt.args.expectedGetCalls)
			purchaseOrderRepositoryMock.On("Delete", tt.args.expectedDeleteCalls)
		})
	}
}

func Test_purchaseOrderService_CountByPurchaseOrderID(t *testing.T) {
	type args struct {
		ctx                          *context.Context
		id                           int
		expectedGetResult            domain.Buyer
		expectedGetError             error
		expectedGetCalls             int
		expectedCountByBuyerIDResult int
		expectedCountByBuyerIDError  error
		expectedCountByBuyerIDCalls  int
	}

	ctx := context.TODO()

	var expectedBuyer domain.Buyer
	expectedBuyerSerialized, _ := os.ReadFile("../../../test/resources/valid_buyer.json")
	if err := json.Unmarshal(expectedBuyerSerialized, &expectedBuyer); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Successfully counting puchaseOrders by buyer id",
			args: args{
				ctx:                          &ctx,
				id:                           1,
				expectedGetResult:            expectedBuyer,
				expectedGetError:             nil,
				expectedGetCalls:             1,
				expectedCountByBuyerIDResult: 1,
				expectedCountByBuyerIDError:  nil,
				expectedCountByBuyerIDCalls:  1,
			},
			want:    1,
			wantErr: nil,
		},
		{
			name: "Error getting purchaseOrders by buyer id",
			args: args{
				ctx:                          &ctx,
				id:                           1,
				expectedGetResult:            expectedBuyer,
				expectedGetError:             nil,
				expectedGetCalls:             1,
				expectedCountByBuyerIDResult: 0,
				expectedCountByBuyerIDError:  assert.AnError,
				expectedCountByBuyerIDCalls:  1,
			},
			want:    0,
			wantErr: assert.AnError,
		},
		{
			name: "Error getting buyer",
			args: args{
				ctx:                          &ctx,
				id:                           1,
				expectedGetResult:            domain.Buyer{},
				expectedGetError:             assert.AnError,
				expectedGetCalls:             1,
				expectedCountByBuyerIDResult: 0,
				expectedCountByBuyerIDError:  nil,
				expectedCountByBuyerIDCalls:  0,
			},
			want:    0,
			wantErr: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buyerRepositoryMock := buyer_mock.NewBuyerRepositoryMock()
			buyerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(tt.args.expectedGetResult, tt.args.expectedGetError)

			purchaseOrderRepositoryMock := mocks.NewMockPurchaseOrderRepository(t)
			purchaseOrderRepositoryMock.On("CountByBuyerID", ctx, mock.AnythingOfType("int")).Return(tt.args.expectedCountByBuyerIDResult, tt.args.expectedCountByBuyerIDError)

			service := NewPurchaseOrderService(purchaseOrderRepositoryMock, buyerRepositoryMock)

			got, err := service.CountByBuyerID(&ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "Get", tt.args.expectedGetCalls)
			purchaseOrderRepositoryMock.AssertNumberOfCalls(t, "CountByBuyerID", tt.args.expectedCountByBuyerIDCalls)

		})
	}
}
