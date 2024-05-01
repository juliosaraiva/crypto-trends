# Use a imagem oficial do Node.js 19
FROM node:19-alpine

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos necessários para o contêiner
COPY package.json yarn.lock ./

# Instale as dependências
RUN yarn install

# Copie o restante dos arquivos do aplicativo
COPY . .

# Construa o aplicativo Next.js para produção
RUN yarn build

# Exponha a porta 3000 para acessar o aplicativo
EXPOSE 3000

# Comando para iniciar o servidor Next.js
CMD ["yarn", "start"]
