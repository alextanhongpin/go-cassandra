# go-cassandra

Basic operations with Cassandra in Golang.

## Quickstart

```bash
# Spins up dockerized Cassandra
$ make up

# Access cqlsh to run queries. Paste the content of schema.cql into the bash to create the keyspace and tables required.
$ make sh

# Execute the queries. Change the `createUser` boolean to true to create a new user.
$ make start

# Spins down the docker containers.
$ make down
```

## Thoughts

- how to run migrations for Cassandra queries?
- how to handle collection queries?
- how to handle User Defined Types (UDT)?
