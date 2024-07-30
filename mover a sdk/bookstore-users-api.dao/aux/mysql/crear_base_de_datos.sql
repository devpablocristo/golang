
/* Borrar base de datos */
DROP DATABASE IF EXISTS users;

/*  Base de datos   */
CREATE DATABASE IF NOT EXISTS users
DEFAULT CHARACTER SET utf8
DEFAULT COLLATE utf8_spanish_ci;

/*  Privilegios usuario.    */
GRANT ALL PRIVILEGES ON users.* TO 'mysql_user'@'localhost';

/*  Refrescar todos los privilegios. siempre que haya un cambio de privilegios. */
FLUSH PRIVILEGES;