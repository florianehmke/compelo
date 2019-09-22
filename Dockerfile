### Build Image
FROM golang:1.13-buster

RUN set -e \
  && apt-get update -q \
  && apt-get install -y -q --no-install-recommends make nodejs npm \
  && npm install -g @angular/cli \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY Makefile go.mod go.sum /usr/src/
COPY api /usr/src/api
COPY cmd /usr/src/cmd
COPY db /usr/src/db
COPY frontend /usr/src/frontend
COPY game /usr/src/game
COPY match /usr/src/match
COPY player /usr/src/player
COPY project /usr/src/project
COPY rating /usr/src/rating
COPY stats /usr/src/stats

WORKDIR /usr/src
RUN make

### APP Image
FROM debian:buster
RUN adduser --home /srv --no-create-home --system --uid 1000 --group app
RUN chown 1000:1000 /srv
COPY --from=0 /usr/src/compelo /usr/local/bin/compelo

VOLUME /srv
ENV COMPELO_DB_PATH=/srv/db.sql
ENV COMPELO_PORT=8080
EXPOSE 8080/tcp

USER app
CMD ["/usr/local/bin/compelo"]
