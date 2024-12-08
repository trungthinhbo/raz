# Create a stage for building the application.
ARG GO_VERSION=1.23.3
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
LABEL org.opencontainers.image.source=https://github.com/trungthinhbo/raz
WORKDIR /src

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN ls
RUN mv tailwindcss-linux-x64 /usr/bin/tailwindcss

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]