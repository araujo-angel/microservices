CREATE DATABASE IF NOT EXISTS `order`;
CREATE DATABASE IF NOT EXISTS `payment`;

-- Criar tabela de estoque
CREATE TABLE IF NOT EXISTS `order`.`stock_items` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `product_code` varchar(255) DEFAULT NULL,
  `name` longtext,
  `quantity` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_stock_items_product_code` (`product_code`),
  KEY `idx_stock_items_deleted_at` (`deleted_at`)
);

-- Inserir produtos de exemplo
INSERT INTO `order`.`stock_items` (product_code, name, quantity, created_at, updated_at) VALUES
('prod', 'Produto Teste', 100, NOW(), NOW()),
('A', 'Produto A', 50, NOW(), NOW()),
('B', 'Produto B', 30, NOW(), NOW()),
('C', 'Produto C', 25, NOW(), NOW()),
('D', 'Produto D', 15, NOW(), NOW());