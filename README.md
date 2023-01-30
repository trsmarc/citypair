# Citypair

## Background

To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

Required JSON structure:

```
[["SFO", "EWR"]] => ["SFO", "EWR"]
[["ATL", "EWR"], ["SFO", "ATL"]] => ["SFO", "EWR"]
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]

```

## Architecture

This project follows SOLID & Clean architecture

## API Document

Calculate city pair from input flights array

<details>
 <summary><code>POST</code> <code><b>/</b></code> <code>v1/calculate</code></summary>

##### Request body

> | name   | type     | data type  | description                                                                                                         |
> | ------ | -------- | ---------- | ------------------------------------------------------------------------------------------------------------------- |
> | flight | required | [][]string | a list of flights, which are defined by a source and destination airport code e.g. [["SFO", "ALT"], ["ALT", "EWR"]] |

##### Responses

> | http code | content-type               | response                                 |
> | --------- | -------------------------- | ---------------------------------------- |
> | `200`     | `text/plain;charset=UTF-8` | `{result: ["SFO", "EWR"]}`                         |
> | `400`     | `application/json`         | `{"code":"400","message":"Bad Request"}` |
> | `405`     | `text/html;charset=utf-8`  | None                                     |

##### Example cURL

> ```javascript
>  curl -X POST -H "Content-Type: application/json" --data @post.json http://localhost:8080/v1/calculate --data-raw '{"flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]}'
> ```

</details>

---

Service health check

<details>
 <summary><code>GET</code> <code><b>/</b></code> <code>v1/healthcheck</code></summary>

##### Request body

> | name | type | data type | description |
> | ---- | ---- | --------- | ----------- |
> | none | -    | -         | N/A         |

##### Responses

> | http code | content-type               | response                                           |
> | --------- | -------------------------- | -------------------------------------------------- |
> | `200`     | `text/plain;charset=UTF-8` | `"passed"`                                         |
> | `500`     | `application/json`         | `{"code":"500","message":"Internal server error"}` |

##### Example cURL

> ```javascript
>  curl -X POST --data @post.json http://localhost:8080/v1/healthcheck
> ```

</details>

---

## Get Started

### Prerequisites

- installed [Golang 1.19](https://golang.org/)
- or run using [Docker](https://www.docker.com/)

### Using installed go

Below are the steps if you want to run locally without docker

1. Set required environment variable

   ENV=local

2. Set configuration

   Change config/local.yaml configuration value properly

3. Run app

   > make run

4. Make a request
   ```
   curl --location --request POST 'http://127.0.0.1:8080/v1/calculate' \
   --header 'Content-Type: application/json' \
   --data-raw '{
       "flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
   }'
   ```
   response:
   ```
   {result: ["SFO", "EWR"]}
   ```

### Run Test

go test -v ./...

### Using Docker Compose

Build docker image and run using docker compose

> docker compose up

Quit & Cleanup

Press Ctrl+C to quit in console then run these commands below to clean up all things:

> docker-compose rm -v

### Layout

```tree
├── cmd
│   ├── server/
│   |   └── route.go
│   |   └── server.go
│   └── main.go
├── config
│   └── local.yaml
├── internal
│   ├── config/
│   ├── flight/
│   └── healthcheck/
└── pkg
    ├── error/
    └── log/
```

A brief description of the layout:

- `cmd` Main applications for this project.
- `pkg` library code intend to use by external applications
- `internal` private application and library code included business logic
  - `flight` contain layered api, handler, service for flight module
  - `healthcheck` api for healthcheck to get status of all connected service providers
- `config` configuration file templates or default configs.

## Notes

- Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
