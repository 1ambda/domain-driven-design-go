-- +migrate Up
INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (1, CURRENT_TIMESTAMP(), 'LG Notebook GRAM 13', '1350000', 'LG gram Thin & Light Laptop - Up to 16.5 hrs, Thunderbolt 3, Finger Print Reader (Windows not installed)', 'Y', 5, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (1, CURRENT_TIMESTAMP(), 'Memory 8 GB+', '160000', '', 'Y', 1, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (2, CURRENT_TIMESTAMP(), 'SSD 256 GiB+', '100000', '', 'Y', 1, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (3, CURRENT_TIMESTAMP(), 'OS Installation', '150000', '', 'Y', 1, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (2, CURRENT_TIMESTAMP(), 'LG Notebook GRAM 15', '1550000', 'Windows Installed', 'N', 5, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (4, CURRENT_TIMESTAMP(), 'Memory 4 GB+', '160000', '', 'N', 2, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (5, CURRENT_TIMESTAMP(), 'SSD 128 GiB+', '100000', '', 'Y', 2, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (3, CURRENT_TIMESTAMP(), 'Samsung 7 Pro', '1450000', 'OS is installed', 'Y', 5, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (6, CURRENT_TIMESTAMP(), 'Memory 4 GB+', '160000', '', 'Y', 3, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (7, CURRENT_TIMESTAMP(), 'SSD 128 GiB+', '100000', '', 'Y', 3, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (4, CURRENT_TIMESTAMP(), 'Samsung 8 Pro', '1950000', 'OS is installed', 'Y', 5, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (5, CURRENT_TIMESTAMP(), 'MS Surface Pro 2', '1100000', '', 'Y', 7, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (8, CURRENT_TIMESTAMP(), 'Surface Pen', '150000', '', 'Y', 5, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (6, CURRENT_TIMESTAMP(), 'MS Surface Pro 3', '1300000', '', 'Y', 7, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (9, CURRENT_TIMESTAMP(), 'Surface Pen', '150000', '', 'Y', 6, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (7, CURRENT_TIMESTAMP(), 'Intel Core 8 i5-4600 (Bulk)', '224000', '', 'Y', 3, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (8, CURRENT_TIMESTAMP(), 'Intel Core 8 i5-8600 (Bulk)', '272000', '', 'Y', 3, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (9, CURRENT_TIMESTAMP(), 'Intel Core 8 i7-8700 (Bulk)', '434000', '', 'Y', 3, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (10, CURRENT_TIMESTAMP(), 'Intel Core 8 i7-8086K (Bulk)', '572400', '', 'Y', 3, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (11, CURRENT_TIMESTAMP(), 'AMD Ryzen 7 2700', '359000', '', 'Y', 3, null);
INSERT INTO `ProductOption` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `product_id`, `image_id`)
VALUES (10, CURRENT_TIMESTAMP(), 'AMD Ryzen 7 2700X', '30000', '', 'N', 11, null);

INSERT INTO `Product` (`id`, `created_at`, `name`, `price`, `description`, `on_sale`, `category_id`, `image_id`)
VALUES (12, CURRENT_TIMESTAMP(), 'AMD Ryzen 5 2600', '234000', '', 'Y', 3, null);


-- +migrate Down
SELECT 1;

