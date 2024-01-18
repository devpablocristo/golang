CREATE TABLE IF NOT EXISTS items (
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
	code varchar(200) DEFAULT NULL,
	description varchar(200) DEFAULT NULL,
	title longtext,
	price bigint(20) DEFAULT NULL,
	stock bigint(20) DEFAULT NULL,
	status varchar(200) DEFAULT NULL,
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	PRIMARY KEY (id),
	UNIQUE KEY code (code)
);