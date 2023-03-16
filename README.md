# Curso [Go: desenvolvendo uma API Rest](https://cursos.alura.com.br/course/go-desenvolvendo-api-rest)  

## Links
[Código original](https://github.com/alura-cursos/api-go-rest)  
[gorilla/mux](https://github.com/gorilla/mux)  
[gorilla/handlers](https://github.com/gorilla/handlers)  
[GORM ORM](https://gorm.io/)  
[Conectar no Postgres usando GORM](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)  

## Comandos  
```
Rodar na pasta do projeto api-go-rest:
go mod init github.com/brunosantanati/api-go-rest
go get -u github.com/gorilla/mux
go get github.com/gorilla/handlers
go get -u gorm.io/gorm
go get gorm.io/driver/postgres
docker-compose up
docker-compose down

docker-compose exec postgres sh
depois do comando acima:
hostname -i
OU
docker inspect <id-container> | grep IPAddress
Exemplo:
docker inspect 2328 | grep IPAddress

Rodar na pasta do projeto frontend-react-personalidades:
sudo apt install npm
npm install
npm update
npm start
```

## Links da aplicação
```
API Go backend:
http://localhost:8000

Aplicação frontend:
http://localhost:3000

pgadmin:
http://localhost:54321
```

## SQL
```
create table personalidades(
    id serial primary key,
    nome varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia) VALUES
('Deodato Petit Wertheimer', 'Deodato Petit Wertheimer foi um médico e político brasileiro, seus primeiros anos de vida foram em São Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Colégio Anchieta dos jesuítas.'),
('Carmela Dutra', 'Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16.º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara.');
```
