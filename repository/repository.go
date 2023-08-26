package repository

import (
	"database/sql"
	"time"

	"github.com/cecepsprd/inventory-app/model"
)

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) InventoryRepository {
	return InventoryRepository{
		db: db,
	}
}

func (r *InventoryRepository) StoreTransactionInHeader(req model.TransactionInHeader) (int64, error) {
	query := `INSERT INTO trx_in_header (trx_in_no, warehouse_id, supplier_id, trx_in_date, notes) VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.TrxInNo, req.WarehouseID, req.SupplierID, time.Now(), req.Notes)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *InventoryRepository) StoreTransactionInDetail(req model.TransactionInDetail) error {
	query := `INSERT INTO trx_in_detail (trx_in_id, trx_in_product_id, trx_in_quantity_dus, trx_in_quantity_pcs) VALUES (?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.TrxInID, req.TrxInProductID, req.TrxInQuantityDus, req.TrxInQuantityPcs)
	if err != nil {
		return err
	}

	return nil
}

func (r *InventoryRepository) StoreTransactionOutHeader(req model.TransactionOutHeader) (int64, error) {
	query := `INSERT INTO trx_out_header (trx_out_no, warehouse_id, supplier_id, trx_out_date, notes) VALUES (?, ?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.TrxOutNo, req.WarehouseID, req.SupplierID, time.Now(), req.Notes)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *InventoryRepository) StoreTransactionOutDetail(req model.TransactionOutDetail) error {
	query := `INSERT INTO trx_out_detail (trx_out_id, trx_out_product_id, trx_out_quantity_dus, trx_out_quantity_pcs) VALUES (?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.TrxOutID, req.TrxOutProductID, req.TrxOutQuantityDus, req.TrxOutQuantityPcs)
	if err != nil {
		return err
	}

	return nil
}

func (r *InventoryRepository) StoreProductWarehouseOutDetail(req model.ProductWarehouse) error {
	query := `INSERT INTO product_warehouse (product_id, warehouse_id, quantity_dus, quantity_pcs) VALUES (?, ?, ?, ?)`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.ProductID, req.WarehouseID, req.QuantityDus, req.QuantityPcs)
	if err != nil {
		return err
	}

	return nil
}

func (r *InventoryRepository) StockReport() (reports []model.StockReport, err error) {
	query := `select p.name as product_name, 
			  w.name as warehouse_name,
			  sum(pw.quantity_dus) as quantity_dus,
			  sum(pw.quantity_pcs) as quantity_pcs 
			  from product p 
			  inner join product_warehouse pw on pw.product_id = p.id 
			  inner join warehouse w on w.id = pw.warehouse_id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		sr := model.StockReport{}
		err = rows.Scan(
			&sr.ProductName,
			&sr.WarehouseName,
			&sr.QuantityDus,
			&sr.QuantityPcs,
		)

		if err != nil {
			return nil, err
		}

		reports = append(reports, sr)
	}

	return reports, nil
}
