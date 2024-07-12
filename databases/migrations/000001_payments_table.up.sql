CREATE TABLE IF NOT EXISTS payments (
    id UUID DEFAULT GEN_RANDOM_UUID() PRIMARY KEY,
    reservation_id UUID NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    payment_method VARCHAR(100),
    payment_status VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
)