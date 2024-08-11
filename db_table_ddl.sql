CREATE TABLE IF NOT EXISTS Orders (
	id INT PRIMARY KEY,
	created_on TIMESTAMP,
	price DECIMAL,
	category TEXT
);
