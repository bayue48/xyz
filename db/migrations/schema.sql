-- =========================================
-- DATABASE
-- =========================================
CREATE DATABASE IF NOT EXISTS xyz_multifinance
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

USE xyz_multifinance;

-- =========================================
-- CUSTOMERS
-- =========================================
CREATE TABLE customers (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    nik VARCHAR(20) NOT NULL UNIQUE,
    full_name VARCHAR(150) NOT NULL,
    legal_name VARCHAR(150) NOT NULL,
    birth_place VARCHAR(100),
    birth_date DATE,
    salary BIGINT NOT NULL,
    ktp_photo VARCHAR(255),
    selfie_photo VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    token VARCHAR(36) UNIQUE
) ENGINE=InnoDB;

-- =========================================
-- CUSTOMER LIMITS
-- Each customer has limits per tenor
-- =========================================
CREATE TABLE customer_limits (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    tenor_month INT NOT NULL,
    limit_amount BIGINT NOT NULL,
    used_amount BIGINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_customer_limit_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_customer_tenor UNIQUE (customer_id, tenor_month)
) ENGINE=InnoDB;

-- =========================================
-- TRANSACTIONS
-- =========================================
CREATE TABLE transactions (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    contract_number VARCHAR(50) NOT NULL UNIQUE,
    customer_id BIGINT NOT NULL,
    otr BIGINT NOT NULL,
    admin_fee BIGINT NOT NULL,
    installment_amount BIGINT NOT NULL,
    interest_amount BIGINT NOT NULL,
    asset_name VARCHAR(150) NOT NULL,
    tenor_month INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_transaction_customer
        FOREIGN KEY (customer_id)
        REFERENCES customers(id)
        ON DELETE RESTRICT
) ENGINE=InnoDB;

-- =========================================
-- INDEXES (PERFORMANCE)
-- =========================================
CREATE INDEX idx_customer_limits_customer
    ON customer_limits (customer_id);

CREATE INDEX idx_transactions_customer
    ON transactions (customer_id);

CREATE INDEX idx_transactions_created_at
    ON transactions (created_at);
