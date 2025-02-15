# Usa a imagem oficial do Go
FROM golang:1.20

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia os arquivos para o container
COPY . .

# Baixa as dependências
RUN go mod tidy

# Compila a aplicação
RUN go build -o server

# Expõe a porta correta
EXPOSE 3000

# Comando para rodar o app
CMD [ "./server" ]
