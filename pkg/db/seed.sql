-- Seed data for Auto Parts Inventory Management System
-- Insert sample categories
INSERT INTO
    categories (category_name, description, parent_category_id)
VALUES
    ('Engine', 'Engine components and parts', NULL),
    (
        'Transmission',
        'Transmission components and parts',
        NULL
    ),
    (
        'Brakes',
        'Brake system components and parts',
        NULL
    ),
    (
        'Suspension',
        'Suspension system components and parts',
        NULL
    ),
    (
        'Electrical',
        'Electrical system components and parts',
        NULL
    ),
    (
        'Exhaust',
        'Exhaust system components and parts',
        NULL
    ),
    ('Body', 'Body parts and accessories', NULL),
    (
        'Interior',
        'Interior components and accessories',
        NULL
    ),
    ('Engine Oil', 'Engine oils and lubricants', 1),
    (
        'Engine Filters',
        'Engine air, oil, and fuel filters',
        1
    ),
    (
        'Transmission Fluid',
        'Transmission fluids and lubricants',
        2
    ),
    ('Brake Pads', 'Brake pads and shoes', 3),
    (
        'Shock Absorbers',
        'Shock absorbers and struts',
        4
    ),
    ('Batteries', 'Car batteries', 5),
    (
        'Spark Plugs',
        'Spark plugs and ignition components',
        5
    ),
    ('Mufflers', 'Exhaust mufflers and resonators', 6),
    ('Bumpers', 'Front and rear bumpers', 7),
    ('Seat Covers', 'Seat covers and upholstery', 8);

-- Insert sample vehicle makes
INSERT INTO
    vehicle_makes (make_name, country)
VALUES
    ('Toyota', 'Japan'),
    ('Ford', 'USA'),
    ('BMW', 'Germany'),
    ('Honda', 'Japan'),
    ('Chevrolet', 'USA'),
    ('Volkswagen', 'Germany'),
    ('Nissan', 'Japan'),
    ('Mercedes-Benz', 'Germany');

-- Insert sample vehicle models
INSERT INTO
    vehicle_models (make_id, model_name)
VALUES
    (1, 'Corolla'),
    (1, 'Camry'),
    (2, 'F-150'),
    (2, 'Mustang'),
    (3, '3 Series'),
    (3, '5 Series'),
    (4, 'Civic'),
    (4, 'Accord'),
    (5, 'Silverado'),
    (5, 'Malibu'),
    (6, 'Golf'),
    (6, 'Passat'),
    (7, 'Altima'),
    (7, 'Maxima'),
    (8, 'C-Class'),
    (8, 'E-Class');

-- Insert sample vehicle submodels
INSERT INTO
    vehicle_submodels (
        model_id,
        submodel_name,
        year_from,
        year_to,
        engine_type,
        engine_displacement,
        fuel_type,
        transmission_type,
        body_type
    )
