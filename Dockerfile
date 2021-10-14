FROM alpine AS builder
WORKDIR /src
RUN apk add go

ADD go.mod /src
ADD go.sum /src
RUN go mod download

ADD . /src
RUN go build -o bpk-dx .

#---

FROM alpine
COPY --from=builder /src/views /views
COPY --from=builder /src/bpk-dx /
CMD /bpk-dx

#CMD ./bpk-dx