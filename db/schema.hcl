schema "public" {
}

table "sensors" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "type" {
    null = true
    type = varchar(50)
  }
  column "location" {
    null = true
    type = varchar(50)
  }
  primary_key {
    columns = [column.id]
  }
}

table "sensor_data" {
  schema = schema.public
  column "time" {
    null = false
    type = timestamptz
  }
  column "sensor_id" {
    null = true
    type = integer
  }
  column "temperature" {
    null = true
    type = double_precision
  }
  column "cpu" {
    null = true
    type = double_precision
  }
  foreign_key "sensor_id" {
    columns     = [column.sensor_id]
    ref_columns = [table.sensors.column.id]
  }
}
