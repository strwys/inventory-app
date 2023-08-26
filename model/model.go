package model

import "time"

type InventoryInRequest struct {
	WarehouseID int64  `json:"warehouse_id"`
	SupplierID  int64  `json:"supplier_id"`
	ProductID   int64  `json:"product_id"`
	QuantityDus int    `json:"quantity_dus"`
	QuantityPcs int    `json:"quantity_pcs"`
	Notes       string `json:"notes"`
}

type TransactionInHeader struct {
	ID          int64     `json:"id"`
	TrxInNo     string    `json:"trx_in_no"`
	WarehouseID int64     `json:"warehouse_id"`
	SupplierID  int64     `json:"supplier_id"`
	TrxInDate   time.Time `json:"trx_in_date"`
	Notes       string    `json:"notes"`
}

type TransactionInDetail struct {
	ID               int64 `json:"id"`
	TrxInID          int64 `json:"trx_in_no"`
	TrxInProductID   int64 `json:"trx_in_product_id"`
	TrxInQuantityDus int   `json:"trx_in_quantity_dus"`
	TrxInQuantityPcs int   `json:"trx_in_quantity_pcs"`
}

type InventoryOutRequest struct {
	WarehouseID int64  `json:"warehouse_id"`
	SupplierID  int64  `json:"supplier_id"`
	ProductID   int64  `json:"product_id"`
	QuantityDus int    `json:"quantity_dus"`
	QuantityPcs int    `json:"quantity_pcs"`
	Notes       string `json:"notes"`
}

type TransactionOutHeader struct {
	ID          int64     `json:"id"`
	TrxOutNo    string    `json:"trx_in_no"`
	WarehouseID int64     `json:"warehouse_id"`
	SupplierID  int64     `json:"supplier_id"`
	TrxInDate   time.Time `json:"trx_in_date"`
	Notes       string    `json:"notes"`
}

type TransactionOutDetail struct {
	ID                int64 `json:"id"`
	TrxOutID          int64 `json:"trx_out_no"`
	TrxOutProductID   int64 `json:"trx_out_product_id"`
	TrxOutQuantityDus int   `json:"trx_out_quantity_dus"`
	TrxOutQuantityPcs int   `json:"trx_out_quantity_pcs"`
}

type ProductWarehouse struct {
	ID          int64 `json:"id"`
	ProductID   int64 `json:"product_id"`
	WarehouseID int64 `json:"warehouse_id"`
	QuantityDus int   `json:"quantity_dus"`
	QuantityPcs int   `json:"quantity_pcs"`
}

type StockReport struct {
	ProductName   string `json:"product_name"`
	WarehouseName string `json:"warehouse_name"`
	QuantityDus   int    `json:"quantity_dus"`
	QuantityPcs   int    `json:"quantity_pcs"`
}
