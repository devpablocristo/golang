# Descripción:
Imagen base para go-template.

# NOTAS:
## .env files en dockerfile: 
    Al parecer no se puede leed .env file desde docker

## Image tagging:
    docker hub ID/repo-project name version directorio:version
    docker push devpablocristo/go-gin-service:0.1-air-alpine


# HACER:
## RUN ["go", "install", "github.com/go-delve/delve/cmd/dlv@latest"]
agregar delve (debugger)