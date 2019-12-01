create table turnos_mascotas(
turnos int not null,
mascotas int not null,
constraint turnos_mascotas_pk primary key (turnos, mascotas))