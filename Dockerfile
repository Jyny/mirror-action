FROM golang

WORKDIR /usr/src/action

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]