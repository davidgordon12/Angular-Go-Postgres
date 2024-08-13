CREATE TABLE IF NOT EXISTS Orders (
	id INT PRIMARY KEY,
	created_on TIMESTAMP,
	price DECIMAL,
	category TEXT
);

INSERT INTO Orders (
    1, CURRENT_TIMESTAMP, 4.29, "WASHER",
    2, CURRENT_TIMESTAMP, 4.30, "DRYER",
    3, CURRENT_TIMESTAMP, 4.29, "COMBO",
    4, CURRENT_TIMESTAMP, 4.28, "WASHER",
    5, CURRENT_TIMESTAMP, 4.31, "WASHER",
    6, CURRENT_TIMESTAMP, 5.27, "WASHER",
    7, CURRENT_TIMESTAMP, 4.25, "DRYER",
    8, CURRENT_TIMESTAMP, 4.27, "COMBO",
    9, CURRENT_TIMESTAMP, 4.26, "COMBO",
    10, CURRENT_TIMESTAMP, 4.26, "COMBO",
    11, CURRENT_TIMESTAMP, 4.44, "COMBO",
    12, CURRENT_TIMESTAMP, 4.27, "COMBO",
    13, CURRENT_TIMESTAMP, 4.28, "WASHER",
    14, CURRENT_TIMESTAMP, 4.25, "WASHER",
    15, CURRENT_TIMESTAMP, 4.27, "WASHER",
    16, CURRENT_TIMESTAMP, 4.25, "WASHER",
    17, CURRENT_TIMESTAMP, 4.25, "WASHER",
    18, CURRENT_TIMESTAMP, 4.25, "WASHER",
    19, CURRENT_TIMESTAMP, 4.25, "WASHER",
    20, CURRENT_TIMESTAMP, 4.25, "WASHER",
    21, CURRENT_TIMESTAMP, 4.25, "WASHER",
    22, CURRENT_TIMESTAMP, 5.25, "WASHER",
)