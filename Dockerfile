# -------------
# build stage
# -------------
FROM golang:1.21-alpine AS build

# Workdir
WORKDIR /src

# Go deps
ADD go.mod go.mod
ADD go.sum go.sum
RUN time go mod download

# Attach sources
ADD . .

# Build
RUN time go build -o cs2bot

# -------------
# runtime stage
# -------------
FROM alpine

# Copy app
WORKDIR /app
COPY --from=build /src/cs2bot /app/

# Entrypoint
ENTRYPOINT ./cs2bot
