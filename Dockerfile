FROM golang:1.13 as builder

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go build .

ENTRYPOINT ["./gotmpl"]

FROM scratch
COPY --from=builder /app/gotmpl /
COPY --from=builder /app/manifest.gotmpl /
COPY --from=builder /app/args.yaml /

ENTRYPOINT ["/gotmpl"]
CMD ["manifest.gotmpl", "args.yaml"]
