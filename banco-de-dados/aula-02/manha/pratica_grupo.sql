-- Enumera os dados dos autores.
SELECT a.idAutor, a.nombre, a.nacionalidad
FROM autor a

-- Indica o nome e a idade dos alunos
SELECT e.nombre, e.edad
FROM estudiante e

-- Que alunos pertencem ao curso de informática?
SELECT e.idLector, e.nombre, e.apellido, e.direccion, e.carrera, e.edad
FROM estudiante e
WHERE e.carrera = 'Informatica'

-- Quais são os autores de nacionalidade francesa ou italiana?
SELECT a.idAutor, a.nombre, a.nacionalidad
FROM autor a
WHERE a.nacionalidade = 'Francesa' OR a.nacionalidad = 'Italiana'

-- Quais os livros que não são da área da Internet?
SELECT l.idLibro, l.titulo, l.editorial, l.area
FROM libro l
WHERE l.area <> 'Internet'

-- Enumera os livros publicados pela Salamandra.
SELECT l.idLibro, l.titulo, l.editorial, l.area
FROM libro l
WHERE l.editorial = 'Salamandra'

-- Enumera os nomes dos alunos cuja idade é superior à média.
SELECT e.nombre
FROM estudiante e
WHERE e.edad > (SELECT AVR(edad) FROM estudiante)

-- Enumera os nomes dos alunos cujo apelido começa com a letra G.
SELECT e.nombre
FROM estudiante e
WHERE e.apellido LIKE 'G%';

-- Faz uma lista dos autores do livro "O Universo: Guia de Viagem". (Apenas os nomes devem ser indicados).
SELECT a.nombre
FROM autor a
INNER JOIN libroautor la ON a.idAutor = la.idAutor
INNER JOIN libro l ON la.idLibro = l.idLibro
WHERE l.titulo  = 'O Universo: Guia de Viagem'

-- Que livros emprestaste ao leitor "Filippo Galli"?
SELECT l.idLibro, l.titulo, l.editorial, l.area
FROM libro l
INNER JOIN prestamo p ON l.idLivro = p.idLibro
INNER JOIN estudiante e on p.idlector = e.idLetor
WHERE e.nombre = 'Filippo Galli'

-- Indica o nome do aluno mais novo.
SELECT e.idLector, e.nombre, e.apellido, e.direccion, e.carrera, e.edad
FROM estudiante e
WHERE e.edad = (SELECT MIN(edad) FROM estudiante)

-- Enumera os nomes dos alunos a quem foram emprestados livros da Base de Dados.
SELECT e.nombre
FROM estudiante e
INNER JOIN prestamo p ON e.idLector = p.idLector
INNER JOIN libro l ON p.idLibro = l.idLibro

-- Enumera os livros que pertencem à autora J.K. Rowling.
SELECT l.idLibro, l.titulo, l.editorial, l.area
FROM libro l
INNER JOIN libroautor la ON l.idlibro = la.idLibro
INNER JOIN autor a ON la.idAutor = a.idAutor
a.nombre = 'J.K. Rowling'

-- Enumera os títulos dos livros que deviam ser devolvidos em 16/07/2021.
SELECT l.titulo
FROM libro l
INNER JOIN prestamo p ON l.idLivro = p.idLibro
WHERE p.fechaDevolucion = '2021-07-16'

-- Criar tabelas
CREATE TABLE autor (
    idAutor INT PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    nacionalidad VARCHAR(50)
);

CREATE TABLE libro (
    idLibro INT PRIMARY KEY,
    titulo VARCHAR(200) NOT NULL,
    editorial VARCHAR(100),
    area VARCHAR(50)
);

CREATE TABLE libroautor (
    idAutor INT,
    idLibro INT,
    PRIMARY KEY (idAutor, idLibro),
    FOREIGN KEY (idAutor) REFERENCES autor(idAutor),
    FOREIGN KEY (idLibro) REFERENCES libro(idLibro)
);

CREATE TABLE estudiante (
    idLector INT PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    apellido VARCHAR(50) NOT NULL,
    direccion VARCHAR(200),
    carrera VARCHAR(50),
    edad INT
);

CREATE TABLE prestamo (
    idLector INT,
    idLibro INT,
    fechaPrestamo DATE NOT NULL,
    fechaDevolucion DATE,
    devuelto BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (idLector, idLibro, fechaPrestamo),
    FOREIGN KEY (idLector) REFERENCES estudiante(idLector),
    FOREIGN KEY (idLibro) REFERENCES libro(idLibro)
);

-- Inserts
INSERT INTO autor (idAutor, nombre, nacionalidad) VALUES 
(1, 'J.K. Rowling', 'Británica'),
(2, 'Gabriel García Márquez', 'Colombiana'),
(3, 'Mario Vargas Llosa', 'Peruana'),
(4, 'Albert Camus', 'Francesa'),
(5, 'Umberto Eco', 'Italiana');

INSERT INTO libro (idLibro, titulo, editorial, area) VALUES
(1, 'Harry Potter y la Piedra Filosofal', 'Salamandra', 'Fantasía'),
(2, 'Cien Años de Soledad', 'Random House', 'Literatura'),
(3, 'La Ciudad y los Perros', 'Alfaguara', 'Literatura'),
(4, 'El Extranjero', 'Alianza', 'Filosofía'),
(5, 'O Universo: Guia de Viagem', 'Nova Fronteira', 'Ciencia'),
(6, 'Bases de Datos Relacionales', 'Prentice Hall', 'Base de Datos'),
(7, 'El Nombre de la Rosa', 'Lumen', 'Novela Histórica');

INSERT INTO libroautor (idAutor, idLibro) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(5, 7);

INSERT INTO estudiante (idLector, nombre, apellido, direccion, carrera, edad) VALUES
(1, 'Juan', 'Pérez', 'Calle A, 123', 'Informática', 22),
(2, 'María', 'González', 'Avenida B, 456', 'Literatura', 24),
(3, 'Carlos', 'García', 'Calle C, 789', 'Informática', 19),
(4, 'Filippo', 'Galli', 'Via Roma, 10', 'Ingeniería', 25),
(5, 'Ana', 'Gutiérrez', 'Avenida D, 321', 'Medicina', 21);

INSERT INTO prestamo (idLector, idLibro, fechaPrestamo, fechaDevolucion, devuelto) VALUES
(1, 6, '2021-06-15', '2021-07-15', TRUE),
(2, 2, '2021-06-20', '2021-07-20', FALSE),
(3, 1, '2021-06-25', '2021-07-16', FALSE),
(4, 5, '2021-06-30', '2021-07-30', FALSE),
(4, 7, '2021-07-01', '2021-08-01', FALSE),
(5, 6, '2021-07-05', '2021-08-05', FALSE);