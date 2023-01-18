# NodeJS_Angular13_NodeJs_Express

# Namazu Restaurant
Made by [`Santiago Soler Llin`](https://github.com/SantiSL5)  and  [`Salva Roig Bataller`](https://github.com/SalRB)

## Prerequisites

* [npm](https://www.npmjs.com/)
* [docker](https://www.docker.com/)
* [docker-compose](https://docs.docker.com/compose/)

## Starting up the app

1. Clone the repo.

2. Create the following .env file on the corresponding folders of the repo with this variables:

    1. /
    * UID=1000
    * GID=1000
    * MYSQL_DATABASE
    * MYSQL_USER
    * MYSQL_PASSWORD
    * MYSQL_ROOT_PASSWORD

    2. /backend/go
    * DB_HOST="mysql"
    * DB_PORT=3306
    * DB_USERNAME
    * DB_PASSWORD
    * DB_DATABASE
    * SECRET

    3. /backend/laravel
    * Create the file following the .env.example file structure.

3. Create a Secret.js file on /frontend/vue3/src folder with this content:

    export default {
        LARAVEL_APP_URL : "http://localhost:6001/api",
        GO_APP_URL : "http://localhost:4000/api",
    }
  
4. Go to repo main folder and do 'docker-compose up'

Following this steps, app is running on [localhost:8080](http://localhost:8080).

## Features

This application have the following modules.

Module | Description
:--- | :---
Home | Main page of the application where you can see a carousel with categories and an infinite scroll with dishtypes.
Reservations | Show available tables depending on the selected filters ant let's reserve one.
Menu | Shows the manu of the restaurant separated by dish types.
Login | It allows you to register and login in the application.
Panel Admin | Admin users can manage every DB table using CRUDs.
Reservations | Logged users can reserve tables and see future reservations in a modal window.

## Technologies

### Deploy

The technology used for deploy is [docker](https://www.docker.com/)

  * Docker
  * docker-compose
  * Env files configuration

### Frontend

The technology used for the client is [Vue3](https://vuejs.org/) in its 3.2.45 version. 

  * Routes
  * Components
      * Reusable Components
  * Authentication
      * Guards
      * Interceptors
      * Services
      * JWT Token
  * Store and composables
  * Toastr
  * Infinite-scroll
  * Pagination
  * Table filters with datepicker
  * Validation

### Backend 1

The technology used for the server is [Gin-Gonic](https://gin-gonic.com/) in its 1.7.7 version.

  * Validators
  * Serializers
  * Repositories
  * Interfaces
  * Relationships
  * Middleware_auth
    * JWT-go
  * Gin v1.7.7
  * Gorm v1.9.12
  * Godotenv v1.4.0

### Backend 2

The technology used for the server is [Laravel](https://laravel.com/) in its 9 version.

  * Migrations
  * Models
  * Controllers
  * Mysql
    * Relationships
    * Schema
  * Middleware_auth
    * Header
    * Token JWT


### Database

Server uses a [MySQL](https://www.mysql.com/) database in its 8.0 version.