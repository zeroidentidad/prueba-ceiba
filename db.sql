USE pruebaceiba;

DROP TABLE IF EXISTS `pagos`;

CREATE TABLE `pagos` (
  `documentoIdentificacionArrendatario` int(11) NOT NULL,
  `codigoInmueble` varchar(10) NOT NULL,
  `valorPagado` int(11) NOT NULL,  
  `fechaPago` date NOT NULL,
  INDEX idx_pagos (`documentoIdentificacionArrendatario`, `codigoInmueble`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;