create table turnos(
turnos int not null,
fecha smalldatetime not null,
descripcion varchar(max),
constraint turnos_pk primary key (turnos))