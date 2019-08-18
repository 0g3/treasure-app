CREATE TABLE IF NOT EXISTS users (
    id VARCHAR (128) NOT NULL,
    display_id VARCHAR (20) NOT NULL,
    name VARCHAR (20) NOT NULL,
    description VARCHAR (20) NOT NULL DEFAULT '',
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY (display_id)
);