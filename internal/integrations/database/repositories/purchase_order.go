package repositories

import (
	"context"
	"database/sql"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

const (
	GetAllPurchaseOrders    = "SELECT  purchase_orders.id, purchase_orders.order_number, purchase_orders.order_date, purchase_orders.tracking_code, purchase_orders.buyer_id, purchase_orders.carrier_id, purchase_orders.order_status_id, purchase_orders.warehouse_id, purchase_orders.product_record_id FROM purchase_orders"
	GetPurchaseOrderByID    = "SELECT  purchase_orders.id, purchase_orders.order_number, purchase_orders.order_date, purchase_orders.tracking_code, purchase_orders.buyer_id, purchase_orders.carrier_id, purchase_orders.order_status_id, purchase_orders.warehouse_id, purchase_orders.product_record_id FROM purchase_orders WHERE id = ?"
	ExistsPurchaseOrderByID = "SELECT id FROM purchase_orders WHERE id=?"
	SavePurchaseOrder       = "INSERT INTO purchase_orders(order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id, product_record_id) VALUES (?,?,?,?,?,?,?,?)"
	UpdatePurchaseOrder     = "UPDATE purchase_orders SET order_number=?, order_date=?, tracking_code=?, buyer_id=?, carrier_id=?, order_status_id=?, warehouse_id=?, product_record_id=? WHERE id=?"
	DeletePurchaseOrderByID = "DELETE FROM purchase_orders WHERE id = ?"
	CountByBuyerID          = "SELECT COUNT(*) from purchase_orders where id = ?"
)

type purchaseOrderRepository struct {
	db *sql.DB
}

func NewPurchaseOrderRepository(db *sql.DB) repositories.PurchaseOrderRepository {
	return &purchaseOrderRepository{
		db: db,
	}
}

func (r *purchaseOrderRepository) GetAll(ctx context.Context) ([]entities.PurchaseOrder, error) {
	localities := make([]entities.PurchaseOrder, 0)

	rows, err := r.db.Query(GetAllPurchaseOrders)
	if err != nil {
		return localities, err
	}

	for rows.Next() {
		purchaseOrder := entities.PurchaseOrder{}
		err := rows.Scan(&purchaseOrder.ID, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate, &purchaseOrder.TrackingCode, &purchaseOrder.BuyerID, &purchaseOrder.CarrierID, &purchaseOrder.OrderStatusID, &purchaseOrder.WarehouseID, &purchaseOrder.ProductRecordID)
		if err != nil {
			return localities, err
		}

		localities = append(localities, purchaseOrder)
	}

	return localities, rows.Err()
}

func (r *purchaseOrderRepository) Get(ctx context.Context, id int) (entities.PurchaseOrder, error) {
	row := r.db.QueryRow(GetPurchaseOrderByID, id)
	purchaseOrder := entities.PurchaseOrder{}
	err := row.Scan(&purchaseOrder.ID, &purchaseOrder.OrderNumber, &purchaseOrder.OrderDate, &purchaseOrder.TrackingCode, &purchaseOrder.BuyerID, &purchaseOrder.CarrierID, &purchaseOrder.OrderStatusID, &purchaseOrder.WarehouseID, &purchaseOrder.ProductRecordID)
	if err != nil {
		return entities.PurchaseOrder{}, err
	}

	return purchaseOrder, nil
}

func (r *purchaseOrderRepository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(ExistsPurchaseOrderByID, id)
	var foundId int
	err := row.Scan(&foundId)
	return err == nil
}

func (r *purchaseOrderRepository) Save(ctx context.Context, purchaseOrder entities.PurchaseOrder) (int, error) {
	stmt, err := r.db.Prepare(SavePurchaseOrder)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&purchaseOrder.OrderNumber, &purchaseOrder.OrderDate, &purchaseOrder.TrackingCode, &purchaseOrder.BuyerID, &purchaseOrder.CarrierID, &purchaseOrder.OrderStatusID, &purchaseOrder.WarehouseID, &purchaseOrder.ProductRecordID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *purchaseOrderRepository) Update(ctx context.Context, purchaseOrder entities.PurchaseOrder) error {
	stmt, err := r.db.Prepare(UpdatePurchaseOrder)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&purchaseOrder.OrderNumber, &purchaseOrder.OrderDate, &purchaseOrder.TrackingCode, &purchaseOrder.BuyerID, &purchaseOrder.CarrierID, &purchaseOrder.OrderStatusID, &purchaseOrder.WarehouseID, &purchaseOrder.ProductRecordID, &purchaseOrder.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *purchaseOrderRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeletePurchaseOrderByID)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return services.ErrNotFound
	}

	return nil
}

func (r *purchaseOrderRepository) CountByBuyerID(ctx context.Context, id int) (int, error) {
	count := 0
	row := r.db.QueryRow(CountByBuyerID, id)
	err := row.Scan(&count)

	return count, err
}
