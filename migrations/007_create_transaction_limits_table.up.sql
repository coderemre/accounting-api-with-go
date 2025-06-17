CREATE TABLE IF NOT EXISTS transaction_limits (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    type ENUM('credit', 'debit', 'transfer') NOT NULL,
    max_amount DECIMAL(10,2),
    max_count INT,
    period ENUM('daily', 'weekly', 'monthly') NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY unique_user_type_period (user_id, type, period),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);