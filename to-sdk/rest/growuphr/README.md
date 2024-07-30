## Numbers Manager API
------------------------------------------------------------
Solution for the proposed challenge.
It is written entirely in Golang and it uses an in-memory database.

To demonstrate horizontal scaling, I created 3 instances that run concurrently, configured on ports 8080, 8081, and 8082.

### Run the tests
```shell script
go test -v ./...
```

### Build the API
```shell script
go build -o num-man cmd/main.go
```

### Run the API
```shell script
./num-man
```

## With Docker compose

```shell script
make up
```
or
```shell script
docker-compose up
```

### down all the containers
```shell script
make down
```
or
```shell script
docker-compose down --remove-orphans
```

### check status
```shell script
docker-compose ps
```

### Check if API is up and running (server response: 200)
```shell script
curl -X GET \
  http://localhost:8080/heartbeat
```

### Reserve a number
```shell script
curl -X POST \
  http://localhost:8080/api/v1/number-manager/reserve/client1 \
  -d '{
  "number": 1
}'
```

```shell script
curl -X POST \
  http://localhost:8081/api/v1/number-manager/reserve/client2 \
  -d '{
  "number": 2
}'
```

```shell script
curl -X POST \
  http://localhost:8082/api/v1/number-manager/reserve/client3 \
  -d '{
  "number": 3
}'
```

### List numbers
```shell script
curl -X GET \
  localhost:8080/api/v1/number-manager/reserved-numbers
```
```shell script
curl -X GET \
  localhost:8081/api/v1/number-manager/reserved-numbers
```
```shell script
curl -X GET \
  localhost:8082/api/v1/number-manager/reserved-numbers
```