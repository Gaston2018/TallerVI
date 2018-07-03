CREATE TABLE `tallervi`.`profesionales` (
  `matricula` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(30) NOT NULL,
  `apellido` VARCHAR(30) NOT NULL,
  `especialidad` VARCHAR(30) NOT NULL,
  `direccion` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`matricula`));