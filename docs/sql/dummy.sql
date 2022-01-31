CREATE DATABASE myrepo;

CREATE TABLE IF NOT EXISTS myrepo.dummy(
	id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    adress VARCHAR(255),
    age INT NOT NULL,
    type ENUM('Undead', 'Human', 'Mutant') DEFAULT 'Human',
    is_active TINYINT(1) DEFAULT '0',
    PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;