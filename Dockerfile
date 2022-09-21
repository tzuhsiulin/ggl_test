# -------
FROM golang:1.18.4-buster as bin
ADD . /backend
WORKDIR /backend
RUN apt-get update && apt-get install make
RUN cd /backend && make clean && make bin

# -------
FROM alpine:3.14
LABEL maintainer="Chris Lin"

RUN mkdir -p /ggl_test && mkdir -p /ggl_test/data/
WORKDIR /ggl_test
COPY --from=bin /backend/bin/api /ggl_test/
COPY --from=bin /backend/bin/db_migration /ggl_test/
COPY --from=bin /backend/data/migration_scripts /ggl_test/data/migration_scripts

ENTRYPOINT [ "sh", "-c", "./db_migration && ./api" ]