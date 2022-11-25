#!/bin/bash
chown -R laravel:laravel /var/www
php artisan serve --host=0.0.0.0 --port=6001
exec "$@";