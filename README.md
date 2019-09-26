# compelo

![CircleCI](https://img.shields.io/circleci/build/github/florianehmke/compelo?style=flat-square) ![Docker Build Status](https://img.shields.io/docker/build/florianehmke/compelo?style=flat-square)

Program for elevating team spirit.

## Development

Build with `-tags=dev`, otherwise generated `*_vfsdata.go` files for the `frontend` and `db` package are expected.
Start the backend with `-dev` flag (enables a `CORS` middleware).

Start the frontend inside `frontend/compelo` via `npm install && npm run start`.

