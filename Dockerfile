FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
WORKDIR $GOPATH/src/legislator-query-api
COPY . .
RUN cp cmd/legislator-query-api/*.go .
RUN ls -ltr
RUN go get -d -v
RUN  CGO_ENABLED=1 go build -o legislators-api

FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/legislator-query-api/legislators-api .
COPY --from=build-env /go/src/legislator-query-api/home.html /app/home.html
COPY --from=build-env /go/src/legislator-query-api/index.html /app/index.html
COPY --from=build-env /go/src/legislator-query-api/about.html /app/about.html
COPY --from=build-env /go/src/legislator-query-api/state.html /app/state.html
COPY --from=build-env /go/src/legislator-query-api/data /app/data

EXPOSE 8080

CMD ["./legislators-api"]


