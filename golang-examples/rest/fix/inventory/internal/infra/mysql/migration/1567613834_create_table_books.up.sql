CREATE TABLE IF NOT EXISTS books (
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	code varchar(200) DEFAULT NULL,
	author varchar(200) DEFAULT NULL,
	title longtext,
	price bigint(20) DEFAULT NULL,
	stock bigint(20) DEFAULT NULL,
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	PRIMARY KEY (id),
	UNIQUE KEY code (code)
);