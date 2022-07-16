CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    picture VARCHAR(555),
    phone_number VARCHAR(255),
    address VARCHAR(255),
    is_active BOOLEAN,
    is_admin BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_roles PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS producers (
    id BIGSERIAL NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    document_number VARCHAR(255) UNIQUE,
    birth_date DATE,
    phone_number VARCHAR(255),
    address VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_producers PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS productions (
    id BIGSERIAL NOT NULL,
    producer BIGINT,
    lote_number VARCHAR(255),
    entry VARCHAR(255),
    name VARCHAR(255),
    production_type VARCHAR(255),
    area DECIMAL(10,2),
    cultivated_area DECIMAL(10,2),
    latitude DECIMAL(10,8) NOT NULL,
    longitude DECIMAL(11,8) NOT NULL,
    picture VARCHAR(255),
    cadastral_registration VARCHAR(255),
    district VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_productions PRIMARY KEY (id),
    CONSTRAINT fk_producers_productions FOREIGN KEY (producer)
    REFERENCES producers(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sections (
    id BIGSERIAL NOT NULL,
    section_number VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_sections PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS intakes (
    id BIGSERIAL NOT NULL,
    intake_number VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    section BIGINT,
    latitude DECIMAL(10,8) NOT NULL,
    longitude DECIMAL(11,8) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_intakes PRIMARY KEY (id),
    CONSTRAINT fk_intakes_sections FOREIGN KEY (section) 
    REFERENCES sections(id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS intakes_productions (
    intake_id BIGINT REFERENCES intakes(id)
    ON DELETE CASCADE,
    production_id BIGINT REFERENCES productions(id)
    ON DELETE CASCADE,
    watering_order INTEGER UNIQUE,
    CONSTRAINT pk_intakes_productions PRIMARY KEY (intake_id, production_id)
);

CREATE TABLE IF NOT EXISTS turns (
    id BIGSERIAL NOT NULL,
    start_date TIMESTAMP NOT NULL,
    turn_hours DECIMAL(10,2) NOT NULL,
    end_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT pk_turns PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS turns_productions (
    turn_id BIGINT REFERENCES turns(id)
    ON DELETE CASCADE,
    production_id BIGINT REFERENCES productions(id)
    ON DELETE CASCADE,
    CONSTRAINT pk_turns_productions PRIMARY KEY (turn_id, production_id)
);
