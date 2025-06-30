-- selecione tudo na tabela de cliente
SELECT id, nome, sobrenome, provincia, cidade, data_nascimento
FROM cliente
-- selecione tudo na tabela de planos
SELECT id, velocidade, preco, desconto, id_cliente 
FROM plano
-- selecione cliente por provincia (Rio Grande do Sul)
SELECT id, nome, sobrenome, provincia, cidade, data_nascimento
FROM cliente
WHERE provincia = 'Rio Grande do Sul'
-- atualizar Rio Grande do Sul para RS
UPDATE cliente
SET provincia='RS'
WHERE provincia = 'Rio Grande do Sul'
-- selecionar planos com desconto
SELECT id, velocidade, preco, desconto, id_cliente 
FROM plano
WHERE desconto <> 0
-- atualizar valores dos planos pra 10% a mais
UPDATE plano
SET preco = preco * 1.10;
-- deletar plano
DELETE FROM plano
WHERE id_cliente = 1
-- descobrir numero de planos contratados por velocidade
SELECT COUNT(id), velocidade 
FROM plano 
GROUP BY velocidade
-- ver valor mais caro pago no plano
SELECT MAX(preco) 
FROM plano
-- ver valor do plano médio
SELECT AVG(preco)
FROM plano
GROUP BY velocidade
