CREATE TABLE `product` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(255),
  PRIMARY KEY (`id`)
);

CREATE TABLE `customer` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(255),
  PRIMARY KEY (`id`)
);

CREATE TABLE `warehouse` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(255),
  PRIMARY KEY (`id`)
);

CREATE TABLE product_warehouse (
  `id` INT AUTO_INCREMENT PRIMARY KEY,
  `product_id` INT,
  `warehouse_id` INT,
  `quantity_dus` INT DEFAULT 0,
  `quantity_pcs` INT DEFAULT 0,  
  FOREIGN KEY (warehouse_id) REFERENCES warehouse(id),
  FOREIGN KEY (product_id) REFERENCES product(id)
);

CREATE TABLE `supplier` (
  `id` INT AUTO_INCREMENT,
  `name` VARCHAR(255),
  PRIMARY KEY (`id`)
);

CREATE TABLE `trx_in_header` (
  `id` INT AUTO_INCREMENT,
  `trx_in_no` VARCHAR(50),
  `warehouse_id` INT,
  `supplier_id` INT,
  `transaction_in_date` DATE,
  `notes` TEXT,
  PRIMARY KEY (`id`)
);

CREATE TABLE `trx_in_detail` (
  `id` INT AUTO_INCREMENT,
  `trx_in_id` INT,
  `trx_in_product_id` INT,
  `trx_in_quantity_dus` INT,
  `trx_in_quantity_pcs` INT,
  PRIMARY KEY (`id`)
);

CREATE TABLE `trx_out_header` (
  `id` INT AUTO_INCREMENT,
  `trx_out_no` VARCHAR(50),
  `warehouse_id` INT,
  `supplier_id` INT,
  `transaction_out_date` DATE,
  `notes` TEXT,
  PRIMARY KEY (`id`)
);

CREATE TABLE `trx_out_detail` (
  `id` INT AUTO_INCREMENT,
  `trx_out_id` INT,
  `trx_out_product_id` INT,
  `trx_out_quantity_dus` INT,
  `trx_out_quantity_pcs` INT,
  PRIMARY KEY (`id`)
);