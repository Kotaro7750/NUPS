FROM mysql:latest
WORKDIR /init
COPY ./ki.csv /init/
COPY init.sql /docker-entrypoint-initdb.d/1-init.sql
COPY ki_table.sql /docker-entrypoint-initdb.d/2-init.sql