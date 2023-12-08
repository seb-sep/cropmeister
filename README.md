# Cropmeister

GitHub repository: https://github.com/seb-sep/cropmeister
Note: This repo includes everything needed for the frotend, as we serve static HTML files with NGINX instead of using AppSmith.

Video link: https://share.descript.com/view/t3S1r1u8l11

## Layout

- `main.go`: The entry point for the API, sets up the DB connection and contains the main API root routes.
- `/static`: The folder containing all our frontend HTML/CSS/JS served by NGINX in our Docker network.
- `schema.sql`: The SQL file dictating our database schema, consumed by `sqlc` to make type-safe db objects.
- `query.sql`: The SQL file dictating our SQL queries, consumed by `sqlc` to make type-safe db queries.
- `/routes`: The folder containing all non-trivial endpoints, grouped by root endpointt (`farm.go` contains all `/farm/` endpoints)
- `/mockaroodata`: The folder containing our Mockaroo CSVs and SQL files which we used to populate the DB
- `/mysql`: The folder which the MySQL Docker image is mounted to to keep data persistent across `docker-compose` runs
- `/db`: The output of `sqlc` which contains typesafe db queries and schema objects

## How to Use
Run the container with `docker compose up`. This may take a while for all the images to download and build.
Then, navigate to http://localhost [here](http://localhost) to check out our frontend, and have fun!




