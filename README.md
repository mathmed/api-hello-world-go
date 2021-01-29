# Simple Pokemon API - Go

Simple API to get Pokemons information

# Pre-requisites

- docker
- docker-compose

# Install

- Run `$ mv .env.example .env`
- Set port value in .env
- Run `$ docker-compose build && docker-compose up`

# Usage

| Endpoint                    | Description                            | Method | Example Return                                                                                          |
| :-------------------------- | :------------------------------------- | ------ | ------------------------------------------------------------------------------------------------------- |
| `/pokemons/{pokedexNumber}` | Returns informations about the pokemon | GET    | { "status": { "message": "Pokemon not Found", "code": 400 } }                                                                      |
| `/pokemons/{pokemonName}`   | Returns informations about the pokemon | GET    | { "status": { "message": "Ok", "code": 200 }, "data": { "id": 1, "name": "bulbasaur", "weight": 69 } } |
