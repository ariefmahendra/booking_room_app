<<<<<<< HEAD
create database booking_room_db;
=======
CREATE DATABASE booking_room_db;
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE role AS ENUM ('GA', 'ADMIN', 'EMPLOYEE');

CREATE TYPE facility_status AS ENUM ('REQUEST','AVAILABLE', 'BOOKED', 'BROKEN');

CREATE TYPE approval_status AS ENUM ('PENDING', 'ACCEPT', 'DECLINE');

CREATE TABLE mst_employee (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    division VARCHAR(100),
    position VARCHAR(100),
    role role,
    contact VARCHAR(15),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE mst_room (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    code_room VARCHAR(10) UNIQUE NOT NULL,
    room_type VARCHAR (50),
    capacity INTEGER,
    facilities VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);


CREATE TABLE mst_facilities (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    code_name VARCHAR(10) UNIQUE NOT NULL,
    facilities_type VARCHAR(255),
    status facility_status DEFAULT 'AVAILABLE',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE tx_room_reservation (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id uuid,
    room_id UUID,
    additional_id uuid,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    notes VARCHAR(255),
    approval_status approval_status DEFAULT 'PENDING',
    approval_note VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (employee_id) REFERENCES mst_employee(id),
    FOREIGN KEY (room_id) REFERENCES mst_room(id),
    FOREIGN KEY (additional_id) REFERENCES mst_facilities(id)
);

CREATE TABLE tx_additional (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    reservation_id uuid,
    facilities_id uuid,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (reservation_id) REFERENCES tx_room_reservation(id),
    FOREIGN KEY (facilities_id) REFERENCES mst_facilities(id)
<<<<<<< HEAD
);
=======
);
>>>>>>> ed7e6ada7c231957f8498b03fd926752a5f88f1d
