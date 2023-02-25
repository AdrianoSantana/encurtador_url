Projeto criado como resolução ao desafio proposto na plataforma dev gym.

Link do desafio:
https://app.devgym.com.br/challenges/3ecd0771-981d-44dc-9eee-5ec69791a745

Descrição:

Nesse desafio você deve criar um servidor que encurte urls e faça redirecionamento.

Requisitos:


* Criar um servidor http que contenha dois endpoints:
* POST / - recebe uma url e retorna um código único
* GET /:code - utiliza o code para redirecionar para a url original
* O code é um código único, a mesma url enviada várias vezes gera códigos diferentes
* O code tem o tamanho de até 6 caracteres
