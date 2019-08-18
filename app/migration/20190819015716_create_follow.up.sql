CREATE TABLE IF NOT EXISTS follows (
    user_id VARCHAR (128) NOT NULL,
    followed_user_id VARCHAR (128) NOT NULL,
    PRIMARY KEY (user_id, followed_user_id),
    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON UPDATE RESTRICT
        ON DELETE CASCADE,
    FOREIGN KEY (followed_user_id)
        REFERENCES users(id)
        ON UPDATE RESTRICT
        ON DELETE CASCADE
);