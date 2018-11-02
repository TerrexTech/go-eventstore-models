#!/usr/bin/env bash

echo $(pwd)

docker-compose -f ./test/docker-compose.yaml up -d

function ping_cassandra() {
  docker exec -it cassandra /usr/bin/nodetool status | grep UN
  res=$?
}

echo "Waiting for Cassandra to be ready."

# Wait for Cassandra to be ready
max_attempts=30
cur_attempts=0
ping_cassandra
while (( res != 0 && ++cur_attempts != max_attempts ))
do
  ping_cassandra
  echo Attempt: $cur_attempts of $max_attempts
  sleep 1
done

if (( cur_attempts == max_attempts )); then
  echo "Cassandra Timed Out."
  exit 1
fi

# The Cassandra image takes more time to be ready despite
# nodetool-status being success.
# There has to be a better way than this.
echo "Waiting additional time for Cassandra to be ready."
add_wait=30
cur_add_wait=0
while (( ++cur_add_wait != add_wait ))
do
  echo Additional Wait: $cur_add_wait of $add_wait seconds
  sleep 1
done

go test -v -race ./...
