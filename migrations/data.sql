CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  user_type ENUM('admin', 'regular') NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(50) NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50)
);

CREATE TABLE variants (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price FLOAT NOT NULL,
  stock INT NOT NULL,
  status VARCHAR(20) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(50) NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50)
);

CREATE TABLE warehouses (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  city VARCHAR(100) NOT NULL,
  province VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(50) NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50)
);

CREATE TABLE images (
  id INT AUTO_INCREMENT PRIMARY KEY,
  image_url VARCHAR(200) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(50) NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50)
);

CREATE TABLE brands (
  id INT AUTO_INCREMENT PRIMARY KEY,
  variant_id INT NOT NULL,
  image_id INT NOT NULL,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(50) NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50),
  FOREIGN KEY (variant_id) REFERENCES variants(id),
  FOREIGN KEY (image_id) REFERENCES images(id)
);

CREATE TABLE products (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  brand_id INT NOT NULL,
  warehouse_id INT NOT NULL,
  name VARCHAR(200) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  created_by VARCHAR(50) NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  updated_by VARCHAR(50) NOT NULL,
  deleted_at TIMESTAMP,
  deleted_by VARCHAR(50),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (brand_id) REFERENCES brands(id),
  FOREIGN KEY (warehouse_id) REFERENCES warehouses(id)
);