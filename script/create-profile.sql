CREATE TABLE profiles (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    mobile VARCHAR(10),
    sex VARCHAR(1) CHECK (sex IN ('M', 'F')),
    status VARCHAR(1) CHECK (status IN ('A', 'I')),
    image VARCHAR(255),
    user_id BIGINT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP(3)
);