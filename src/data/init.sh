#!/bin/sh
mongoimport -d repositories_db -c repositories --file data.json --jsonArray --uri "mongodb://mongo:27017"