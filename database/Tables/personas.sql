create table personas(
personas int not null,
DNI int not null,
direccion varchar(max),
telefono int,
mail varchar(max),
nombre varchar(20) not null,
apellido varchar(30) not null,
fechaingreso smalldatetime not null,
constraint personas_pk primary key (personas))