# Docker Digrams - Generate Diagrams from Docker Compose files

## Intro

Docker Diagrams is a tool that generates diagrams from Docker Compose files.

## Features

1. supports Docker Compose files compliant with Docker Compose speicification
2. supports multiple environment variables in Docker Compose files

## Installation

```bash
go install github.com/alex-guoba/docker-diagrams
```

## Usage

### Generate diagram

Make sure you have a Docker Compose file in your current directory.

```bash
docker-diagrams
```

### Generate diagram with custom options

```bash
docker-diagrams -i=<path_to_your_docker_compose_file> -e=<environment_file>
```


Docker-Diagrams will create a folder in the output directory( default to 'go-diagrams') with the graphviz DOT file and any image assets.

Create an ouput image with any graphviz compatible renderer:

```bash
dot -Tpng <graphviz_file>  <path_to_your_output_image>.png
```


## Exampe

```bash
docker-diagrams -i ./testcase/docker-compose.yml -e testcase/.env.dev 
cd go-diagrams
dot -Tpng docker-compose.dot output.png
```

<p align="center">
  <img alt="Example Page" src="./diagram.png" width="689">
</p>

## Opiptions

```
  -e string
        path to environment file (default ".env")
  -i string
        path to docker compose file (default "docker-compose.yml")
```
