#!/bin/bash
# Restore from the mounted /dump folder
# We use the env vars provided in compose to authenticate
mongorestore --username "$MONGO_INITDB_ROOT_USERNAME" --password "$MONGO_INITDB_ROOT_PASSWORD" --authenticationDatabase admin --archive=openfoodfacts-mongodbdump.gz --gzip
