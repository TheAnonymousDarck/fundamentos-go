CREATE TABLE meli_items {
    id VARCHAR(15) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    conditions VARCHAR(20) NOT NULL,
    category_id VARCHAR(50) NOT NULL,
    url TEXT NOT NULL 
}