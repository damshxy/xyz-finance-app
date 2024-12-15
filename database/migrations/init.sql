CREATE TABLE "consumers" (
    id SERIAL PRIMARY KEY,
    nik VARCHAR(16) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255),
    birth_place VARCHAR(255),
    birth_date DATE,
    salary DECIMAL(10, 2),
    ktp_photo TEXT,
    selfie_photo TEXT
);

CREATE TABLE "transactions" (
    id SERIAL PRIMARY KEY,
    consumer_id INT,
    contract_number VARCHAR(50),
    otr DECIMAL(10, 2),
    admin_fee DECIMAL(10, 2),
    installment DECIMAL(10, 2),
    interest DECIMAL(10, 2),
    asset_name VARCHAR(255),
    FOREIGN KEY (consumer_id) REFERENCES consumers(id)
);
