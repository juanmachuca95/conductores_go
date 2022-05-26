# Deployment

Para este challenge utilizaré docker y docker-compose para crear el contenedor de la api y para la conexión a la base de datos con mysql, simplemente configuraré las variables de entorno (```.env```) al servicio externo. 

## Step by Step

Iniciamos nuestra conexion via ssh o de su preferencia para conectarse con su instancia para deploy de aplicación de acuerdo a su proveedor, en mi caso será un ubuntu 20.04 - en digitalocean. Primero clonamos el repositorio:

```sh
git clone https://github.com/juanmachuca95/conductores_go
```

```sh
cd conductores_go
nano .env
```

2. Seteamos nuestras variables de entorno:

```
DATABASE=""
PORT="3306"
USERNAME=""
PASSWORD=""
HOSTNAME=""

TOKEN_API_AUTH=""
API_PORT=":8080"
```


3. Una vez especificada nuestra conexión a la base de datos externa. 
En nuestro directorio raiz ```/db/mysq.sql``` se encuentra el archivo que importaremos en nuestro servicio mysql.
Al termniar la importación podemos continuar con docker.

```docker
docker-compose build
docker-compose up -d 
```

4. 👍 Todo listo! La esta en marcha.
<br><br>
## Primeras pruebas
Lancemos algunas solicitudes a la api para comprobar su correcto funcionamiento. Vamos logearnos con un usuario administrador.

```sh
curl -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{"email":"admin@spaceguru.com", "password":"123456"}'
```

Con el token de la respuesta. Nos dirigimos a https://jwt.io y pegamos nuestro token.

```json
{
  "user": {
    "id": 1,
    "name": "Administrador",
    "email": "admin@spaceguru.com"
  },
  "roles": [
    "admin"
  ],
  "exp": 1653711999
}
```

Ingresamos una cuenta de conductor:
```bash 
curl -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{"email":"conductor@spaceguru.com", "password":"123456"}'
```

```json
{
  "user": {
    "id": 102,
    "name": "Conductor test",
    "email": "conductor@spaceguru.com"
  },
  "roles": [
    "conductor"
  ],
  "exp": 1653712180
}
```


2. Obtener los conductores - utilizando paginación

```sh
curl -X GET http://localhost:8080/conductores -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImlkIjoxLCJuYW1lIjoiQWRtaW5pc3RyYWRvciIsImVtYWlsIjoiYWRtaW5Ac3BhY2VndXJ1LmNvbSJ9LCJyb2xlcyI6WyJhZG1pbiJdLCJleHAiOjE2NTM2OTYzMzJ9.TKWpE2ut1HfKFlsjmi7jvpWRd_jwJepcW_lAwOeVp00' -d '{"page":5}'
```

```json
{
  "conductores": [
    {
      "id": 11,
      "users_id": 31,
      "name": "Mrs. Clarissa McDermott",
      "email": "gCAdxcw@AxUsYny.org",
      "matricula": "472",
      "vehiculo": "Prof. Emely Goodwin",
      "created_at": "2022-05-25 17:11:13",
      "updated_at": "2022-05-25 17:11:13"
    },
    {
      "id": 12,
      "users_id": 37,
      "name": "Miss Zita Emard",
      "email": "ggLhBly@UXspyFv.ru",
      "matricula": "1502",
      "vehiculo": "Miss Dasia Nienow",
      "created_at": "2022-05-25 17:11:13",
      "updated_at": "2022-05-25 17:11:13"
    },
    {
      "id": 13,
      "users_id": 45,
      "name": "Queen Juana Schroeder",
      "email": "hrcxbwv@VkaByfn.net",
      "matricula": "171",
      "vehiculo": "Queen Mikayla Walker",
      "created_at": "2022-05-25 17:11:13",
      "updated_at": "2022-05-25 17:11:13"
    },
    {
      "id": 14,
      "users_id": 39,
      "name": "Princess Kenyatta Mertz",
      "email": "Hygujte@hVVrYJT.net",
      "matricula": "1719",
      "vehiculo": "Queen Jeanne Jacobs",
      "created_at": "2022-05-25 17:11:13",
      "updated_at": "2022-05-25 17:11:13"
    },
    {
      "id": 15,
      "users_id": 34,
      "name": "Princess Hortense Gleason",
      "email": "IfKTxIY@ZrigpOe.com",
      "matricula": "1371",
      "vehiculo": "Lady Callie Labadie",
      "created_at": "2022-05-25 17:11:13",
      "updated_at": "2022-05-25 17:11:13"
    }
  ],
  "status": "Success"
}
```