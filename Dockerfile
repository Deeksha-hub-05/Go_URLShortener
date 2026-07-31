FROM golang:alpha as builder

#stage1
RUN mkdir /build

ADD ./build/

WORKDIR /build

RUN go build -o main .


#stage2
FROM alpha

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY . /app

COPY --from=builder /build/main /app/

WORKDIR /app

EXPOSE 3000

CMD [ "./main" ]

