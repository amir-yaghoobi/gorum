# Gorum 

[![Drone Status](https://cloud.drone.io/api/badges/irgophers/gorum/status.svg)](https://cloud.drone.io/irgophers/gorum) [![Docker Status](https://img.shields.io/docker/cloud/build/irgophers/gorum.svg)](https://hub.docker.com/r/irgophers/gorum)

Message board and forum created with love by and for Iranian gophers.

_This project is currently under development._

## Run with Docker

The latest development version of the docker image can be obtained by running:

```bash
$ docker pull irgophers/gorum:latest
```

However, `latest` images are intended to be used for previews and testing new features. One must consider using a release tagged image in production.

Do not forget to run database migrations. First, migrations need to be initialized:

```bash
$ docker exec [container] migrate init
```

Note that this command needs to be run only once.

Then, run migrations:

```bash
$ docker exec [container] migrate
```

## License

This software is—and always will be—a free software. It is developed by a free community of Iranian gophers and its resources will be available for everyone. No ads, no commercials, no bullshit.

Its credits goes to the developers who are listed in the `CONTRIBUTORS` file.