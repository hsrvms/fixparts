-- Auto Parts Inventory Management System Database Schema
-- MVP Version

-- Drop tables if they exist (for clean reinstallation)
DROP TABLE IF EXISTS sales CASCADE;
DROP TABLE IF EXISTS purchases CASCADE;
DROP TABLE IF EXISTS compatibility CASCADE;
DROP TABLE IF EXISTS items CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS vehicle_submodels CASCADE;
DROP TABLE IF EXISTS vehicle_models CASCADE;
DROP TABLE IF EXISTS vehicle_makes CASCADE;
DROP TABLE IF EXISTS suppliers CASCADE;

-- Create extension for UUID generation if needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create sequences for IDs
CREATE SEQUENCE IF NOT EXISTS category_id_seq;
CREATE SEQUENCE IF NOT EXISTS make_id_seq;
CREATE SEQUENCE IF NOT EXISTS model_id_seq;
CREATE SEQUENCE IF NOT EXISTS submodel_id_seq;
CREATE SEQUENCE IF NOT EXISTS item_id_seq;
CREATE SEQUENCE IF NOT EXISTS supplier_id_seq;
CREATE SEQUENCE IF NOT EXISTS purchase_id_seq;
CREATE SEQUENCE IF NOT EXISTS sale_id_seq;

