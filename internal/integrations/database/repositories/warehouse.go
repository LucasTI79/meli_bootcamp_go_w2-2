package repositories

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

// Warehouse queries
const (
	GetAllWarehouses    = "SELECT warehouses.id, warehouses.address, warehouses.telephone, warehouses.warehouse_code, warehouses.minimum_capacity, warehouses.minimum_temperature FROM warehouses"
	GetWarehouseByID    = "SELECT warehouses.id, warehouses.address, warehouses.telephone, warehouses.warehouse_code, warehouses.minimum_capacity, warehouses.minimum_temperature FROM warehouses WHERE id=?"
	ExistsWarehouseByID = "SELECT id FROM warehouses WHERE id=?"
	SaveWarehouse       = "INSERT INTO warehouses(address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?,?,?,?,?)"
	UpdateWarehouse     = "UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"
	DeleteWarehouseByID = "DELETE FROM warehouses WHERE id=?"
)

type warehouseRepository struct {
	db *sql.DB
}

func NewWarehouseRepository(db *sql.DB) repositories.WarehouseRepository {
	return &warehouseRepository{
		db: db,
	}
}

func (r *warehouseRepository) GetAll(ctx context.Context) ([]entities.Warehouse, error) {
	rows, err := r.db.Query(GetAllWarehouses)
	if err != nil {
		return nil, err
	}

	var warehouses []entities.Warehouse

	for rows.Next() {
		w := entities.Warehouse{}
		_ = rows.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
		warehouses = append(warehouses, w)
	}

	return warehouses, nil
}

func (r *warehouseRepository) Get(ctx context.Context, id int) (entities.Warehouse, error) {
	row := r.db.QueryRow(GetWarehouseByID, id)
	w := entities.Warehouse{}
	err := row.Scan(&w.ID, &w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		return entities.Warehouse{}, err
	}

	return w, nil
}

func (r *warehouseRepository) Exists(ctx context.Context, warehouseCode string) bool {
	row := r.db.QueryRow(ExistsWarehouseByID, warehouseCode)
	err := row.Scan(&warehouseCode)
	return err == nil
}

func (r *warehouseRepository) Save(ctx context.Context, w entities.Warehouse) (int, error) {
	stmt, err := r.db.Prepare(SaveWarehouse)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *warehouseRepository) Update(ctx context.Context, w entities.Warehouse) error {
	stmt, err := r.db.Prepare(UpdateWarehouse)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&w.Address, &w.Telephone, &w.WarehouseCode, &w.MinimumCapacity, &w.MinimumTemperature, &w.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *warehouseRepository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DeleteWarehouseByID)
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
