create table roles(
id_rol serial,
descripcion varchar(30),
primary key (id_rol)
)

INSERT INTO ROLES (DESCRIPCION) VALUES ('ADMINISTRADOR')
INSERT INTO ROLES (DESCRIPCION) VALUES ('OPERADOR')
INSERT INTO ROLES (DESCRIPCION) VALUES ('VETERINARIO')
INSERT INTO ROLES (DESCRIPCION) VALUES ('MASCOTA')
INSERT INTO ROLES (DESCRIPCION) VALUES ('DUEÑO')

create table usuarios (
id_usuario serial,
nombre varchar(50) NOT NULL,
apellido varchar(50) not null,
nombre_usuario varchar(50) UNIQUE not null,
password varchar(50) not null,
id_rol smallint not null,
primary key (id_usuario),
foreign key (id_rol) referenceS roles(id_rol)
)

INSERT INTO usuarios(NOMBRE,APELLIDO,NOMBRE_USUARIO,password,ID_ROL) 
VALUES ('ADMINISTRADOR','DE APLICACION','ADMIN','ADMIN',1)
VALUES ('OPERADOR','VETERINARIA','OPERADOR','OPERADOR',2)
VALUES ('JUAN','GONZALEZ','JGONZALEZ','1234',3)
VALUES ('RICARDO','SANCHEZ','RSANCHEZ','C137',3)

create table clientes (
id_cliente serial,
nombre varchar(50) NOT NULL,
apellido varchar(50) not null,
direccion varchar(50) UNIQUE not null,
telefono varchar(50) not null,
id_rol smallint not null,
primary key (id_cliente),
foreign key (id_rol) referenceS roles(id_rol)
)

create table mascotas (
id_mascota serial,
nombre varchar(50) NOT NULL,
apellido varchar(50) not null,
nombre_usuario varchar(50) UNIQUE not null,
password varchar(50) not null,
id_rol smallint not null,
id_cliente smallint not null,
primary key (id_mascota),
foreign key (id_rol) referenceS roles(id_rol),
foreign key (id_cliente) referenceS clientes(id_cliente)
)

create table turnos (
id_turno serial,
fecha date not null,
hora time not null,
id_usuario smallint not null,
id_cliente smallint not null,
id_mascota smallint not null,
primary key (id_turno),
foreign key (id_usuario) referenceS usuarios(id_usuario),
foreign key (id_cliente) referenceS clientes(id_cliente),
foreign key (id_mascota) referenceS mascotas(id_mascota)
)