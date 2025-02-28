CREATE TABLE IF NOT EXISTS monsters (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,  -- Ensure that the name is unique
    generation INT,
    types JSON,
    species VARCHAR(100) NOT NULL,
    height VARCHAR(50) NOT NULL,
    weight VARCHAR(50) NOT NULL,
    abilities JSON,
    image VARCHAR(150) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL 
);
