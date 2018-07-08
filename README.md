# Movies

An experiment in Go and microservices.

## Deployment

### Docker Compose

Using Compose, you may start the application in the background:
```sh
docker-compose up -d
```

Check container status with:
```sh
docker-compose ps
```

Follow application logs with:
```sh
docker-compose logs -f
```

Finally, tear down with:
```sh
docker-compose down -v
```

## API endpoints

### Create movie

```sh
POST /movies
{
  "id": "tron-legacy",
  "title": "TronLegacy",
  "director": "Joseph Kosinski",
  "actors": ["Jeff Bridges", "Olivia Wilde"]
}
```

> Example
```
curl 'http://localhost:8080/movies' --data '{"id":"tron-legacy","title":"TronLegacy","director":"Joseph Kosinski","actors":["Jeff Bridges","Olivia Wilde"]}'
```

### Get movie list

```sh
GET /movies
```

> Example
```
curl 'http://localhost:8080/movies'
```

### Create review

```sh
POST /reviews/:movie_id
{
  "text": "Review",
  "rating": 5.0
}
```

> Example
```
# Replace :movie_id with existing identifier from GET /movies call
curl 'http://localhost:8080/reviews/:movie_id' --data '{"text":"Review","rating":5.0}'
```

### Get reviews

```sh
GET /reviews/:movie_id
```

> Example
```
# Replace :movie_id with existing identifier from GET /movies call
curl 'http://localhost:8080/reviews/:movie_id'
```

### Delete movie

```sh
DELETE /movies/:movie_id
```

> Example
```
# Replace :movie_id with existing identifier from GET /movies call
curl -X DELETE 'http://localhost:8080/movies/:movie_id'
```
