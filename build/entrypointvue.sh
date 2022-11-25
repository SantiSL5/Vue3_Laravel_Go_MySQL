#!/bin/bash
chown -R vue:vue /var/www
su vue
bash
npm install
npm run serve

exec "$@";