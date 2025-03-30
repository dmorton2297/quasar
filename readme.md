# Quasar
_A massive and extremely remote celestial object, emitting exceptionally large amounts of energy_

---

## Goal - create an extendable relational data store service that
1. Has relational database for storing data for a strongly typed entity
2. Generates messages for the following events.
    1. Event was deleted
        1. Metadata - ID and blob that was deleted
    2. Event was created - ID and blob that was created
    3. Event was updated - ID and updated blob vs original
    4. event was read - ID and data that was read

Message will eventually be consumed downstream by api-gateway, that will employ a server side events pattern for interfacing with clients

## Spin up db
```bash
docker-compose -f infra/db-compose.yml up -d    
```

## Start application
```bash
go run server.go
```

## Useful curls
1. Create md content
```bash
curl -X POST http://localhost:8080/md_content/create \  
     -H "Content-Type: application/json" \
     -d '{"name": "My Document", "content": "This is markdown content."}'
```

## Connet to db
```bash
psql postgresql://username:password@host:5432/db_name # replace username, password, host and db_name
```


