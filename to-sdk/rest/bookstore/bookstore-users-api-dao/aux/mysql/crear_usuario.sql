/* 	Crear usuario con contraseña.
	'%' puede conectar desde cualquier lado.
*/
CREATE USER IF NOT EXISTS 'mysq_user'@'%' IDENTIFIED BY 'pass';

/* 
	Garantizar privilegios usuario.
	Nota: aqui estoy dandole todos los privilegios, seria conveniente restringirlos en el futuro.
*/
REVOKE ALL PRIVILEGES ON *.* FROM 'mysq_user'@'%t'; 
GRANT ALL PRIVILEGES ON *.* TO 'mysq_user'@'%' 
REQUIRE NONE WITH GRANT OPTION 
MAX_QUERIES_PER_HOUR 0 
MAX_CONNECTIONS_PER_HOUR 0 
MAX_UPDATES_PER_HOUR 0 
MAX_USER_CONNECTIONS 0;

/*  Refrescar todos los privilegios. siempre que haya un cambio de privilegios. */
FLUSH PRIVILEGES;