DROP TABLE IF EXISTS cliente;

CREATE TABLE CLIENTE (
    id_cliente SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    telefono VARCHAR(15),
    identificacion VARCHAR(20) NOT NULL,
    correo VARCHAR(100) NOT NULL
);
