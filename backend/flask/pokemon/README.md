
# Pokemon Data API
This project provides an API to manage and display Pokémon data, including their types, abilities, height, weight, and other details. The data is stored in a MySQL database, and the application is built with Flask.

## Features
- Retrieve Pokémon data.
- Pagination and filtering for Pokémon listings.
- Ability to store Pokémon data in a MySQL database.
- Built with Flask and MySQL.

---

Make sure you have the following installed:
- Docker (if you want to run the app using Docker).
- Python 3.11+ (if you want to run the app without Docker).
- MySQL database setup.

Create a .env file at the root of the project with the following contents to configure the database connection:

```
MYSQL_HOST=mysql
MYSQL_DATABASE=pokemon_db
MYSQL_USER=root
MYSQL_PASSWORD=rootpassword
```

or rename .env.example to .env

then if you used docker you can run by this command
```
docker-compose up --build
```

or run without docker

```
1. pip3 install -r requirement.txt
2. setup database in you local
3. python3 crawler.py -o pokemon_data.csv # to save the crawler to you local as csv file
4. python3 seeder.py -o pokemon_data.csv # to insert the data to database
5. python3 app.py # to running the server
```

## Sample cURL Requests and Responses
Sample Request 1: Get all Pokémon with Pagination

`curl "http://127.0.0.1:8888/api/v1/pokemon?page=1&limit=10"`

Response:
```
{
  "data": {
    "page": 1,
    "pokemon": [
      {
        "abilities": [ "Synchronize", "Inner Focus", "Magic Guard"],
        "created_at": "Sat, 25 Jan 2025 17:54:09 GMT",
        "deleted_at": null,
        "generation": 1,
        "height": "0.9\u00a0m (2\u203211\u2033)",
        "id": 63,
        "image": "https://img.pokemondb.net/sprites/home/normal/2x/abra.jpg",
        "name": "Abra",
        "species": "Psi Pok\u00e9mon",
        "types": ["Psychic]",
        "updated_at": "Sat, 25 Jan 2025 17:54:09 GMT",
        "weight": "19.5\u00a0kg (43.0\u00a0lbs)"
      },
     {...}
    ],
    "total_count": 68,
    "total_pages": 7
  },
  "message": "Success",
  "status_code": 200
}
```

* you can filter by name & gen
