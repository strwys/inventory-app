package service

import (
	"github.com/cecepsprd/inventory-app/model"
	"github.com/cecepsprd/inventory-app/repository"
	"github.com/cecepsprd/inventory-app/utils"
	"github.com/cecepsprd/inventory-app/utils/logger"
)

type InventoryService struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return InventoryService{
		repo: repo,
	}
}

func (s *InventoryService) InventoryIn(req model.InventoryInRequest) error {
	transactionNumber := utils.GenerateTransactionNumber("IN")

	trxID, err := s.repo.StoreTransactionInHeader(model.TransactionInHeader{
		TrxInNo:     transactionNumber,
		WarehouseID: req.WarehouseID,
		SupplierID:  req.SupplierID,
		Notes:       req.Notes,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = s.repo.StoreTransactionInDetail(model.TransactionInDetail{
		TrxInID:          trxID,
		TrxInProductID:   req.ProductID,
		TrxInQuantityDus: req.QuantityDus,
		TrxInQuantityPcs: req.QuantityPcs,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = s.repo.StoreProductWarehouseOutDetail(model.ProductWarehouse{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		QuantityDus: req.QuantityDus,
		QuantityPcs: req.QuantityPcs,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *InventoryService) InventoryOut(req model.InventoryOutRequest) error {
	transactionNumber := utils.GenerateTransactionNumber("OUT")

	trxOutID, err := s.repo.StoreTransactionOutHeader(model.TransactionOutHeader{
		TrxOutNo:    transactionNumber,
		WarehouseID: req.WarehouseID,
		SupplierID:  req.SupplierID,
		Notes:       req.Notes,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = s.repo.StoreTransactionOutDetail(model.TransactionOutDetail{
		TrxOutID:          trxOutID,
		TrxOutProductID:   req.ProductID,
		TrxOutQuantityDus: req.QuantityDus,
		TrxOutQuantityPcs: req.QuantityPcs,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	err = s.repo.StoreProductWarehouseOutDetail(model.ProductWarehouse{
		ProductID:   req.ProductID,
		WarehouseID: req.WarehouseID,
		QuantityDus: -req.QuantityDus,
		QuantityPcs: -req.QuantityPcs,
	})

	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}

	return nil
}

func (s *InventoryService) StockReport() ([]model.StockReport, error) {
	reports, err := s.repo.StockReport()
	if err != nil {
		logger.Log.Error(err.Error())
		return nil, err
	}
	return reports, err
}
