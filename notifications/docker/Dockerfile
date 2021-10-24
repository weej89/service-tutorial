FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./
RUN cd notifications && go build -o /notifications
EXPOSE 4040
CMD ["/notifications"]