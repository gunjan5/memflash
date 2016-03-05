FROM iron/go:1.5.1
WORKDIR /app

# copy binary into image
COPY app /app/

EXPOSE 8080

ENTRYPOINT ["./app"]