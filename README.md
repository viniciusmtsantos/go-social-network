<h1 align="center"> Aplicação de Rede Social em Go Lang </h1>

# Índice
* [Descrição do Projeto](#descrição-do-projeto)
* [Status do Projeto](#status-do-projeto)
* [Funcionalidades da Aplicação](#funcionalidades-da-aplicação)
* [Packages de Golang utilizados](#packages-de-golang-utilizados)

## Descrição do Projeto
Uma aplicação em Golang que desenvolve uma rede social desenvolvida para que a parte front-end faça uma requisiçao para uma API, sendo, por sua vez, responsável pela comunicação entre o banco de dados e a aplicação em si, devolvendo uma resposta.
## Status do Projeto
<p align="center">
<img src="https://img.shields.io/badge/status-Desenvolvido-green"/>
</p>

## Funcionalidades da Aplicação
- `Criar e Excluir usuário criado`: A aplicação permite que o usuário crie sua conta com um email, nick, nome e senha para realização do login. O usuário cadastrado, caso queira, pode excluir sua conta da rede social.
- `Seguir e Deixar de Seguir outro usuário`: A aplicação permite que o usuário siga outro(s) usuario(s) cadastrados na rede social. Caso queira, o usuário pode deixar de seguir outro(s) usuario(s).
- `Atualizar informações do usuario`: A aplicação permite que o usuário atualize suas informações como nome, senha, nick e email. As duas últimas só poderão ser alteradas desde que não estejam sendo utilizadas por outro usuário.
- `Criar e Excluir publicações`: A aplicação permite que o usuário publique um conteúdo para que os seus seguidores a vejam. Esta publicação contém um título e um conteúdo.
- `Atualizar publicações`: A aplicação permite que o usuário atualize qualquer publicação de sua autoria, editando o seu título e o seu conteúdo.
- `Curtir e Descurtir publicações`: A aplicação permite que o usuário curta ou deixe de curtir uma publicação de autoria própria ou de outrem.

## Packages de Golang utilizados

`package controllers:`

A função deste pacote é ser uma camada intermediária para receber as requisições
