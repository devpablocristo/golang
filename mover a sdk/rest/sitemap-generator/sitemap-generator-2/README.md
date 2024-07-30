# Commands Examples:

```sh
-u http://localhost:8081/index.html -p 4 -of sitemap.xml -md 0
```

```sh
-u http://www.example.com -p 4 -of sitemap.xml -md 0
```

```sh
-u http://localhost:8081/index.html -p 4 -of ./file/sitemap.xml -md 2
```

```sh
-u http://localhost:8081/index.html -p 1 -of ./file/sitemap.xml -md 5
```

```sh
-u http://localhost:8081/index.html -p 2 -of ./file/sitemap.xml -md 4
```

```sh
-u http://localhost:8081/index.html -p 1 -of ./sitemap.xml -md 10
```

```sh
-u http://localhost:8081/index.html -p 0 -of ./file.xml -md 5
```

## Testing webserver 

For testing purposes I have created a web server in localhost:8081

it can be launched with the following command:

```sh
go run test-webserver/server.go
```