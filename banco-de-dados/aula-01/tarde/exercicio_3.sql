-- É solicitado que você crie um novo banco de dados chamado "empresa_internet".
	CREATE DATABASE empresa_internet
-- Incorpore 10 registros na tabela de clientes e 5 na tabela de planos de Internet.
-- Faça as associações/relacionamentos correspondentes entre esses registros.
USE empresa_internet;
CREATE TABLE cliente(
	id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	nome VARCHAR(150) NOT NULL,
	sobrenome VARCHAR(300) NOT NULL,
	provincia VARCHAR(200) NOT NULL,
	cidade VARCHAR(200) NOT NULL,
	data_nascimento DATETIME NOT NULL
	
);

CREATE TABLE plano(
id iNT NOT NULL AUTO_INCREMENT PRIMARY KEY,
velocidade INT NOT NULL,
preco FLOAT NOT NULL,
desconto INT NOT NULL DEFAULT 0,
id_cliente INT NOT NULL,
CONSTRAINT `plano_id_cliente_fk` FOREIGN KEY (`id_cliente`) REFERENCES `cliente` (`id`)
);

INSERT INTO cliente (nome, sobrenome, provincia, cidade, data_nascimento) VALUES
('Ana', 'Silva', 'Minas Gerais', 'Belo Horizonte', '1990-05-12 00:00:00'),
('Carlos', 'Santos', 'São Paulo', 'Campinas', '1985-11-23 00:00:00'),
('Mariana', 'Oliveira', 'Rio de Janeiro', 'Niterói', '1992-07-30 00:00:00'),
('Pedro', 'Souza', 'Bahia', 'Salvador', '1988-03-15 00:00:00'),
('Juliana', 'Ferreira', 'Paraná', 'Curitiba', '1995-09-08 00:00:00'),
('Lucas', 'Almeida', 'Pernambuco', 'Recife', '1991-01-20 00:00:00'),
('Fernanda', 'Costa', 'Ceará', 'Fortaleza', '1987-06-14 00:00:00'),
('Rafael', 'Gomes', 'Rio Grande do Sul', 'Porto Alegre', '1993-12-03 00:00:00'),
('Patrícia', 'Martins', 'Santa Catarina', 'Florianópolis', '1996-04-27 00:00:00'),
('Bruno', 'Barbosa', 'Goiás', 'Goiânia', '1989-08-19');

INSERT INTO plano (velocidade, preco, desconto, id_cliente) VALUES
(100, 99.90, 10, 1),
(200, 149.90, 15, 2),
(300, 199.90, 20, 3),
(150, 119.90, 5, 4),
(500, 299.90, 25, 5);