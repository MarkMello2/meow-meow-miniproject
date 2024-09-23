CREATE TABLE malls (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256),
    description VARCHAR(1024),
    image VARCHAR(256),
    created_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3),
    updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3),
    deleted_at TIMESTAMP(3)
);
