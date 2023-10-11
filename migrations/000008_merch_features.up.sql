CREATE TABLE IF NOT EXISTS merch.category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS merch.product (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    category_id INT NOT NULL,
    primary_image_id INT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS merch.size (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    product_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS merch.color (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    product_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS merch.price (
    id SERIAL PRIMARY KEY,
    price VARCHAR(255) NOT NULL,
    product_id INT NOT NULL,
    size_id INT NOT NULL,
    color_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS merch.transaction (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT NOT NULL,
    coupon_id INT,
    proof_of_payment TEXT,
    status VARCHAR(255) NOT NULL,
    paid_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS merch.image (
    id SERIAL PRIMARY KEY,
    path TEXT NOT NULL,
    product_id INT NOT NULL REFERENCES merch.product(id),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS merch.coupon (
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    discount_type VARCHAR(255) NOT NULL,
    discount_value NUMERIC NOT NULL,
    max_use_count INT NOT NULL,
    used_count INT NOT NULL DEFAULT 0,
    valid_from TIMESTAMP,
    valid_until TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

ALTER TABLE merch.product
ADD CONSTRAINT product_category_id_fkey FOREIGN KEY (category_id) REFERENCES merch.category(id),
ADD CONSTRAINT product_primary_image_id_fkey FOREIGN KEY (primary_image_id) REFERENCES merch.image(id);

ALTER TABLE merch.size
ADD CONSTRAINT size_product_id_fkey FOREIGN KEY (product_id) REFERENCES merch.product(id);

ALTER TABLE merch.color
ADD CONSTRAINT color_product_id_fkey FOREIGN KEY (product_id) REFERENCES merch.product(id);

ALTER TABLE merch.price
ADD CONSTRAINT price_product_id_fkey FOREIGN KEY (product_id) REFERENCES merch.product(id),
ADD CONSTRAINT price_size_id_fkey FOREIGN KEY (size_id) REFERENCES merch.size(id),
ADD CONSTRAINT price_color_id_fkey FOREIGN KEY (color_id) REFERENCES merch.color(id);

ALTER TABLE merch.transaction
ADD CONSTRAINT transaction_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id),
ADD CONSTRAINT transaction_product_id_fkey FOREIGN KEY (product_id) REFERENCES merch.product(id),
ADD CONSTRAINT transaction_coupon_id_fkey FOREIGN KEY (coupon_id) REFERENCES merch.coupon(id);

ALTER TABLE merch.image
ADD CONSTRAINT merch_image_product_id_fkey FOREIGN KEY (product_id) REFERENCES merch.product(id);