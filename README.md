Trivy-Image-Scanner is a web interface implementation that makes use of [aquasec/trivy](https://github.com/aquasecurity/trivy) tool to scan image for vulnerabilities and 
display the scan result in a `JSON` or `Table` format, as specified.

# Pulling the image
The image is pushed to dockerhub repository [`sarrocksdev/trivy-image-scanner`](https://hub.docker.com/repository/docker/sarrocksdev/trivy-image-scanner/general). 
To pull the image:
```bash
docker pull sarrocksdev/trivy-image-scanner
```
# Running the image container
The port `8888` is exposed by default.

To run the image on port `9000` locally:
```bash
docker run -d -p 9000:8888 sarrocksdev/trivy-image-scanner
```
Access the service running at [localhost:9000](http://127.0.0.1:9000)

