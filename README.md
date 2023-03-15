# Curso [Go: desenvolvendo uma API Rest](https://cursos.alura.com.br/course/go-desenvolvendo-api-rest)  

## Links
[Código original](https://github.com/alura-cursos/api-go-rest)  
[gorilla/mux](https://github.com/gorilla/mux)  

## Comandos  
```
Rodar na pasta do projeto:
go mod init github.com/brunosantanati/api-go-rest 
go get -u github.com/gorilla/mux  
docker-compose up

docker-compose exec postgres sh
depois do comando acima:
hostname -i
OU
docker inspect <id-container> | grep IPAddress
Exemplo:
docker inspect 2328 | grep IPAddress
```

## Links da aplicação
```
aplicação:
http://localhost:8000

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