#!/bin/bash
# Restore from the mounted /dump folder
# We use the env vars provided in compose to authenticate
mongorestore --username "$MONGO_INITDB_ROOT_USERNAME" --password "$MONGO_INITDB_ROOT_PASSWORD" --authenticationDatabase admin --archive=openfoodfacts-mongodbdump.gz --gzip
mongosh --username "$MONGO_INITDB_ROOT_USERNAME" --password "$MONGO_INITDB_ROOT_PASSWORD" --authenticationDatabase admin "mongodb://localhost:27017/off" --eval 'db.products.dropIndexes()'
mongosh --username "$MONGO_INITDB_ROOT_USERNAME" --password "$MONGO_INITDB_ROOT_PASSWORD" --authenticationDatabase admin "mongodb://localhost:27017/off" --eval 'db.products.createIndex({ _keywords: "text", product_name: "text" }, { name: "TextSearchIndex" });'
