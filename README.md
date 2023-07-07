# Consideraciones referidas a la Base de Datos

Instalar MySQL: brew install mysql Opción2: -arm64 brew install mysql

Verificar estado de MySQL: Ejecutar el comando 'make build-database' desde el root del proyecto (comando declarado en el
archivo Makefile)

Correr el comando de creación de la BD: Chequear con 'mysql.server' status para verificar que el servicio se encuentra
inicializado. Caso contrario ejecutar comando 'mysql.server start'

EN NINGÚN CASO DEBERÁN PONER PASSWORD

# Considerações relacionadas ao Banco de Dados

Instale o MySQL: brew install mysql Option2: -arm64 brew install mysql

Verifique o status do MySQL: execute o comando 'make build-database' da raiz do projeto (comando declarado no Makefile)

Execute o comando de criação do banco de dados: Verifique com o status 'mysql.server' para verificar se o serviço foi
inicializado. Caso contrário, execute o comando 'mysql.server start'

EM NENHUM CASO VOCÊ DEVE COLOCAR UMA SENHA

# Para documentar com o Swagger

Rode os seguintes comandos pelo terminal na pasta principal do projeto, separadamente:

go install github.com/swaggo/swag/cmd/swag@latest

go mod tidy

export PATH=$PATH:$HOME/go/bin

# Após fazer suas alterações de documentação, rode sempre o comando abaixo para atualizá-las no projeto

`swag init -g cmd/main.go`

# Após fazer alterações em interfaces, executar o comando para gerar mocks:

`mockery`

* é necessário ter o mockery instalado