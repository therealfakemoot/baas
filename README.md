# Installation

## dockerhub
Butts-As-A-Service is available on [dockerhub](https://hub.docker.com/r/therealfakemoot/baas), just pull the image, `docker pull therealfakemoot/baas`, and run the container as described below.

The container may or may not have multiple configuration options, I don't know.

```
docker run -d -p 8008:8008 --name=baas therealfakemoot/baas:latest
```

## source

```
go install github.com/therealfakemoot/baas
./baas
```

# Usage
Using Butts As A Service ( *baas* ) is as simple as running a curl against your domain:
[asciinema link](https://cast.ndumas.com/a/xpesSQ97NLw4rZOMKy716FGWB)
