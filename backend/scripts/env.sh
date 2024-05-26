#!/bin/bash

# Defina a variável ENV como 'dev' se não estiver definida
: "${ENV:=development}"

# Determine o nome do arquivo .env a ser carregado
ENV_FILE=".env.${ENV}"

# Verifique se o arquivo .env correspondente existe
if [ ! -f "$ENV_FILE" ]; then
  echo "Arquivo $ENV_FILE não encontrado!"
  return 1
fi

# Carrega as variáveis do arquivo .env correspondente e exporta-as
while IFS='=' read -r key value; do
  # Ignora linhas em branco e comentários
  if [[ ! "$key" =~ ^#.*$ ]] && [[ -n "$key" ]]; then
    export "$key"="$value"
  fi
done < "$ENV_FILE"

echo "Variáveis de ambiente carregadas do arquivo $ENV_FILE"