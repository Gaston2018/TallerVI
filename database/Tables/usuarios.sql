create table usuarios(
usuarios int not null,
personas int not null,
mail varchar(40),
constrase�a varchar(100),
constraint usuarios_pk primary key (usuarios),
constraint personas_fk foreign key (personas) references personas(personas))