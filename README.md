# udp-console

tiny udp server to log messages to the console

**not something you want to use in production :)**

i initially put this together to log statsd UDP messages during local
development. send it a UDP message, and it'll print that message to the console

## Usage

The easiest way to utilize is via docker

```sh
HOST="0.0.0.0" PORT=8125 docker run -it -p 8125:8125/udp knomedia/udp-console
```

or docker-compose to cordinate with your other services locally


```yaml
version: "2.3"
services:
  statsd:
      image: knomedia/udp-console
      ports:
        - "8125:8125/udp"
      environment:
        HOST: "0.0.0.0"
        PORT: 8125
```

## ENV Vars

Pass in `HOST` and `PORT` ENV vars to customize where the server is listening
(they default to `0.0.0.0` and `8125` respectively)
