# Image Collection Center Core

![](https://github.com/xylonx/icc-core/actions/workflows/ci.yml/badge.svg)

![](./doc/img/icc-title.png)


This project is the backend part of the ICC(Image Collection Center). You can find the frontend part at [icc-frontend](https://github.com/xylonx/icc-frontend)

## What is Image Collection Center

In brief, ICC can show(without auth) and upload images(with auth). In this perspective, it is just like a pic-host service.
Moreover, it provides some improved functions like tagging the images and searching by tags. It is simple but stronger. Additionally, it also provides out-of-box API.

Unlike other pic-host services, ICC keeps min auth info. Up to now, only image uploading and tag modifying actions need auth - just a token.

ICC just tracks uploaded bytes for every token preventing malicious users.

You can preview it at [here](https://icc.xylonx.com)

For backend, the real image storage is a S3 compatible backend. (for those s3 no-compatible backend, you can proxy it by [s3-gateway](https://github.com/xylonx/s3-gateway)).

> Therefore, thanks [Backblaze](https://backblaze.com/) cheep Object Storage providing :)

It provides 2 different type APIs: 

- HTTP
- *protobuf*(in progress)
  > you can find the protobuf file at `proto/icc/icc.proto`.

## API Doc

available API doc:

- [v0.2.0](./doc/api/api-v0.2.0.yaml)

## How to Deploy

**base service need**

- PostgreSQL
- S3 access info

you can file the template config file at `config.default.yaml`. its content like below:

> Fill your DB and S3 info into `database.postgres.dsn` and `storage.s3` blocks.

```yaml
# application - configure the server
application:
    # grpc bind host
    grpc_host: 0.0.0.0
    # grpc bind port
    grpc_port: 30000
    # http bind host
    http_host: 0.0.0.0
    # http bind port
    http_port: 5000
    # http server read timeout. if read(body) action exceeds, it will close the connection
    # default value is 60
    http_read_timeout_seconds: 60
    # http server read timeout. if write(body) action exceeds, it will close the connection
    # default value is 60
    http_write_timeout_seconds: 60
    # admin token - used to generated auth Bearer token
    admin_token:
    # http allow origins - using for cors
    http_allow_origins:
        - https://icc.xylonx.com

# database - configure database connection
database:
    # using postgres database
    # visit gorm connecting to database(https://gorm.io/docs/connecting_to_the_database.html)
    # for more information
    postgres:
        # postgres database dsn
        dsn: host= user= password= dbname=icccore port=5432 sslmode=disable TimeZone=Asia/Shanghai
        # database max open connection
        # default value is 10
        max_open_conn: 10
        # database max idle connection keeping
        # default value is 100
        max_idle_conn: 100
        # database max connection keeping seconds
        # default value is 600
        max_life_seconds: 600

# storage - configure image storage
storage:
    # cdn host - using for generated cdn image link
    cdn_host:
    # sign upload seconds - duration for signing upload
    sign_upload_seconds: 600
    # using s3 param
    s3:
        # s3 endpoint
        endpoint:
        # s3 access id
        access_id:
        # s3 access key
        access_secret:
        # s3 bucket
        bucket:
        # s3 region
        region:
```

`admin_token` is a meta-token using to generate other token. Don't leak it.

## Docker Deploy

ICC core provides the docker image tracking the latest git tag.

recommend deploying it using docker-compose like below:

> remember to fill the config path and psql password

```docker-compose
version: '3.9'

services:
  icc-core:
    image: xylonx/icc-core:${version:-latest}
    container_name: icc-core
    volumes:
      - #{fill config file path}:/opt/icc-core/config.yaml
    environment:
      - GIN_MODE=release
    ports:
      - 5000:5000

  psql:
    image: postgres:14
    restart: always
    container_name: db_postgres
    volumes:
      - $PWD/data/postgres:/var/lib/postgresql/data
    environment:
      LANG: C.UTF-8
      POSTGRES_PASSWORD: 
```