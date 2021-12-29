FROM postgres

ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB footy_picks

COPY *.sql /docker-entrypoint-initdb.d/

# docker build -t footy-picks-db .
# docker run -d --name footy-picks-db -p 5432:5432 footy-picks-db
# psql -h 127.0.0.1 -p 5432 -U postgres -d footy_picks