# Prueba ceiba: *API para pagos de arrendamientos*

- Bajar dependencias: `go mod vendor`

- Configuraciones necesarias: `.env`

- Contenedor DB: `docker-compose up -d`

- **Importar DB**: `./db.sh`

- Para iniciar con hot-reload: `air` [Binario: https://github.com/cosmtrek/air]

- Ejecución preparando entorno: `./air.sh`

![demo run air](./x.img/air.png)
![demo exit air](./x.img/air_exit.png)

- Ejecución directa sin esperar instancia db: `go run .`

- **Tests**: `./tests/routes_test.go`

![demo test Post](./.img/test.png)