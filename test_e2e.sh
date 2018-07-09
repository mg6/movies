#!/bin/bash

set -xeuo pipefail

API="${MOVIE_API_URL:-http://localhost:8080}"

get_movies() {
  curl -s "$API/movies"
}

create_movie() {
  curl -s "$API/movies" --data '{"id":"tron-legacy","title":"TronLegacy","director":"Joseph Kosinski","actors":["Jeff Bridges","Olivia Wilde"]}'
}

get_first_movie_id() {
  curl -s "$API/movies" | jq -r '.[0].id'
}

delete_movie() {
  curl -s -X DELETE "$API/movies/$1"
}

create_review() {
  curl -s "$API/reviews/$1" --data '{"movie":"tron-legacy","text":"Review","rating":5.0}'
}

get_reviews() {
  curl -s "$API/reviews/$1"
}

test_e2e() {
  if [ $(get_movies | jq length) -eq 0 ]; then
    create_movie
    [ $(get_movies | jq length) -eq 1 ]
  fi

  movie_id="$(get_first_movie_id)"

  [ $(get_reviews "$movie_id" | jq length) -eq 0 ]
  create_review "$movie_id"
  [ $(get_reviews "$movie_id" | jq length) -eq 0 ]

  sleep 2
  [ $(get_reviews "$movie_id" | jq length) -eq 1 ]

  delete_movie "$movie_id"
  [ $(get_movies | jq length) -eq 0 ]
}

test_e2e