VALUES
    (
        1,
        'LE',
        2015,
        2020,
        'Inline-4',
        1.8,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        1,
        'SE',
        2015,
        2020,
        'Inline-4',
        1.8,
        'Gasoline',
        'Manual',
        'Sedan'
    ),
    (
        2,
        'XLE',
        2018,
        2023,
        'Inline-4',
        2.5,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        3,
        'XL',
        2010,
        2020,
        'V6',
        3.5,
        'Gasoline',
        'Automatic',
        'Truck'
    ),
    (
        4,
        'GT',
        2015,
        2023,
        'V8',
        5.0,
        'Gasoline',
        'Manual',
        'Coupe'
    ),
    (
        5,
        '330i',
        2019,
        2023,
        'Inline-4',
        2.0,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        6,
        '530i',
        2020,
        2023,
        'Inline-4',
        2.0,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        7,
        'LX',
        2016,
        2021,
        'Inline-4',
        1.5,
        'Gasoline',
        'CVT',
        'Sedan'
    ),
    (
        8,
        'Touring',
        2018,
        2023,
        'Inline-4',
        1.5,
        'Hybrid',
        'CVT',
        'Sedan'
    ),
    (
        9,
        'LTZ',
        2015,
        2020,
        'V8',
        5.3,
        'Gasoline',
        'Automatic',
        'Truck'
    ),
    (
        10,
        'Premier',
        2016,
        2021,
        'Inline-4',
        2.0,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        11,
        'GTI',
        2017,
        2023,
        'Inline-4',
        2.0,
        'Gasoline',
        'Manual',
        'Hatchback'
    ),
    (
        12,
        'SE',
        2018,
        2023,
        'Inline-4',
        1.8,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        13,
        'SV',
        2015,
        2020,
        'Inline-4',
        2.5,
        'Gasoline',
        'CVT',
        'Sedan'
    ),
    (
        14,
        'SL',
        2016,
        2021,
        'V6',
        3.5,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        15,
        'AMG',
        2019,
        2023,
        'V8',
        4.0,
        'Gasoline',
        'Automatic',
        'Sedan'
    ),
    (
        16,
        'AMG',
        2020,
        2023,
        'Inline-6',
        3.0,
        'Gasoline',
        'Automatic',
        'Sedan'
    );

-- Insert sample suppliers
INSERT INTO
    suppliers (
        name,
        contact_person,
        phone,
        email,
        address,
        tax_id,
        payment_terms,
        notes
    )
VALUES
    (
        'AutoParts Inc.',
        'John Doe',
        '123-456-7890',
        'john@autoparts.com',
        '123 Main St, Anytown, USA',
        '123456789',
        'Net 30',
        'Primary supplier for engine parts'
    ),
    (
        'Brake Masters',
        'Jane Smith',
        '987-654-3210',
        'jane@brakemasters.com',
        '456 Elm St, Othertown, USA',
        '987654321',
        'Net 60',
        'Specializes in brake components'
    ),
    (
        'Suspension Pros',
        'Mike Johnson',
        '555-123-4567',
        'mike@suspensionpros.com',
        '789 Oak St, Sometown, USA',
        '555123456',
        'Net 45',
        'High-quality suspension parts'
    ),
    (
        'Electro Auto',
        'Sarah Lee',
        '444-555-6666',
        'sarah@electroauto.com',
        '321 Pine St, Newtown, USA',
        '444555666',
        'Net 30',
        'Electrical components supplier'
    ),
    (
        'Exhaust World',
        'Chris Brown',
        '777-888-9999',
        'chris@exhaustworld.com',
        '654 Cedar St, Oldtown, USA',
        '777888999',
        'Net 30',
        'Exhaust systems and parts'
    );

-- Insert sample items (auto parts)
INSERT INTO
    items (
        item_name,
        part_number,
        description,
        category_id,
        buy_price,
        sell_price,
        current_stock,
        minimum_stock,
        barcode,
        supplier_id,
        location_aisle,
        location_shelf,
        location_bin,
        weight_kg,
        dimensions_cm,
        warranty_period,
        image_url,
        is_active,
        notes
    )
