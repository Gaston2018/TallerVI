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