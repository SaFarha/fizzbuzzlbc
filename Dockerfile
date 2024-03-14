FROM gcr.io/distroless/base:latest AS distroless

COPY target/app /

CMD ["/app"]