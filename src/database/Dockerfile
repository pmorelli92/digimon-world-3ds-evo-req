FROM postgres:13.1-alpine
COPY /schema.sql /docker-entrypoint-initdb.d/01.sql

ENV POSTGRES_USER="dbuser"
ENV POSTGRES_DB="digimonsql"
ENV POSTGRES_PASSWORD="dbpassword"
