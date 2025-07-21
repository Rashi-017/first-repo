from go-1.23.1-alpine

workdir ./app

copy go mod go sum ./
copy . .

run go mod download

run go build -o .main

cmd ["./main"]