-- Categories table with hierarchical structure
CREATE TABLE categories (
    category_id INTEGER PRIMARY KEY DEFAULT nextval('category_id_seq'),
    category_name VARCHAR(100) NOT NULL,
    description TEXT,
    parent_category_id INTEGER REFERENCES categories(category_id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_category_name UNIQUE (category_name)
);

-- Vehicle Makes
CREATE TABLE vehicle_makes (
    make_id INTEGER PRIMARY KEY DEFAULT nextval('make_id_seq'),
    make_name VARCHAR(100) NOT NULL,
    country VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_make_name UNIQUE (make_name)
);

-- Vehicle Models (Parent models like A3, 418, etc.)
CREATE TABLE vehicle_models (
    model_id INTEGER PRIMARY KEY DEFAULT nextval('model_id_seq'),
    make_id INTEGER NOT NULL REFERENCES vehicle_makes(make_id) ON DELETE CASCADE,
    model_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_make_model UNIQUE (make_id, model_name)
);

-- Vehicle Submodels (like A3 Cabrio, 418d, etc.)
CREATE SEQUENCE IF NOT EXISTS submodel_id_seq;
CREATE TABLE vehicle_submodels (
    submodel_id INTEGER PRIMARY KEY DEFAULT nextval('submodel_id_seq'),
    model_id INTEGER NOT NULL REFERENCES vehicle_models(model_id) ON DELETE CASCADE,
    submodel_name VARCHAR(100) NOT NULL,
    year_from INTEGER NOT NULL,
    year_to INTEGER,
    engine_type VARCHAR(100),
    engine_displacement DECIMAL(3,1),
    fuel_type VARCHAR(50),
    transmission_type VARCHAR(50),
    body_type VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_submodel UNIQUE (model_id, submodel_name, year_from)
);

-- Suppliers
CREATE TABLE suppliers (
    supplier_id INTEGER PRIMARY KEY DEFAULT nextval('supplier_id_seq'),
    name VARCHAR(200) NOT NULL,
    contact_person VARCHAR(200),
    phone VARCHAR(50),
    email VARCHAR(200),
    address TEXT,
    tax_id VARCHAR(100),
    payment_terms VARCHAR(100),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_supplier_name UNIQUE (name)
);

-- Items (Auto Parts)
CREATE TABLE items (
    item_id INTEGER PRIMARY KEY DEFAULT nextval('item_id_seq'),
    item_name VARCHAR(200) NOT NULL,
    part_number VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    category_id INTEGER REFERENCES categories(category_id) ON DELETE SET NULL,
    buy_price DECIMAL(10,2) NOT NULL,
    sell_price DECIMAL(10,2) NOT NULL,
    current_stock INTEGER NOT NULL DEFAULT 0,
    minimum_stock INTEGER NOT NULL DEFAULT 5,
    barcode VARCHAR(100) UNIQUE,
    supplier_id INTEGER REFERENCES suppliers(supplier_id) ON DELETE SET NULL,
    location_aisle VARCHAR(50),
    location_shelf VARCHAR(50),
    location_bin VARCHAR(50),
    weight_kg DECIMAL(10,3),
    dimensions_cm VARCHAR(50), -- Format: LxWxH
    warranty_period VARCHAR(50),
    image_url VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_part_number UNIQUE (part_number),
    CONSTRAINT unique_item_name UNIQUE (item_name),
    CONSTRAINT positive_buy_price CHECK (buy_price >= 0),
    CONSTRAINT positive_sell_price CHECK (sell_price >= 0),
    CONSTRAINT non_negative_stock CHECK (current_stock >= 0)
);

-- Compatibility mapping between parts and vehicle submodels
CREATE TABLE compatibility (
    compat_id SERIAL PRIMARY KEY,
    item_id INTEGER NOT NULL REFERENCES items(item_id) ON DELETE CASCADE,
    submodel_id INTEGER NOT NULL REFERENCES vehicle_submodels(submodel_id) ON DELETE CASCADE,
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_item_submodel UNIQUE (item_id, submodel_id)
);

-- Purchases
CREATE TABLE purchases (
    purchase_id INTEGER PRIMARY KEY DEFAULT nextval('purchase_id_seq'),
    date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    supplier_id INTEGER NOT NULL REFERENCES suppliers(supplier_id) ON DELETE RESTRICT,
    item_id INTEGER NOT NULL REFERENCES items(item_id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL,
    cost_per_unit DECIMAL(10,2) NOT NULL,
    total_cost DECIMAL(10,2) NOT NULL,
    invoice_number VARCHAR(100),
    received_by VARCHAR(100),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT positive_quantity CHECK (quantity > 0),
    CONSTRAINT positive_cost_per_unit CHECK (cost_per_unit >= 0),
    CONSTRAINT positive_total_cost CHECK (total_cost >= 0)
);

-- Sales
CREATE TABLE sales (
    sale_id INTEGER PRIMARY KEY DEFAULT nextval('sale_id_seq'),
    date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    item_id INTEGER NOT NULL REFERENCES items(item_id) ON DELETE RESTRICT,
    quantity INTEGER NOT NULL,
    price_per_unit DECIMAL(10,2) NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    transaction_number VARCHAR(100),
    customer_name VARCHAR(200),
    customer_phone VARCHAR(50),
    customer_email VARCHAR(200),
    sold_by VARCHAR(100),
    notes TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT positive_quantity CHECK (quantity > 0),
    CONSTRAINT positive_price_per_unit CHECK (price_per_unit >= 0),
    CONSTRAINT positive_total_price CHECK (total_price >= 0)
);

-- Create indexes for performance
CREATE INDEX idx_categories_parent ON categories(parent_category_id);
CREATE INDEX idx_vehicle_models_make ON vehicle_models(make_id);
CREATE INDEX idx_items_category ON items(category_id);
CREATE INDEX idx_items_supplier ON items(supplier_id);
CREATE INDEX idx_items_barcode ON items(barcode);
CREATE INDEX idx_compatibility_item ON compatibility(item_id);
CREATE INDEX idx_compatibility_submodel ON compatibility(submodel_id);
CREATE INDEX idx_purchases_supplier ON purchases(supplier_id);
CREATE INDEX idx_purchases_item ON purchases(item_id);
CREATE INDEX idx_purchases_date ON purchases(date);
CREATE INDEX idx_sales_item ON sales(item_id);
CREATE INDEX idx_sales_date ON sales(date);

-- Create triggers for updated_at timestamp
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_categories_timestamp
BEFORE UPDATE ON categories
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_vehicle_makes_timestamp
BEFORE UPDATE ON vehicle_makes
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_vehicle_models_timestamp
BEFORE UPDATE ON vehicle_models
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_vehicle_submodels_timestamp
BEFORE UPDATE ON vehicle_submodels
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_suppliers_timestamp
BEFORE UPDATE ON suppliers
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_items_timestamp
BEFORE UPDATE ON items
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_purchases_timestamp
BEFORE UPDATE ON purchases
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

CREATE TRIGGER update_sales_timestamp
BEFORE UPDATE ON sales
FOR EACH ROW EXECUTE PROCEDURE update_timestamp();

-- Create a trigger to update inventory on purchase
CREATE OR REPLACE FUNCTION update_inventory_on_purchase()
RETURNS TRIGGER AS $$
BEGIN
   UPDATE items
   SET current_stock = current_stock + NEW.quantity,
       updated_at = CURRENT_TIMESTAMP
   WHERE item_id = NEW.item_id;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_inventory_on_purchase
AFTER INSERT ON purchases
FOR EACH ROW EXECUTE PROCEDURE update_inventory_on_purchase();

-- Create a trigger to update inventory on sale
CREATE OR REPLACE FUNCTION update_inventory_on_sale()
RETURNS TRIGGER AS $$
BEGIN
   UPDATE items
   SET current_stock = current_stock - NEW.quantity,
       updated_at = CURRENT_TIMESTAMP
   WHERE item_id = NEW.item_id;
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_inventory_on_sale
AFTER INSERT ON sales
FOR EACH ROW EXECUTE PROCEDURE update_inventory_on_sale();

-- Create view for low stock alerts
CREATE OR REPLACE VIEW low_stock_items AS
SELECT
    i.item_id,
    i.part_number,
    i.description,
    i.current_stock,
    i.minimum_stock,
    c.category_name,
    s.name as supplier_name,
    s.phone as supplier_phone,
    s.email as supplier_email
FROM
    items i
JOIN
    categories c ON i.category_id = c.category_id
JOIN
    suppliers s ON i.supplier_id = s.supplier_id
WHERE
    i.current_stock <= i.minimum_stock
ORDER BY
    (i.current_stock::float / i.minimum_stock) ASC;

-- Create view for item sales velocity (items sold per day)
CREATE OR REPLACE VIEW item_sales_velocity AS
SELECT
    i.item_id,
    i.part_number,
    i.description,
    COUNT(s.sale_id) as total_sales,
    SUM(s.quantity) as total_quantity_sold,
    (CURRENT_DATE - DATE '2023-01-01') as days_since_jan1,
    ROUND(SUM(s.quantity)::numeric / ((CURRENT_DATE - DATE '2023-01-01')::numeric), 2) as daily_sales_rate,
    CASE
        WHEN i.current_stock > 0 AND (SUM(s.quantity)::numeric / ((CURRENT_DATE - DATE '2023-01-01')::numeric)) > 0
        THEN ROUND(i.current_stock / (SUM(s.quantity)::numeric / ((CURRENT_DATE - DATE '2023-01-01')::numeric)))
        ELSE NULL
    END as estimated_days_until_stockout
FROM
    items i
LEFT JOIN
    sales s ON i.item_id = s.item_id
WHERE
    s.date >= '2023-01-01'
GROUP BY
    i.item_id, i.part_number, i.description, i.current_stock
ORDER BY
    daily_sales_rate DESC;

-- Create view for top selling items
CREATE OR REPLACE VIEW top_selling_items AS
SELECT
    i.item_id,
    i.part_number,
    i.description,
    c.category_name,
    COUNT(s.sale_id) as number_of_sales,
    SUM(s.quantity) as total_quantity_sold,
    SUM(s.total_price) as total_revenue,
    SUM(s.total_price) - (SUM(s.quantity) * i.buy_price) as estimated_profit
FROM
    items i
JOIN
    sales s ON i.item_id = s.item_id
JOIN
    categories c ON i.category_id = c.category_id
GROUP BY
    i.item_id, i.part_number, i.description, i.buy_price, c.category_name
ORDER BY
    total_revenue DESC;

-- Create view for vehicle compatibility count
CREATE OR REPLACE VIEW part_compatibility_summary AS
SELECT
    i.item_id,
    i.part_number,
    i.description,
    COUNT(c.compat_id) as compatible_vehicle_count,
    COUNT(DISTINCT vmod.model_id) as compatible_model_count,
    STRING_AGG(DISTINCT vm.make_name, ', ') as compatible_makes
FROM
    items i
LEFT JOIN
    compatibility c ON i.item_id = c.item_id
LEFT JOIN
    vehicle_submodels vsub ON c.submodel_id = vsub.submodel_id
LEFT JOIN
    vehicle_models vmod ON vsub.model_id = vmod.model_id
LEFT JOIN
    vehicle_makes vm ON vmod.make_id = vm.make_id
GROUP BY
    i.item_id, i.part_number, i.description
ORDER BY
    compatible_vehicle_count DESC;

-- Create function to get compatible parts for a specific vehicle submodel
CREATE OR REPLACE FUNCTION get_compatible_parts(make_name TEXT, model_name TEXT, submodel_name TEXT, model_year INTEGER)
RETURNS TABLE (
    item_id INTEGER,
    part_number VARCHAR,
    description TEXT,
    category_name VARCHAR,
    sell_price DECIMAL,
    current_stock INTEGER,
    barcode VARCHAR
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        i.item_id,
        i.part_number,
        i.description,
        c.category_name,
        i.sell_price,
        i.current_stock,
        i.barcode
    FROM
        items i
    JOIN
        categories c ON i.category_id = c.category_id
    JOIN
        compatibility comp ON i.item_id = comp.item_id
    JOIN
        vehicle_submodels vsub ON comp.submodel_id = vsub.submodel_id
    JOIN
        vehicle_models vm ON vsub.model_id = vm.model_id
    JOIN
        vehicle_makes vma ON vm.make_id = vma.make_id
    WHERE
        vma.make_name = make_name
        AND vm.model_name = model_name
        AND vsub.submodel_name = submodel_name
        AND vsub.year_from <= model_year
        AND (vsub.year_to IS NULL OR vsub.year_to >= model_year)
    ORDER BY
        c.category_name, i.part_number;
END;
$$ LANGUAGE plpgsql;

-- Create function to get all compatible parts for a model (across all submodels)
CREATE OR REPLACE FUNCTION get_model_compatible_parts(make_name TEXT, model_name TEXT)
RETURNS TABLE (
    item_id INTEGER,
    part_number VARCHAR,
    description TEXT,
    category_name VARCHAR,
    sell_price DECIMAL,
    current_stock INTEGER,
    compatible_submodels TEXT
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        i.item_id,
        i.part_number,
        i.description,
        c.category_name,
        i.sell_price,
        i.current_stock,
        STRING_AGG(DISTINCT vsub.submodel_name, ', ') as compatible_submodels
    FROM
        items i
    JOIN
        categories c ON i.category_id = c.category_id
    JOIN
        compatibility comp ON i.item_id = comp.item_id
    JOIN
        vehicle_submodels vsub ON comp.submodel_id = vsub.submodel_id
    JOIN
        vehicle_models vm ON vsub.model_id = vm.model_id
    JOIN
        vehicle_makes vma ON vm.make_id = vma.make_id
    WHERE
        vma.make_name = make_name
        AND vm.model_name = model_name
    GROUP BY
        i.item_id, i.part_number, i.description, c.category_name, i.sell_price, i.current_stock
    ORDER BY
        c.category_name, i.part_number;
END;
$$ LANGUAGE plpgsql;

-- Function to get category prefix
CREATE OR REPLACE FUNCTION get_category_prefix(category_id INTEGER)
RETURNS TEXT AS $$
DECLARE
    category_name TEXT;
BEGIN
    SELECT UPPER(LEFT(category_name, 2))
    INTO category_name
    FROM categories
    WHERE category_id = $1;
    RETURN category_name;
END;
$$ LANGUAGE plpgsql;
