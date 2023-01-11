#!/bin/bash
chown -R laravel:laravel /var/www
composer install
php artisan key:generate
php artisan route:list
php artisan vendor:publish --provider="PHPOpenSourceSaver\JWTAuth\Providers\LaravelServiceProvider"
# php artisan jwt:secret
php artisan migrate:status
php artisan migrate
php artisan serve --host=0.0.0.0 --port=6001
exec "$@";