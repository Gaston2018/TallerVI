CREATE TABLE profesionales(
	nombre varchar(60) not null,
    apellido varchar(30) not null,
    matricula smallint,
    especialidad varchar(50) not null,
    
    primary key (matricula)
)