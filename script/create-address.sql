CREATE TABLE address (
    id BIGSERIAL PRIMARY KEY,  -- auto-incrementing primary key
    first_name VARCHAR(255) NOT NULL,  -- recipient's first name
    last_name VARCHAR(255) NOT NULL,  -- recipient's last name
    mobile VARCHAR(10),  -- phone number (optional)
    address VARCHAR(1024),  -- detailed address (optional)
    type BIGINT DEFAULT 0,  -- 1 for primary address, 0 otherwise
    user_id BIGINT REFERENCES users(id),  -- foreign key referencing 'users' table
    created_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,  -- timestamp for record creation
    updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,  -- timestamp for record updates
    deleted_at TIMESTAMP(3)  -- timestamp for soft deletes (optional)
);

