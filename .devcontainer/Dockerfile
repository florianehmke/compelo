FROM mcr.microsoft.com/vscode/devcontainers/typescript-node:16-buster

ARG GO_VERSION=1.17.5

RUN npm install -g @angular/cli

RUN wget https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz

EXPOSE 4200
EXPOSE 8080