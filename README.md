# HTTP Golang Project Example
An empty project for a Golang Echo HTTP service.  This project includes a docker setup since most docker examples use unstructured golang projects.  Additionally in the `docker-compose.yml` there is a Redis instance just to have an on hand example of how to manage docker service dependencies.

# How to Build
```bash
# For development you most likely just want the HTTP logs, and you want them to stream
# Otherwise use -d for detached mode.
docker-compose up --build app
```