VALUES
    (
        'Engine Oil 5W-30',
        'EOIL-5W30',
        'Synthetic engine oil 5W-30',
        9,
        25.00,
        35.00,
        100,
        20,
        '123456789012',
        1,
        'A1',
        'S1',
        'B1',
        5.0,
        '10x10x20',
        '1 year',
        'http://example.com/eoil5w30.jpg',
        TRUE,
        'High-quality synthetic oil'
    ),
    (
        'Air Filter',
        'AFIL-123',
        'Engine air filter',
        10,
        10.00,
        15.00,
        50,
        10,
        '234567890123',
        1,
        'A2',
        'S2',
        'B2',
        0.5,
        '15x15x5',
        '6 months',
        'http://example.com/afilter123.jpg',
        TRUE,
        'Standard air filter'
    ),
    (
        'Brake Pads Set',
        'BPAD-456',
        'Front brake pads set',
        12,
        30.00,
        50.00,
        30,
        5,
        '345678901234',
        2,
        'B1',
        'S1',
        'B1',
        2.0,
        '20x10x5',
        '1 year',
        'http://example.com/bpad456.jpg',
        TRUE,
        'Ceramic brake pads'
    ),
    (
        'Shock Absorber',
        'SHOCK-789',
        'Front shock absorber',
        13,
        50.00,
        80.00,
        20,
        5,
        '456789012345',
        3,
        'C1',
        'S1',
        'B1',
        5.5,
        '30x10x10',
        '2 years',
        'http://example.com/shock789.jpg',
        TRUE,
        'Heavy-duty shock absorber'
    ),
    (
        'Car Battery',
        'BATT-101',
        '12V car battery',
        14,
        100.00,
        150.00,
        15,
        5,
        '567890123456',
        4,
        'D1',
        'S1',
        'B1',
        15.0,
        '30x20x20',
        '3 years',
        'http://example.com/batt101.jpg',
        TRUE,
        'Maintenance-free battery'
    ),
    (
        'Spark Plug',
        'SPARK-202',
        'Iridium spark plug',
        15,
        5.00,
        10.00,
        200,
        50,
        '678901234567',
        4,
        'E1',
        'S1',
        'B1',
        0.1,
        '5x5x5',
        '1 year',
        'http://example.com/spark202.jpg',
        TRUE,
        'High-performance spark plug'
    ),
    (
        'Muffler',
        'MUFF-303',
        'Stainless steel muffler',
        16,
        80.00,
        120.00,
        10,
        2,
        '789012345678',
        5,
        'F1',
        'S1',
        'B1',
        8.0,
        '50x20x20',
        '2 years',
        'http://example.com/muff303.jpg',
        TRUE,
        'Durable muffler'
    ),
    (
        'Front Bumper',
        'BUMP-404',
        'Front bumper for sedan',
        17,
        150.00,
        250.00,
        5,
        1,
        '890123456789',
        1,
        'G1',
        'S1',
        'B1',
        10.0,
        '150x50x10',
        '1 year',
        'http://example.com/bump404.jpg',
        TRUE,
        'OEM-style bumper'
    ),
    (
        'Seat Cover Set',
        'SEAT-505',
        'Leather seat cover set',
        18,
        60.00,
        100.00,
        20,
        5,
        '901234567890',
        1,
        'H1',
        'S1',
        'B1',
        3.0,
        '50x50x10',
        '6 months',
        'http://example.com/seat505.jpg',
        TRUE,
        'Premium leather seat covers'
    );

-- Insert sample compatibility mappings
INSERT INTO
    compatibility (item_id, submodel_id, notes)
VALUES
    (
        1,
        1,
        'Compatible with Toyota Corolla LE 2015-2020'
    ),
    (
        1,
        2,
        'Compatible with Toyota Corolla SE 2015-2020'
    ),
    (
        2,
        1,
        'Compatible with Toyota Corolla LE 2015-2020'
    ),
    (
        2,
        2,
        'Compatible with Toyota Corolla SE 2015-2020'
    ),
    (3, 4, 'Compatible with Ford F-150 XL 2010-2020'),
    (4, 5, 'Compatible with Ford Mustang GT 2015-2023'),
    (5, 6, 'Compatible with BMW 330i 2019-2023'),
    (6, 7, 'Compatible with BMW 530i 2020-2023'),
    (7, 8, 'Compatible with Honda Civic LX 2016-2021'),
    (
        8,
        9,
        'Compatible with Honda Accord Touring 2018-2023'
    ),
    (
        9,
        10,
        'Compatible with Chevrolet Silverado LTZ 2015-2020'
    );

