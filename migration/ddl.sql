
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    deposit_amount NUMERIC(10,2) DEFAULT 0,
    jwt_token TEXT
);

CREATE TABLE books (
    book_id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    stock_availability INTEGER NOT NULL CHECK (stock_availability >= 0),
    rental_cost NUMERIC(10,2) NOT NULL,
    category VARCHAR(100)
);

CREATE TABLE rent_books (
    rent_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    book_id INTEGER NOT NULL,
    duration INTEGER NOT NULL CHECK (duration > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    due_date_rent TIMESTAMP NOT NULL,

    -- Foreign Keys
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_book
        FOREIGN KEY (book_id)
        REFERENCES books(book_id)
        ON DELETE CASCADE
);

CREATE INDEX idx_rent_user ON rent_books(user_id);
CREATE INDEX idx_rent_book ON rent_books(book_id);