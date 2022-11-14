insert into usuarios (nome, nick, email, senha)
values
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$zbHUJbE/xogiWXHiUn3ry.Xcrb3W9b2MazwnnOn0dQtU5ejf5P4e."), -- usuario 1
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$zbHUJbE/xogiWXHiUn3ry.Xcrb3W9b2MazwnnOn0dQtU5ejf5P4e."), -- usuario 2
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$zbHUJbE/xogiWXHiUn3ry.Xcrb3W9b2MazwnnOn0dQtU5ejf5P4e."); -- usuario 3

insert into seguidores(usuario_id, seguidor_id)
values
(2,1),
(2,3),
(1,2),
(3,1);

insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicacao do Usuario 1", "Essa e a publicação 1", 1),
("Publicacao do Usuario 2", "Essa e a publicação 2", 2),
("Publicacao do Usuario 3", "Essa e a publicação 3", 3);

select * from usuarios; select * from seguidores; select * from publicacoes;