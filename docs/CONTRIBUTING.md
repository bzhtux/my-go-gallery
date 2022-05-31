# How to contribute

Great ! I'm so glad you want to contribite to this project. Here are some guidelines to help in contributing.

# Context and App goal

This App is supposed to be a sample app to test a fresh kubernetes installation. It will be a photos galery written in Golang composed by:

* a PostgreSQL database
* an API/Backend service
* a frontend (JS ninjas are welcome to help, otherwise it will be simple (very simple) HTML code)

# Requirements

Here is the list of what you need to contribute:

* Golang installed (of course) - I use `go version go1.17.6 darwin/amd64` on my laptop.
* Docker engine/desktop (20.10.14)
* docker image for postgres (latest) `docker pull postgres`

# GitFlow

The main working branch is `develop` but it shouldn't be used directly. I prefer to create a new branch per feature (e.g `feat/my-new-feature`). When the feature is done create a PR with `develop`.
If you have any question regarding this gitflow, feel free to ask.

# Tasks and Todo

All the tasks are in the projects tab ;-)

# Extras

To run a postgres DB using docker, create a `pg_data` directory first and then run the script below:

```shell
#!/usr/bin/env bash

set -euo pipefail

# DB design is taken from https://dev.to/amckean12/designing-a-relational-database-for-a-cookbook-4nj6



PG_PASS="<CHANGE_ME>"
GAL_USER="gallery"
GAL_USER_PASS="<CHANGE_ME>
GAL_DB="gallery"

if [ ! -d "./pg_data" ];then

    mkdir -p pg_data

fi

# Run postgres in a docker container
docker run \
    --name docker-pg-gal \
    -e POSTGRES_PASSWORD="${PG_PASS}" \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v "$PWD"/pg_data/:/var/lib/postgresql/data \
    -p 5432:5432 \
    -d postgres

# Wait postgres to be ready to accept connections
echo "Waiting for PostgreSQL ..."
sleep 60

    
# Create ROLE and PASS
docker exec docker-pg-gal psql -U postgres -c "CREATE USER ${GAL_USER};"
docker exec docker-pg-gal psql -U postgres -c "ALTER ROLE ${GAL_USER} WITH PASSWORD '${GAL_USER_PASS}';"
docker exec docker-pg-gal psql -U postgres -c "CREATE DATABASE ${GAL_DB} WITH OWNER ${GAL_USER};"
docker exec docker-pg-gal psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE ${GAL_DB} TO ${GAL_USER};"
docker exec docker-pg-gal psql -U postgres -c "ALTER ROLE ${GAL_USER} superuser;"
docker exec -e PGPASSWORD=${GAL_USER_PASS} docker-pg-gal psql -U "${GAL_USER}" "${GAL_DB}" -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
docker exec docker-pg-gal psql -U postgres -c "ALTER ROLE ${GAL_USER} nosuperuser;"
```

To clean your local DB:

```shell
#!/usr/bin/env bash

set -eu

echo -ne "* Stopping docker-pg-gal container: "
docker stop docker-pg-gal
echo -ne "\033[32mDONE\033[0m\n"

echo -ne "* Remove stopped docker-pg-gal container: "
docker rm -f docker-pg-gal
echo -ne "\033[32mDONE\033[0m\n"

echo -ne "* cleanup pgdata directory: "
rm -rf pg_data/pgdata/
echo -ne "\033[32mDONE\033[0m\n"
```

Let's work together now !
