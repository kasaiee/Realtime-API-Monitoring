# Realtime-API-Monitoring

This project is based on Python Django, Golang Core, Postgresql, Influxdb, Grafana and Docker. 
This project is under construction.

Run Postgresql container to store API request details.
` docker run --name postgres_name -d \
    -p 5432:5432 \
    --env 'DB_NAME=YOUR_POSTGRES_NAME' \
    --env 'DB_USER=postgres_user' --env 'DB_PASS=' \
    --env 'DB_EXTENSION=pg_trgm' \
    --volume /srv/docker/postgres/api:/var/lib/postgresql_api \
    postgres:latest `
