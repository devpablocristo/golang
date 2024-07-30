# Event Management Service

A simple event management service implemented using FastAPI.

## Setup

### Using Docker Compose

1. **Run Docker Compose**:
    ```bash
    docker-compose up
    ```

2. **Access the application**: Open your browser and go to `http://127.0.0.1:8000`.

### Additional Commands

- **Stop the containers**:
    ```bash
    docker-compose down
    ```

- **Rebuild the image and run the containers**:
    ```bash
    docker-compose up --build
    ```

### Running Locally

1. **Install dependencies**:
    ```bash
    pip install -r requirements.txt
    ```

2. **Run the application**:
    ```bash
    uvicorn cmd.main:app --reload
    ```

3. **Access the application**: Open your browser and go to `http://127.0.0.1:8000`.

## API Endpoints

### Create Event

**POST** `/events/`

Request body:
```json
{
    "name": "Event Name",
    "location": "Event Location",
    "event_date": "2024-05-20",
    "description": "Event Description"
}
```

### Get Event

**GET** `/events/{name}`

### Add Attendee

**POST** `/events/{event_name}/attendees/`

Request body:
```json
{
    "attendee": "Attendee Name"
}
```

### Remove Attendee

**DELETE** `/events/{event_name}/attendees/`

Request body:
```json
{
    "attendee": "Attendee Name"
}
```

## License

This project is licensed under the MIT License.
```
