#! /bin/sh

# these vars are not available at buildtime so we need to import them at runtime
export COMMENTO_PORT=$PORT
export COMMENTO_POSTGRES=$DATABASE_URL

./commento
