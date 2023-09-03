

FROM --platform=linux/amd64 golang:1.20

WORKDIR /app

COPY . .

# RUN go get
RUN go build -o battery-tracking .

ENTRYPOINT [ "/app/battery-tracking" ]
EXPOSE 8080
