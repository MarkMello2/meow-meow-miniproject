CREATE TABLE favorites (
    id BIGSERIAL PRIMARY KEY,
    price FLOAT NOT NULL,
    quantity BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP(3),
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES products,
    CONSTRAINT fk_userId FOREIGN KEY (user_id) REFERENCES users
);