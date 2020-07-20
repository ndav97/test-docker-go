# Backend Tienda Virtual Fiber - Docker

[![N|Solid](https://raw.githubusercontent.com/gofiber/docs/master/static/fiber_v2_logo.svg)](https://nodesource.com/products/nsolid)

## Documentación
[Documentación de Fiber](https://github.com/gofiber/fiber)

### Requisitos

* Instalar Docker en su última version
* Tener activo el WSL 2
* Tener instalada una distro habilitando WSL 2

### Como correr el proyecto

Hacer build de la imagen

```sh
$ docker build -t [nombre_de_la_imagen] .
```

Ejecutar el proyecto desde docker-compose

```sh
$ docker-compose up
```