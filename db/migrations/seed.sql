USE xyz_multifinance;

-- =========================================
-- CUSTOMERS
-- =========================================
INSERT INTO customers (
    nik,
    full_name,
    legal_name,
    birth_place,
    birth_date,
    salary,
    ktp_photo,
    selfie_photo,
    token
) VALUES
(
    '3175010101010001',
    'Budi Santoso',
    'Budi Santoso',
    'Jakarta',
    '1990-01-01',
    8000000,
    'ktp_budi.jpg',
    'selfie_budi.jpg',
    'b2e0a063-5dca-465c-8289-59cc4ff94234'
),
(
    '3175020202020002',
    'Annisa Putri',
    'Annisa Putri',
    'Bandung',
    '1992-02-02',
    15000000,
    'ktp_annisa.jpg',
    'selfie_annisa.jpg',
    'cf0ecf8f-94cd-439f-a374-84071112d3dd'
);

-- =========================================
-- CUSTOMER LIMITS
-- Budi: 1,2,3,6 months
-- =========================================
INSERT INTO customer_limits (
    customer_id,
    tenor_month,
    limit_amount
) VALUES
-- Budi (ID = 1)
(1, 1, 100000),
(1, 2, 200000),
(1, 3, 500000),
(1, 6, 700000);

-- =========================================
-- CUSTOMER LIMITS
-- Annisa: 1,2,3,6 months
-- =========================================
INSERT INTO customer_limits (
    customer_id,
    tenor_month,
    limit_amount
) VALUES
-- Annisa (ID = 2)
(2, 1, 1000000),
(2, 2, 1200000),
(2, 3, 1500000),
(2, 6, 2000000);
