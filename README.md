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

You can now run E2E test:
```sh
# Requires curl & jq installed
bash test_e2e.sh
```

Follow application logs with:
```sh
docker-compose logs -f
```

Finally, tear down with:
```sh
docker-compose down -v
```

### Minikube

Start a local cluster:
```sh
# default settings
minikube start --disk-size=10g

# KVM2 driver if using libvirt
minikube start --disk-size=10g --vm-driver kvm2
```

Build images in Minikube's environment:
```sh
eval $(minikube docker-env)
docker-compose build
```

Deploy application to Kubernetes:
```sh
kubectl apply -f kubernetes/
```

Check deployment, waiting for Running status on all pods:
```sh
kubectl get all
```

Run E2E test on the cluster:
```sh
bash test_e2e.sh "$(minikube service movies --url)"
```

> If you reached this point, you may now access other API endpoints as described below.

Try getting movies, etc.:
```sh
% curl $(minikube service movies --url)/movies
[]
```

Tear down application:
```sh
kubectl delete -f kubernetes/
```

Optionally, delete local cluster:
```sh
minikube delete
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

curl "$(minikube service movies --url)/movies" --data '{"id":"tron-legacy","title":"TronLegacy","director":"Joseph Kosinski","actors":["Jeff Bridges","Olivia Wilde"]}'
```

### Get movie list

```sh
GET /movies
```

> Example
```
curl 'http://localhost:8080/movies'

curl "$(minikube service movies --url)/movies"
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

curl "$(minikube service movies --url)/reviews/:movie_id" --data '{"text":"Review","rating":5.0}'
```

### Get reviews

```sh
GET /reviews/:movie_id
```

> Example
```
# Replace :movie_id with existing identifier from GET /movies call
curl 'http://localhost:8080/reviews/:movie_id'

curl "$(minikube service movies --url)/reviews/:movie_id"
```

### Delete movie

```sh
DELETE /movies/:movie_id
```

> Example
```
# Replace :movie_id with existing identifier from GET /movies call
curl -X DELETE 'http://localhost:8080/movies/:movie_id'

curl -X DELETE "$(minikube service movies --url)/movies/:movie_id"
```
