FROM golang:1.22.4

RUN wget https://github.com/aquasecurity/trivy/releases/download/v0.52.2/trivy_0.52.2_Linux-64bit.deb && \
    dpkg -i trivy_0.52.2_Linux-64bit.deb
WORKDIR /go/src/app
COPY . .
RUN go get -d -v
EXPOSE 8080
RUN go build -o main .
CMD ["./main"]

# time DOCKER_BUILDKIT=1 docker image build --no-cache -t my-image-name-come-here .