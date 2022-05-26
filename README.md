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


Ejemplo de login para usuario administrador

```json
{
    "email":"admin@spaceguru.com",
    "password":"123456"
}
```