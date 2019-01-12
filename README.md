# Realtime-API-Monitoring

This project is based on Python Django, Golang Core, Postgresql, Influxdb, Grafana and Docker. 
This project is under construction.

Run Postgresql container to store API request details. <br>
    
```bash
docker run --name postgresql -d \
    -p 5432:5432 \
    --env 'DB_NAME=YOUR_POSTGRES_NAME' \
    --env 'DB_USER=YOUR_POSTGRES_USER' --env 'DB_PASS=YOUR_POSTGRES_PASSWORD' \
    --env 'DB_EXTENSION=pg_trgm' \
    --volume /srv/docker/postgres/api:/var/lib/postgresql_api \
    postgres:latest
```
