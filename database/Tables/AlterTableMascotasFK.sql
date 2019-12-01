alter table mascotas add constraint especies_fk foreign key (especies) references especies(especies);
alter table mascotas add constraint razas_fk foreign key (razas) references razas(razas);
alter table mascotas add constraint mascotas_personas_fk foreign key (personas) references personas(personas);
alter table mascotas add constraint historiasclinicas_fk foreign key (historiaclinicas) references historiasclinicas(historiasclinicas)