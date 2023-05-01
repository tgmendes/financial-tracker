 atlas schema inspect --env local --format "{{ sql . }}" > db/sql/schema.sql
 sqlc generate --experimental