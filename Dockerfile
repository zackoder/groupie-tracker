FROM golang:1.5-alpine

WORKDIR /app

COPY . .

LABEL ProjectName="groupie-trucker"

LABEL Version="v1.0"

CMD [ "go", "run", "main.go" ]