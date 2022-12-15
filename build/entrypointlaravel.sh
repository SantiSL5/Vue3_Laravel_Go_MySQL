#!/bin/bash
chown -R laravel:laravel /var/www
composer install
php artisan migrate:status
php artisan migrate
php artisan serve --host=0.0.0.0 --port=6001
exec "$@";