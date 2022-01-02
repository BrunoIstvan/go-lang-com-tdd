# API Go Rest

## Criando o projeto

Para criar o projeto, abra o terminal na pasta onde estarão os fontes e execute o seguinte comando:

    go mod init github.com/BrunoIstvan/go-rest-api


Instalando a dependência gorilla mux

    go get -u github.com/gorilla/mux

Instalando a dependência para tratar do CORS

    go get github.com/gorilla/handlers

Instalando a dependência gorm

    go get -u gorm.io/gorm

Instalando a dependência contendo o driver do postgres para o Go

    go get gorm.io/driver/postgres

## Executando a aplicação

Para rodar a aplicação, executar o seguinte comando:

    go run main.go

## Criando banco de dados

Subindo a imagem do Docker contendo um banco de dados postgres:

    docker-compose up

Para pegar o IP do servidor:

    docker-compose exec postgres sh
    hostname -i

Para rodar o pgAdmin, em um novo browser, digite:

    http://localhost:54321

Para logar, colocar o email e senha especificado no serviço pgadmin-compose dentro do arquivo docker-compose.yml


