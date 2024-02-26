--Crear la base de datos
CREATE DATABASE factura;

-- Crear tabla CLIENTE
CREATE TABLE CLIENTE (
    id_cliente SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    telefono VARCHAR(15),
    identificacion VARCHAR(20) NOT NULL,
    correo VARCHAR(100) NOT NULL
);

-- Crear tabla FACTURA
CREATE TABLE FACTURA (
    id_factura SERIAL PRIMARY KEY,
    fecha DATE NOT NULL,
    descripcion VARCHAR(255) NOT NULL,
    id_cliente INTEGER REFERENCES CLIENTE(id_cliente) ON DELETE CASCADE
);

-- Crear tabla ITEM
CREATE TABLE ITEM (
    id_item SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    descripcion VARCHAR(255),
    valor NUMERIC(10, 2) NOT NULL
);

--Crear tabla de rompimiento
CREATE TABLE FACTURA_ITEM (
    id_factura_item SERIAL PRIMARY KEY,
    id_factura INTEGER REFERENCES FACTURA(id_factura) ON DELETE CASCADE,
    id_item INTEGER REFERENCES ITEM(id_item) ON DELETE CASCADE
);

-- Crear índices únicos para evitar duplicados en la relación muchos a muchos
CREATE UNIQUE INDEX idx_factura_item ON FACTURA_ITEM (id_factura, id_item);
