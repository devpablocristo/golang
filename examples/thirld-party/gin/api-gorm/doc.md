Para cargar datos:
1. Correr programa.
2. Usar la sentencia cURL.

Cuidado con las comillas!

No funciona bien:
curl -i -X POST http://localhost:8080/people -d ‘{ “FirstName”: “Elvis”, “LastName”: “Presley”}’

Funciona bien:
curl -d '{"FirstName":"lala", "LastName":"xxxxx"}' -H "Content-Type: application/json" -X POST http://localhost:8080/people

curl  -i -X POST http://localhost:8080/people -d '{"FirstName":"Elvis", "LastName":"Presley"}'
