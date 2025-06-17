CREATE TABLE scheduled_transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    type ENUM('credit', 'debit') NOT NULL,
    currency VARCHAR(3) NOT NULL DEFAULT 'TRY',
    amount DECIMAL(10,2) NOT NULL,
    scheduled_time DATETIME NOT NULL,
    processed BOOLEAN DEFAULT FALSE,
    executed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);