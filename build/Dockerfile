FROM node:16.13.0 as vue

ARG UID

RUN yarn global add @vue/cli
RUN yarn global add @vue/cli-service
RUN yarn global add @vue/cli-init

WORKDIR /var/www

RUN deluser node
RUN useradd -G www-data,root -u $UID -d /home/vue vue

COPY entrypointvue.sh /entrypointvue.sh
ENTRYPOINT ["/bin/bash", "/entrypointvue.sh"]

ENV HOST 0.0.0.0

#########################################################################


FROM golang:1.17.3 as gingonic
ARG UID

RUN mkdir -p /go/src/api

WORKDIR /go/src/api

RUN go get -u github.com/gin-gonic/gin
RUN go install github.com/acoshift/goreload@master

RUN useradd -G www-data,root -u $UID -d /home/go go && \
    mkdir -p /home/go && \
    chown -R go:go /home/go

COPY entrypointgo.sh /entrypointgo.sh

ENTRYPOINT ["/bin/bash", "/entrypointgo.sh"]

#########################################################################

# Use image which contains apache with php
FROM php:8.0 as laravel

ARG UID

# Install system dependencies
RUN apt-get update && apt-get install -y \
    git \
    curl \
    libpng-dev \
    libonig-dev \
    libxml2-dev \
    zip \
    unzip

# Clear cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Install PHP extensions
RUN docker-php-ext-install pdo_mysql mbstring exif pcntl bcmath gd

# Get latest Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Create system user to run Composer and Artisan Commands
RUN useradd -G www-data,root -u $UID -d /home/laravel laravel
RUN mkdir -p /home/laravel/.composer && \
    chown -R laravel:laravel /home/laravel

COPY entrypointlaravel.sh /entrypointlaravel.sh

# USER laravel

# Set working directory
WORKDIR /var/www

ENTRYPOINT ["/bin/bash", "/entrypointlaravel.sh"]

#################################################################################################


