-- Exibir o título e o nome do gênero de todas as séries.
select s.title, g.name  from series s 
inner join genres g on s.genre_id = g.id 

-- Mostre o título dos episódios, o nome e o sobrenome dos atores que trabalham em cada episódio.

select e.title, a.first_name, a.last_name from episodes e
inner join actor_episode ae on e.id = ae.episode_id
inner join actors a on ae.actor_id  = a.id

-- Mostre o título de todas as séries e o número total de temporadas de cada série.

select s.title, count(s2.id) from series s 
inner join seasons s2 on s.id = s2.serie_id
group by s2.serie_id, s.title

-- Mostre o nome de todos os gêneros e o número total de filmes de cada gênero, desde que seja maior ou igual a 3.

select g.name, count(m.id) as numero_filmes from genres g 
inner join movies m on g.id = m.genre_id
group by g.name
having numero_filmes >=3

-- Mostre apenas o nome e o sobrenome dos atores que trabalharam em todos os filmes de Guerra nas Estrelas e não os repita.

select DISTINCT a.first_name, a.last_name from actors a 
inner join actor_movie am on a.id = am.actor_id
inner join movies m on am.movie_id = m.id
where m.title LIKE 'La Guerra de las galaxias%'