CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INTEGER NOT NULL,
    name VARCHAR(50) NOT NULL,
    price NUMERIC NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    starting_date TIMESTAMP NOT NULL,
    payment_date TIMESTAMP NOT NULL,
    subscription_renew TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT chk_subscription_renew
        CHECK (subscription_renew IN ('daily', 'weekly', 'monthly', 'yearly'))
);