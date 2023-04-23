 atlas schema inspect --env local --format "{{ sql . }}" > db/schema.sql
 sqlc generate --experimental