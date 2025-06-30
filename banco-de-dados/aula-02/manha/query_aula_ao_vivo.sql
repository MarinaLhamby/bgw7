
-- São necessárias as seguintes consultas:
-- Selecciona o nome, a posição e a localização dos departamentos onde os vendedores trabalham.
SELECT f.nome, f.posto, d.localidad 
from funcionario f
inner join departamento d on f.depto_nro = d.depto_nro

-- Mostra os departamentos com mais de cinco empregados.
SELECT * 
FROM departamento d
INNER JOIN funcionario f on d.depto_nro = f.depto_nro
WHERE (SELECT COUNT(DEPTO_NRO) FROM departamento GROUP BY depto_nro )>5

-- Mostra o nome, o salário e o nome do departamento dos empregados que têm a mesma posição que "Mito Barchuk".
SELECT f.nome, f.salario, d.nombre_depto 
from funcionario f
inner join departamento d on f.depto_nro = d.depto_nro
where d.posto = (SELECT posto from funcionario WHERE nome = 'Mito' AND sobrenome = 'Barchuk')

-- Mostra os detalhes dos empregados que trabalham no departamento de contabilidade, ordenados por nome.
SELECT f.cod_emp, f.nome, f.sobrenome, f.posto, f.data_alta, f.salario, f.comissao, f.depto_nro
from funcionario f
inner join departamento d on f.depto_nro = d.depto_nro
where d.nombre_depto = 'Contabilidade'
order by f.nome

-- Mostra o nome do empregado com o salário mais baixo.
SELECT f.nome
from funcionario f
where f.salario = (SELECT MIN(salario) from funcionario)

-- Mostra os detalhes do empregado com o salário mais alto no departamento de "Vendas".
SELECT f.cod_emp, f.nome, f.sobrenome, f.posto, f.data_alta, f.salario, f.comissao, f.depto_nro
FROM funcionario f
INNER JOIN departamento d ON f.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Vendas'
AND f.salario = (
    SELECT MAX(f2.salario)
    FROM funcionario f2
    INNER JOIN departamento d2 ON f2.depto_nro = d2.depto_nro
    WHERE d2.nombre_depto = 'Vendas'
  );