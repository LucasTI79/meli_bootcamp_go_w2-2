package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"os"
	"regexp"
	"testing"
)

func Test_purchaseOrderRepository_GetAll(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	validPurchaseOrdersSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_orders.json")
	var validPurchaseOrders []entities.PurchaseOrder
	if err := json.Unmarshal(validPurchaseOrdersSerialized, &validPurchaseOrders); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.PurchaseOrder
		wantErr bool
	}{
		{
			name: "Successfully get purchaseOrders",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want:    validPurchaseOrders,
			wantErr: false,
		},
		{
			name: "Error getting purchaseOrders",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
			},
			want:    []entities.PurchaseOrder{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewPurchaseOrderRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id", "order_number", "order_date", "tracking_code", "buyer_id", "carrier_id", "order_status_id", "warehouse_id", "product_record_id"})
			for _, purchaseOrder := range validPurchaseOrders {
				rows.AddRow(purchaseOrder.ID, purchaseOrder.OrderNumber, purchaseOrder.OrderDate, purchaseOrder.TrackingCode, purchaseOrder.BuyerID, purchaseOrder.CarrierID, purchaseOrder.OrderStatusID, purchaseOrder.WarehouseID, purchaseOrder.ProductRecordID)
			}
			if tt.wantErr {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(GetAllPurchaseOrders).
				WithArgs().
				WillReturnRows(rows)

			got, err := r.GetAll(tt.args.ctx)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_purchaseOrderRepository_Get(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	purchaseOrderSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_order.json")
	var validPurchaseOrder entities.PurchaseOrder
	if err := json.Unmarshal(purchaseOrderSerialized, &validPurchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.PurchaseOrder
		wantErr bool
	}{
		{
			name:   "Successfully get purchaseOrder",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  validPurchaseOrder.ID,
			},
			want:    validPurchaseOrder,
			wantErr: false,
		},
		{
			name:   "Error nonexistent purchaseOrder id",
			fields: fields{db},
			args: args{
				ctx: ctx,
				id:  999,
			},
			want:    entities.PurchaseOrder{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewPurchaseOrderRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id", "order_number", "order_date", "tracking_code", "buyer_id", "carrier_id", "order_status_id", "warehouse_id", "product_record_id"})
			rows.AddRow(tt.want.ID, tt.want.OrderNumber, tt.want.OrderDate, tt.want.TrackingCode, tt.want.BuyerID, tt.want.CarrierID, tt.want.OrderStatusID, tt.want.WarehouseID, tt.want.ProductRecordID)

			mock.ExpectQuery(GetPurchaseOrderByID).
				WithArgs(tt.want.ID).
				WillReturnRows(rows)

			got, err := r.Get(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_purchaseOrderRepository_Exists(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	purchaseOrderSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_order.json")
	var purchaseOrder entities.PurchaseOrder
	if err := json.Unmarshal(purchaseOrderSerialized, &purchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Exists",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  purchaseOrder.ID,
			},
			want: true,
		},
		{
			name: "Does not exists",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  purchaseOrder.ID,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewPurchaseOrderRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(purchaseOrder.ID)

			if !tt.want {
				rows.RowError(0, sql.ErrNoRows)
			}

			mock.ExpectQuery(ExistsPurchaseOrderByID).
				WithArgs(tt.args.id).
				WillReturnRows(rows)

			got := r.Exists(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_purchaseOrderRepository_Save(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx           context.Context
		purchaseOrder entities.PurchaseOrder
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	purchaseOrderSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_order.json")
	var purchaseOrder entities.PurchaseOrder
	if err := json.Unmarshal(purchaseOrderSerialized, &purchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Successfully save purchaseOrder",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:           ctx,
				purchaseOrder: purchaseOrder,
			},
			want:    purchaseOrder.ID,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewPurchaseOrderRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(SavePurchaseOrder))
			mock.ExpectExec(regexp.QuoteMeta(SavePurchaseOrder)).
				WithArgs(
					tt.args.purchaseOrder.OrderNumber,
					tt.args.purchaseOrder.OrderDate,
					tt.args.purchaseOrder.TrackingCode,
					tt.args.purchaseOrder.BuyerID,
					tt.args.purchaseOrder.CarrierID,
					tt.args.purchaseOrder.OrderStatusID,
					tt.args.purchaseOrder.WarehouseID,
					tt.args.purchaseOrder.ProductRecordID,
				).
				WillReturnResult(sqlmock.NewResult(int64(tt.want), 1))

			got, err := r.Save(tt.args.ctx, tt.args.purchaseOrder)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_purchaseOrderRepository_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx           context.Context
		purchaseOrder entities.PurchaseOrder
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	purchaseOrderSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_order.json")
	var purchaseOrder entities.PurchaseOrder
	if err := json.Unmarshal(purchaseOrderSerialized, &purchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully update purchaseOrder",
			fields: fields{
				db: db,
			},
			args: args{
				ctx:           ctx,
				purchaseOrder: purchaseOrder,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewPurchaseOrderRepository(tt.fields.db)

			mock.ExpectPrepare(regexp.QuoteMeta(UpdatePurchaseOrder))
			mock.ExpectExec(regexp.QuoteMeta(UpdatePurchaseOrder)).
				WithArgs(
					tt.args.purchaseOrder.OrderNumber,
					tt.args.purchaseOrder.OrderDate,
					tt.args.purchaseOrder.TrackingCode,
					tt.args.purchaseOrder.BuyerID,
					tt.args.purchaseOrder.CarrierID,
					tt.args.purchaseOrder.OrderStatusID,
					tt.args.purchaseOrder.WarehouseID,
					tt.args.purchaseOrder.ProductRecordID,
					tt.args.purchaseOrder.ID,
				).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := r.Update(tt.args.ctx, tt.args.purchaseOrder)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_purchaseOrderRepository_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New()
	ctx := context.TODO()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successfully delete purchaseOrder",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  6700,
			},
			wantErr: false,
		},
		{
			name: "No purchaseOrder to delete",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  6700,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewPurchaseOrderRepository(tt.fields.db)

			rowsAffected := int64(1)
			if tt.wantErr {
				rowsAffected = 0
			}

			mock.ExpectPrepare(regexp.QuoteMeta(DeletePurchaseOrderByID))
			mock.ExpectExec(regexp.QuoteMeta(DeletePurchaseOrderByID)).
				WithArgs(
					tt.args.id,
				).
				WillReturnResult(sqlmock.NewResult(1, rowsAffected))

			err := r.Delete(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func Test_purchaseOrderRepository_CountByBuyerID(t *testing.T) {
	type fields struct {
		db *sql.DB
	}

	type args struct {
		ctx context.Context
		id  int
	}

	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	ctx := context.TODO()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Successfully count purchaseOrder",
			fields: fields{
				db: db,
			},
			args: args{
				ctx: ctx,
				id:  6700,
			},
			want:    1,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewPurchaseOrderRepository(tt.fields.db)

			rows := sqlmock.NewRows([]string{"count(*)"}).AddRow(tt.want)

			mock.ExpectQuery(CountByBuyerID).
				WithArgs(tt.args.id).
				WillReturnRows(rows)

			got, err := r.CountByBuyerID(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
