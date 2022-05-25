utilidades:
gin:
go get -u github.com/gin-gonic/gin


```
curl -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{"email":"admin@spaceguru.com", "password":"123456"}'
{"_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIyLTA1LTI1VDA5OjUzOjQ3Ljk5MTcyNzEwOC0wMzowMCIsInN1YiI6IjEiLCJ1c2VyIjp7ImlkIjoxLCJuYW1lIjoiSnVhbiBHYWJyaWVsIE1hY2h1Y2EiLCJlbWFpbCI6ImFkbWluQHNwYWNlZ3VydS5jb20iLCJwYXNzd29yZCI6IiQyeSQxMCR1eVVkb1RuSmRrSUt4aDdnNHQudUZlT3JuY1JuT0JhU0d0U3pram55QkRWUmZWeUtxSzZoYSIsImNyZWF0ZWRfYXQiOiIiLCJ1cGRhdGVkX2F0IjoiIn19.bDLXBzTn34TDZwq8dEHudSZcQ81YAj_A5evl4tzstOI","status":"Estas loggeado"}%
```

Ejemplo de login para usuario administrador

```json
{
    "email":"admin@spaceguru.com",
    "password":"123456"
}
```