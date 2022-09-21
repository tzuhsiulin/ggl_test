# -------
FROM golang:1.18.4-buster as bin
ADD . /backend
WORKDIR /backend
RUN apt-get update && apt-get install make
RUN cd /backend && make clean && make bin

# -------
FROM alpine:3.14
LABEL maintainer="Chris Lin"

RUN addgroup -S ggltest && adduser -S ggltest -G ggltest
USER ggltest

RUN mkdir -p /home/ggltest/data/
COPY --from=bin /backend/bin/api /home/ggltest/api
COPY --from=bin /backend/bin/db_migration /home/ggltest/db_migration
COPY --from=bin /backend/data/migration_scripts /home/ggltest/data/migration_scripts

WORKDIR /home/ggltest
ENTRYPOINT [ "sh", "-c", "./db_migration && ./api" ]