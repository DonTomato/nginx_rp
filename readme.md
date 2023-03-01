The create docker images:

```bash
docker build . -t docker.io/dontomato/serviceexp:gs-ping --platform linux/aarch64
docker push docker.io/dontomato/serviceexp:gs-ping

docker build . -t docker.io/dontomato/serviceexp:gs-rp --platform linux/aarch64
docker push docker.io/dontomato/serviceexp:gs-rp
```
