# Realtime-API-Monitoring

This project is based on Python Django, Golang Core, Postgresql, Influxdb, Grafana and Docker. 
This project is under construction.

Run Postgresql container to store API request details. <br>
` docker run --name postgresql -d \`<br>
    `-p 5432:5432 \`<br>
    `--env 'DB_NAME=YOUR_POSTGRES_NAME' \`<br>
    `--env 'DB_USER=YOUR_POSTGRES_USER' --env 'DB_PASS=YOUR_POSTGRES_PASSWORD' \`<br>
    `--env 'DB_EXTENSION=pg_trgm' \`<br>
    `--volume /srv/docker/postgres/api:/var/lib/postgresql_api \`<br>
    `postgres:latest `
