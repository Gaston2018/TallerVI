create table roles_permisos(
rol int not null,
permisos int not null,
constraint rol_permisos_pk primary key (rol, permisos))