-- Insert sample purchases
INSERT INTO
    purchases (
        supplier_id,
        item_id,
        quantity,
        cost_per_unit,
        total_cost,
        invoice_number,
        received_by,
        notes
    )
VALUES
    (
        1,
        1,
        50,
        20.00,
        1000.00,
        'INV-1001',
        'John Doe',
        'Initial stock purchase'
    ),
    (
        1,
        2,
        100,
        8.00,
        800.00,
        'INV-1002',
        'John Doe',
        'Bulk purchase for air filters'
    ),
    (
        2,
        3,
        30,
        25.00,
        750.00,
        'INV-1003',
        'Jane Smith',
        'Brake pads for inventory'
    ),
    (
        3,
        4,
        20,
        40.00,
        800.00,
        'INV-1004',
        'Mike Johnson',
        'Shock absorbers for stock'
    ),
    (
        4,
        5,
        15,
        90.00,
        1350.00,
        'INV-1005',
        'Sarah Lee',
        'Car batteries for inventory'
    ),
    (
        4,
        6,
        200,
        4.00,
        800.00,
        'INV-1006',
        'Sarah Lee',
        'Spark plugs bulk purchase'
    ),
    (
        5,
        7,
        10,
        70.00,
        700.00,
        'INV-1007',
        'Chris Brown',
        'Mufflers for stock'
    ),
    (
        1,
        8,
        5,
        120.00,
        600.00,
        'INV-1008',
        'John Doe',
        'Front bumpers for inventory'
    ),
    (
        1,
        9,
        20,
        50.00,
        1000.00,
        'INV-1009',
        'John Doe',
        'Seat covers for stock'
    );

-- Insert sample sales
INSERT INTO
    sales (
        item_id,
        quantity,
        price_per_unit,
        total_price,
        transaction_number,
        customer_name,
        customer_phone,
        customer_email,
        sold_by,
        notes
    )
VALUES
    (
        1,
        10,
        35.00,
        350.00,
        'TXN-1001',
        'Alice Johnson',
        '111-222-3333',
        'alice@example.com',
        'John Doe',
        'Customer purchase'
    ),
    (
        2,
        5,
        15.00,
        75.00,
        'TXN-1002',
        'Bob Smith',
        '222-333-4444',
        'bob@example.com',
        'Jane Smith',
        'Regular customer'
    ),
    (
        3,
        2,
        50.00,
        100.00,
        'TXN-1003',
        'Charlie Brown',
        '333-444-5555',
        'charlie@example.com',
        'Mike Johnson',
        'Brake pads replacement'
    ),
    (
        4,
        1,
        80.00,
        80.00,
        'TXN-1004',
        'David Lee',
        '444-555-6666',
        'david@example.com',
        'Sarah Lee',
        'Shock absorber replacement'
    ),
    (
        5,
        1,
        150.00,
        150.00,
        'TXN-1005',
        'Eve White',
        '555-666-7777',
        'eve@example.com',
        'Chris Brown',
        'Battery replacement'
    ),
    (
        6,
        10,
        10.00,
        100.00,
        'TXN-1006',
        'Frank Harris',
        '666-777-8888',
        'frank@example.com',
        'John Doe',
        'Spark plugs bulk purchase'
    ),
    (
        7,
        1,
        120.00,
        120.00,
        'TXN-1007',
        'Grace Taylor',
        '777-888-9999',
        'grace@example.com',
        'Jane Smith',
        'Muffler replacement'
    ),
    (
        8,
        1,
        250.00,
        250.00,
        'TXN-1008',
        'Henry Clark',
        '888-999-0000',
        'henry@example.com',
        'Mike Johnson',
        'Front bumper replacement'
    ),
    (
        9,
        2,
        100.00,
        200.00,
        'TXN-1009',
        'Ivy Lewis',
        '999-000-1111',
        'ivy@example.com',
        'Sarah Lee',
        'Seat covers purchase'
    );
