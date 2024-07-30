# Start

- Create .env file from .env.example
- Run run the project:

```bash
sudo make up
```

- Build or rebuild the project:

```bash
sudo make bu
```

- Stop the project:

```bash
sudo make down
```

## Endpoints

### Get all events

```bash
curl --request GET localhost:8080/api/v1/events/get-all
```

### Create event

```bash
curl --header "Content-Type: application/json" \
     --request POST \
     --data '{"start_time": "2023-05-16T23:30:00.000Z", "end_time": "2023-05-17T00:00:00.000Z"}' \
     http://localhost:8080/api/v1/events/create
```

### Get Overlaping events

```bash
curl --request GET localhost:8080/api/v1/events/get-overlaping
```
