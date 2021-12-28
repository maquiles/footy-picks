FROM postgres

ENV POSTGRES_PASSWORD docker
ENV POSTGRES_DB footy_picks

COPY *.sql /docker-entrypoint-initdb.d/

# docker build -t footy-picks-db .
# docker run --name footy-picks-db -p 5432:5432 footy-picks-db