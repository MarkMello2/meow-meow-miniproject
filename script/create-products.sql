CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,             
    code VARCHAR(255) UNIQUE NOT NULL,    
    name VARCHAR(255) UNIQUE NOT NULL,    
    description VARCHAR(1024) NOT NULL,   
    price FLOAT,                          
    rating BIGINT,                        
    image VARCHAR(255),                   
    category_id BIGINT,                  
    created_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3), 
    updated_at TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP(3),  
    deleted_at TIMESTAMP(3),
    CONSTRAINT fk_category
        FOREIGN KEY (category_id) 
        REFERENCES categories(id)
);