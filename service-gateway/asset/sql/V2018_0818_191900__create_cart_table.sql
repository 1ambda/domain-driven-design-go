-- +migrate Up
CREATE TABLE `Cart` (
  -- primary key
  `id`          INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`  TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_Cart_deleted_at` (`deleted_at`),

  -- columns
  `total_price` INTEGER(10) UNSIGNED NOT NULL,

  -- FK columns
  `user_id`     INTEGER(10) UNSIGNED NOT NULL,

  CONSTRAINT `fk_Cart_user_id`
  FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

CREATE TABLE `CartItem` (
  -- primary key
  `id`            INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`    TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`    TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`    TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_CartItem_deleted_at` (`deleted_at`),

  -- columns
  `index`         INTEGER(10) UNSIGNED NOT NULL,
  `quantity`      INTEGER(10) UNSIGNED NOT NULL,
  `product_price` INTEGER(10) UNSIGNED NOT NULL,
  `total_price`   INTEGER(10) UNSIGNED NOT NULL,

  -- FK columns
  `cart_id`       INTEGER(10) UNSIGNED NOT NULL,
  CONSTRAINT `fk_CartItem_cart_id`
  FOREIGN KEY (`cart_id`) REFERENCES `Cart` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,

  `product_id`    INTEGER(10) UNSIGNED NOT NULL,
  CONSTRAINT `fk_CartItem_product_id`
  FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

CREATE TABLE `CartItemOption` (
  -- primary key
  `id`                   INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`           TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`           TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`           TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_CartItemOption_deleted_at` (`deleted_at`),

  -- columns
  `quantity`             INTEGER(10) UNSIGNED NOT NULL,
  `product_option_price` INTEGER(10) UNSIGNED NOT NULL,

  -- FK columns
  `cart_item_id`         INTEGER(10) UNSIGNED NOT NULL,
  CONSTRAINT `fk_CartItemOption_cart_item_id`
  FOREIGN KEY (`cart_item_id`) REFERENCES `CartItem` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,

  `product_option_id`    INTEGER(10) UNSIGNED NOT NULL,
  CONSTRAINT `fk_CartItemOption_product_id`
  FOREIGN KEY (`product_option_id`) REFERENCES `ProductOption` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS `CartItemOption`;
DROP TABLE IF EXISTS `CartItem`;
DROP TABLE IF EXISTS `Cart`;
