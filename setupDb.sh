docker run -p 5432:5432 --name lilvue-postgres -e POSTGRES_PASSWORD=test -d postgres

# Provide password so we don't need to do so in prompt 
export PGPASSWORD=test

# Connect
# psql -h 127.0.0.1 -p 5432 -U postgres 

# Create DB
# psql -h 127.0.0.1 -U postgres postgres -f createDb.sql

# Init schema
# psql -h 127.0.0.1 -U postgres postgres -d lilvue -f schema.sql