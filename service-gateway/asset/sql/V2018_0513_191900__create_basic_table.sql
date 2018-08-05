-- +migrate Up

CREATE TABLE `User` (
  -- primary key
  `id`         INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at` TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at` TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at` TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_USER_deleted_at` (`deleted_at`),

  -- columns
  `email`      VARCHAR(50)          NOT NULL,
  `phone`      VARCHAR(50)          NULL     DEFAULT NULL,
  `name`       VARCHAR(50)          NULL     DEFAULT NULL,
  `birthday`   VARCHAR(20)          NULL     DEFAULT NULL,
  `address`    TEXT                 NULL     DEFAULT NULL,

  CONSTRAINT `uniq_USER_email` UNIQUE (`email`)

  -- FK columns
);

-- model for `qor/auth`
-- * https://godoc.org/github.com/qor/auth/auth_identity#AuthIdentity
CREATE TABLE `AuthIdentity` (
  -- primary key
  `id`                 int(10) UNSIGNED     NOT NULL               AUTO_INCREMENT PRIMARY KEY,

  -- primary key
  `created_at`         timestamp            NULL                   DEFAULT NULL,
  `updated_at`         timestamp            NULL                   DEFAULT NULL,
  `deleted_at`         timestamp            NULL                   DEFAULT NULL,
  INDEX `idx_AuthIdentity_deleted_at` (`deleted_at`),

  -- columns
  `provider`           varchar(255)         NOT NULL,
  `uid`                varchar(20)          NOT NULL,
  `encrypted_password` TEXT                 NOT NULL,
  CONSTRAINT `uniq_AuthIdentity_uid` UNIQUE (`uid`),

  -- FK columns
  `user_id`            INTEGER(10) UNSIGNED NULL                   DEFAULT NULL,

  FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

CREATE TABLE `Category` (
  -- primary key
  `id`                 INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`         TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`         TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`         TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_Category_deleted_at` (`deleted_at`),

  -- columns
  `name`               VARCHAR(255)         NOT NULL,
  `path`               VARCHAR(255)         NOT NULL,
  `display_name`       VARCHAR(255)         NOT NULL,
  `description`        TEXT                 NOT NULL,

  INDEX `idx_Category_path` (`path`),
  CONSTRAINT `uniq_Category_path` UNIQUE (`path`),

  -- FK columns
  `parent_category_id` INTEGER(10) UNSIGNED NULL,
  INDEX `idx_Category_parent_category_id` (`parent_category_id`),
  CONSTRAINT `fk_Product_parent_category_id`
  FOREIGN KEY (`parent_category_id`) REFERENCES `Category` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);


CREATE TABLE `Image` (
  -- primary key
  `id`         INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at` TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at` TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at` TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_Image_deleted_at` (`deleted_at`),

  -- columns
  `name`       VARCHAR(255)         NOT NULL,
  `type`       VARCHAR(255)         NOT NULL,
  `path`       TEXT                 NOT NULL

  -- FK columns
);

CREATE TABLE `Product` (
  -- primary key
  `id`          INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`  TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_Product_deleted_at` (`deleted_at`),

  -- columns
  `name`        VARCHAR(255)         NOT NULL,
  `price`       INTEGER(10) UNSIGNED NOT NULL,
  `description` TEXT                 NOT NULL,
  `on_sale`     VARCHAR(4)           NOT NULL,

  -- FK columns
  `category_id` INTEGER(10) UNSIGNED NOT NULL,
  `image_id`    INTEGER(10) UNSIGNED NULL,

  CONSTRAINT `fk_Product_category_id`
  FOREIGN KEY (`category_id`) REFERENCES `Category` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_Product_image_id`
  FOREIGN KEY (`image_id`) REFERENCES `Image` (`id`)
    ON DELETE SET NULL
    ON UPDATE CASCADE
);

CREATE TABLE `ProductOption` (
  -- primary key
  `id`          INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`  TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`  TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_ProductOption_deleted_at` (`deleted_at`),

  -- columns
  `name`        VARCHAR(255)         NOT NULL,
  `price`       INTEGER(10) UNSIGNED NOT NULL,
  `description` TEXT                 NOT NULL,
  `on_sale`     VARCHAR(4)           NOT NULL,

  -- FK columns
  `product_id`  INTEGER(10) UNSIGNED NOT NULL,
  `image_id`    INTEGER(10) UNSIGNED NULL,

  CONSTRAINT `fk_ProductOption_product_id`
  FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,
  CONSTRAINT `fk_ProductOption_image_id`
  FOREIGN KEY (`image_id`) REFERENCES `Image` (`id`)
    ON DELETE SET NULL
    ON UPDATE CASCADE
);

CREATE TABLE `Order` (
  -- primary key
  `id`                INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`        TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`        TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`        TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_Order_deleted_at` (`deleted_at`),

  -- columns
  `state`             VARCHAR(30)          NOT NULL,
  `amount`            INTEGER(10) UNSIGNED NOT NULL,

  `shipping_country`  VARCHAR(50)          NOT NULL,
  `shipping_city`     VARCHAR(50)          NOT NULL,
  `shipping_state`    VARCHAR(50)          NOT NULL,
  `shipping_zipcode`  VARCHAR(20)          NOT NULL,
  `shipping_address1` TEXT                 NOT NULL,
  `shipping_address2` TEXT                 NOT NULL,
  `shipping_message`  TEXT                 NOT NULL,

  `orderer_name`      VARCHAR(50)          NOT NULL,
  `orderer_phone`     VARCHAR(50)          NOT NULL,
  `orderer_email`     VARCHAR(50)          NOT NULL,
  `recipient_name`    VARCHAR(50)          NOT NULL,
  `recipient_phone`   VARCHAR(50)          NOT NULL,
  `recipient_email`   VARCHAR(50)          NOT NULL,

  -- FK columns
  `user_id`           INTEGER(10) UNSIGNED NOT NULL,

  CONSTRAINT `fk_Order_user_id`
  FOREIGN KEY (`user_id`) REFERENCES `User` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

CREATE TABLE OrderDetail (
  -- primary key
  `id`                INTEGER(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,

  -- timestamp
  `created_at`        TIMESTAMP            NULL     DEFAULT NULL,
  `updated_at`        TIMESTAMP            NULL     DEFAULT NULL,
  `deleted_at`        TIMESTAMP            NULL     DEFAULT NULL,
  INDEX `idx_OrderDetail_deleted_at` (`deleted_at`),

  -- columns
  `index`             INTEGER(10) UNSIGNED NOT NULL,
  `price`             INTEGER(10) UNSIGNED NOT NULL,
  `quantity`          INTEGER(10) UNSIGNED NOT NULL,
  `amount`            INTEGER(10) UNSIGNED NOT NULL,

  -- FK columns
  `order_id`          INTEGER(10) UNSIGNED NOT NULL,
  `product_id`        INTEGER(10) UNSIGNED NOT NULL,
  `product_option_id` INTEGER(10) UNSIGNED NOT NULL,

  CONSTRAINT `fk_OrderDetail_order_id`
  FOREIGN KEY (`order_id`) REFERENCES `Order` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,

  CONSTRAINT `fk_OrderDetail_product_id`
  FOREIGN KEY (`product_id`) REFERENCES `Product` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE,


  CONSTRAINT `fk_OrderDetail_product_option_id`
  FOREIGN KEY (`product_option_id`) REFERENCES `ProductOption` (`id`)
    ON DELETE RESTRICT
    ON UPDATE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS `User`;
DROP TABLE IF EXISTS `Category`;
DROP TABLE IF EXISTS `Image`;
DROP TABLE IF EXISTS `Product`;
DROP TABLE IF EXISTS `Order`;
DROP TABLE IF EXISTS `OrderDetail`;


