### Frontend Build
FROM node:11 AS FRONTEND

COPY frontend /usr/src/frontend
COPY Makefile  /usr/src/

WORKDIR /usr/src

RUN make frontend

### Backend Build
FROM golang:1.13-buster AS BACKEND

RUN apt-get install -y -q --no-install-recommends make

COPY Makefile go.mod go.sum /usr/src/
COPY cmd /usr/src/cmd
COPY pkg /usr/src/pkg
COPY internal /usr/src/internal

RUN mkdir -p /usr/src/frontend/compelo/
COPY frontend/*.go /usr/src/frontend/
COPY --from=FRONTEND /usr/src/frontend/compelo/dist /usr/src/frontend/compelo/dist

WORKDIR /usr/src

ENV GOOS=linux
ENV GOARCH=amd64
RUN make backend

### APP Image
FROM debian:buster
RUN adduser --home /srv --no-create-home --system --uid 1000 --group app
RUN chown 1000:1000 /srv
COPY --from=BACKEND /usr/src/compelo /usr/local/bin/compelo

VOLUME /srv
ENV COMPELO_DB_PATH=/srv/db.sql
ENV COMPELO_PORT=8080
EXPOSE 8080/tcp

USER app
CMD ["/usr/local/bin/compelo"]
