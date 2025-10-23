# Se crea una nueva película para obtener un ID.
POST http://localhost:8080/api/movies
Content-Type: application/json
{
    "title": "Lord of the Strings",
    "synopsis": "A man that can manipulate every string variable.",
    "released_at": "1994-10-14T00:00:00Z",
    "duration_minutes": 255,
    "poster_url": "https://imgs.search.brave.com/8beLkj9LsaTyqIIEdbEELICTyQJkSQnTz6UQfl4x5oI/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pLmVi/YXlpbWcuY29tL2lt/YWdlcy9nL3hwNEFB/ZVN3cGNwb2xKRFQv/cy1sMjI1LmpwZw"
}

HTTP 201
[Asserts]
jsonpath "$.id" isInteger
[Captures]
movieId: jsonpath "$.id"

# ----------------------------------------------------

# Se actualiza la película recién creada usando el ID capturado.
PUT http://localhost:8080/api/movies/{{movieId}}
Content-Type: application/json
{
    "title": "Lord of the Strings 2: The Pointers Counterattack",
    "synopsis": "A man that can manipulate every string variable, now is attacked by pointers.",
    "released_at": "1994-10-14T00:00:00Z",
    "duration_minutes": 255,
    "poster_url": "https://imgs.search.brave.com/8beLkj9LsaTyqIIEdbEELICTyQJkSQnTz6UQfl4x5oI/rs:fit:860:0:0:0/g:ce/aHR0cHM6Ly9pLmVi/YXlpbWcuY29tL2lt/YWdlcy9nL3hwNEFB/ZVN3cGNwb2xKRFQv/cy1sMjI1LmpwZw"
}

HTTP 200
[Asserts]
jsonpath "$.id" == {{movieId}}
jsonpath "$.title" == "Lord of the Strings 2: The Pointers Counterattack"
jsonpath "$.duration_minutes" == 255

# ----------------------------------------------------

# Se verifica que la película se modifico.
GET http://localhost:8080/api/movies/{{movieId}}

HTTP 200
[Asserts]
jsonpath "$.title" == "Lord of the Strings 2: The Pointers Counterattack"

# ----------------------------------------------------

# Se elimina la película usando el ID.
DELETE http://localhost:8080/api/movies/{{movieId}}

HTTP 204

# ----------------------------------------------------

# Se verifica que la película ya no existe.
GET http://localhost:8080/api/movies/{{movieId}}

HTTP 404
[Asserts]
jsonpath "$.message" == "invalid movie identifier"

# ----------------------------------------------------

# Se verifica que se elimina una pelicula inexistente
DELETE http://localhost:8080/api/movies/{{movieId}}

HTTP 404
[Asserts]
jsonpath "$.message" == "invalid movie identifier"
