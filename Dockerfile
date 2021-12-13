### Frontend Build
FROM node:16 AS FRONTEND

COPY . /usr/src/

WORKDIR /usr/src

RUN make frontend

### Backend Build
FROM golang:1.17-buster AS BACKEND

RUN apt-get install -y -q --no-install-recommends make

COPY . /usr/src/
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
