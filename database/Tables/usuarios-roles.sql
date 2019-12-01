create table usuarios_roles(
usuarios int not null,
roles int not null,
constraint usuarios_roles_pk primary key (usuarios, roles))