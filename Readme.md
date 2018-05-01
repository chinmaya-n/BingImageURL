# BING IMAGE URL service

* Instructions on compiling Go code into a minimal container are taken from [here](https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/)

1. `$ CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .`
2. Copy certs file: `cp /etc/ssl/certs/ca-certificates.crt .`
3. `docker build -t bing-img-url:latest .`
4. `docker run -it --rm --name BingImageURL -p 9001:9001 bing-img-url:latest 9001`
