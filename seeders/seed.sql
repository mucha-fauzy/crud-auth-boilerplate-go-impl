-- Insert dummy data into the 'users' table
INSERT INTO users (username, user_type, created_at, created_by, updated_at, updated_by)
VALUES
  ('JohnDoe', 'admin', NOW(), 'System', NOW(), 'System'),
  ('JaneSmith', 'regular', NOW(), 'System', NOW(), 'System'),
  ('MikeJohnson', 'regular', NOW(), 'System', NOW(), 'System'),
  ('EmilyBrown', 'admin', NOW(), 'System', NOW(), 'System'),
  ('DavidWilson', 'regular', NOW(), 'System', NOW(), 'System'),
  ('SophiaDavis', 'regular', NOW(), 'System', NOW(), 'System'),
  ('MatthewLee', 'admin', NOW(), 'System', NOW(), 'System'),
  ('OliviaClark', 'regular', NOW(), 'System', NOW(), 'System'),
  ('EthanMiller', 'regular', NOW(), 'System', NOW(), 'System'),
  ('AvaJohnson', 'admin', NOW(), 'System', NOW(), 'System');

-- Insert dummy data into the 'variants' table with some duplicates and random data
INSERT INTO variants (name, price, stock, status, created_at, created_by, updated_at, updated_by)
SELECT 
  CONCAT('Variant', FLOOR(RAND() * 10) + 1),
  ROUND(RAND() * 100, 2),
  FLOOR(RAND() * 100),
  IF(RAND() < 0.5, 'available', 'out of stock'),
  NOW(), 
  (SELECT username FROM users ORDER BY RAND() LIMIT 1),
  NOW(),
  (SELECT username FROM users ORDER BY RAND() LIMIT 1)
FROM variants
LIMIT 10;

-- Insert dummy data into the 'warehouses'
INSERT INTO warehouses (name, city, province, created_at, created_by, updated_at, updated_by)
VALUES
  ('Warehouse1', 'New York', 'New York', NOW(), 'JohnDoe', NOW(), 'JohnDoe'),
  ('Warehouse2', 'Los Angeles', 'California', NOW(), 'JohnDoe', NOW(), 'JohnDoe'),
  ('Warehouse3', 'Chicago', 'Illinois', NOW(), 'JaneSmith', NOW(), 'JaneSmith'),
  ('Warehouse4', 'Houston', 'Texas', NOW(), 'MikeJohnson', NOW(), 'MikeJohnson'),
  ('Warehouse5', 'Phoenix', 'Arizona', NOW(), 'DavidWilson', NOW(), 'DavidWilson'),
  ('Warehouse6', 'Philadelphia', 'Pennsylvania', NOW(), 'SophiaDavis', NOW(), 'SophiaDavis'),
  ('Warehouse7', 'San Antonio', 'Texas', NOW(), 'MatthewLee', NOW(), 'MatthewLee'),
  ('Warehouse8', 'San Diego', 'California', NOW(), 'OliviaClark', NOW(), 'OliviaClark'),
  ('Warehouse9', 'Dallas', 'Texas', NOW(), 'EthanMiller', NOW(), 'EthanMiller'),
  ('Warehouse10', 'San Jose', 'California', NOW(), 'AvaJohnson', NOW(), 'AvaJohnson');

-- Insert dummy data into the 'images' table
INSERT INTO images (image_url, created_at, created_by, updated_at, updated_by)
SELECT 
  CONCAT('https://example.com/image', FLOOR(RAND() * 10) + 1, '.jpg'),
  NOW(),
  (SELECT username FROM users ORDER BY RAND() LIMIT 1),
  NOW(),
  (SELECT username FROM users ORDER BY RAND() LIMIT 1)
FROM images
LIMIT 10;

-- Insert dummy data into the 'brands' table
INSERT INTO brands (variant_id, image_id, name, created_at, created_by, updated_at, updated_by)
SELECT 
  FLOOR(RAND() * 10) + 1,
  FLOOR(RAND() * 10) + 1,
  CONCAT('Brand', FLOOR(RAND() * 10) + 1),
  NOW(),
  (SELECT username FROM users ORDER BY RAND() LIMIT 1),
  NOW(),
  (SELECT username FROM users ORDER BY RAND() LIMIT 1)
FROM brands
LIMIT 10;

-- Insert dummy data into the 'products' table
INSERT INTO products (user_id, brand_id, warehouse_id, name, created_at, created_by, updated_at, updated_by)
SELECT
  FLOOR(RAND() * 10) + 1,
  FLOOR(RAND() * 10) + 1,
  FLOOR(RAND() * 10) + 1,
  CONCAT('Product', FLOOR(RAND() * 10) + 1),
  NOW(),
  (SELECT username FROM users ORDER BY RAND() LIMIT 1),
  NOW(),
  (SELECT username FROM users ORDER BY RAND() LIMIT 1)
FROM products
LIMIT 10;
