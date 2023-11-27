# eBusinessCardBe 

# Golang REST API 


## Development Setup

Run the following command to start development server.

```bash
docker compose up
```
> before running this command **docker** and **docker compose** must be installed.

## Technology
- Language (golang)


### Libraries
- Router (gorilla/mux)
- Server (net/http)


## API Documentation

### GetPlaceDetails
> **GET** ``/place-details/{placeId}``

Get details about place spcefied by id.

##### Request

```json
http://localhost:8080/place-details/GXvPAor1ifNfpF0U5PTG0w
```

#### Output

```json
{
  "name": "Casa Ferlin",
  "address": "Stampfenbachstrasse 38, 8006 Zürich",
  "openingHours": [
    {
      "Days": [
        "monday",
        "tuesday",
        "wednesday",
        "thursday",
        "friday"
      ],
      "WorkingBlocks": [
        {
          "start": "11:30",
          "end": "14:00",
          "type": "OPEN"
        },
        {
          "start": "18:30",
          "end": "22:00",
          "type": "OPEN"
        }
      ]
    },
    {
      "Days": [
        "saturday",
        "sunday"
      ],
      "WorkingBlocks": null
    }
  ]
}
```
