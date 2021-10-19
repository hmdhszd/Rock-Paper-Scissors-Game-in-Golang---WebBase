FROM golang:latest

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o rockpaperscissors .

CMD ["/app/rockpaperscissors"]