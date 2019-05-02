#FROM alpine:3.4

FROM gcr.io/distroless/base
#RUN apk -U add ca-certificates

EXPOSE 8080

ADD ./hellokube-go /

CMD ["./hellokube-go", "-p", "8080"]
