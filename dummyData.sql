-- Inserting products
INSERT INTO products (product_name, price, stock) VALUES 
('Product A', 10.99, 100),
('Product B', 20.50, 75),
('Product C', 15.75, 50),
('Product D', 8.25, 120),
('Product E', 30.00, 90),
('Product F', 12.49, 60),
('Product G', 25.99, 80),
('Product H', 18.75, 110),
('Product I', 22.50, 70),
('Product J', 9.99, 95);

-- Inserting sales
INSERT INTO sales (product_id, quantity, sale_date) VALUES
(1, 5, '2024-05-01'),
(2, 8, '2024-05-02'),
(3, 3, '2024-05-03'),
(4, 10, '2024-05-04'),
(5, 6, '2024-05-05'),
(6, 2, '2024-05-06'),
(7, 7, '2024-05-07'),
(8, 4, '2024-05-08'),
(9, 9, '2024-05-09'),
(10, 1, '2024-05-10'),
(1, 3, '2024-05-11'),
(2, 6, '2024-05-12'),
(3, 2, '2024-05-13'),
(4, 7, '2024-05-14'),
(5, 4, '2024-05-15');

-- Inserting staff
INSERT INTO staff (staff_name, email, position) VALUES
('John Doe', 'john.doe@example.com', 'Sales Manager'),
('Jane Smith', 'jane.smith@example.com', 'Sales Associate'),
('Bob Johnson', 'bob.johnson@example.com', 'Store Manager'),
('Alice Lee', 'alice.lee@example.com', 'Cashier'),
('Michael Brown', 'michael.brown@example.com', 'Stock Clerk');

INSERT INTO users (username, password, security_question, security_answer) VALUES
('john_doe', 'password123', 'What is your mother\'s maiden name?', 'Smith'),
('jane_smith', 'letmein', 'What city were you born in?', 'New York'),
('bob_johnson', 'securepass', 'What is your favorite color?', 'Blue'),
('alice_lee', 'abc123', 'What is the name of your first pet?', 'Fluffy'),
('michael_brown', 'qwerty', 'What is the make and model of your first car?', 'Toyota Corolla');
