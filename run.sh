#!/bin/bash

if [ -f ./configs/.env ]
then
  export $(cat ./configs/.env | sed 's/#.*//g' | xargs)
fi


if [ -z "$PROJECT_NAME" ]; then
    echo "Você deve criar um arquivo .env e informar a variavel PROJECT_NAME"
    exit 1
fi

if [ -z "$DATABASE_PASSWORD" ]; then
    echo "Você deve criar um arquivo .env e informar a variavel DATABASE_PASSWORD"
    exit 1
fi

if [ -z "$DATABASE_USER" ]; then
    echo "Você deve criar um arquivo .env e informar a variavel DATABASE_USER"
    exit 1
fi

if [ -z "$DATABASE_URL" ]; then
    echo "Você deve criar um arquivo .env e informar a variavel DATABASE_URL"
    exit 1
fi


docker-compose down

docker-compose up -d --build

printf "\n\nACESSE A APLICACAO: http://localhost/\n"

printf "\n\nACESSE O PHPMYADMIN: http://localhost:3000\n"
printf "\nServidor: mysql"
printf "\nUtilizador: root"
printf "\nPalavra-passe: 12345\n\n"