INSERT INTO users (name, email, password, deposit_amount, jwt_token) VALUES
('Andi', 'andi@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 100000, NULL),
('Budi', 'budi@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 150000, NULL),
('Citra', 'citra@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 200000, NULL),
('Dewi', 'dewi@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 120000, NULL),
('Eka', 'eka@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 130000, NULL),
('Fajar', 'fajar@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 140000, NULL),
('Gina', 'gina@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 160000, NULL),
('Hadi', 'hadi@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 170000, NULL),
('Intan', 'intan@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 180000, NULL),
('Joko', 'joko@mail.com', crypt('Rahasiabanget123', gen_salt('bf')), 190000, NULL);


INSERT INTO books (name, stock_availability, rental_cost, category) VALUES
('Database Fundamentals', 5, 10000, 'Education'),
('Clean Code', 3, 15000, 'Programming'),
('Harry Potter', 10, 8000, 'Fantasy'),
('Lord of the Rings', 7, 12000, 'Fantasy'),
('Atomic Habits', 4, 11000, 'Self Development'),
('Deep Work', 6, 13000, 'bookivity'),
('Python Crash Course', 5, 14000, 'Programming'),
('The Hobbit', 8, 9000, 'Fantasy'),
('Rich Dad Poor Dad', 9, 10000, 'Finance'),
('Sapiens', 6, 15000, 'History');


INSERT INTO rent_books (user_id, book_id, duration, created_at, due_date_rent) VALUES
(1, 1, 3, NOW(), NOW() + INTERVAL '3 days'),
(2, 2, 5, NOW(), NOW() + INTERVAL '5 days'),
(3, 3, 2, NOW(), NOW() + INTERVAL '2 days'),
(4, 4, 7, NOW(), NOW() + INTERVAL '7 days'),
(5, 5, 4, NOW(), NOW() + INTERVAL '4 days'),
(6, 6, 6, NOW(), NOW() + INTERVAL '6 days'),
(7, 7, 3, NOW(), NOW() + INTERVAL '3 days'),
(8, 8, 2, NOW(), NOW() + INTERVAL '2 days'),
(9, 9, 5, NOW(), NOW() + INTERVAL '5 days'),
(10, 10, 1, NOW(), NOW() + INTERVAL '1 day');
