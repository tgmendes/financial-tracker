schema "public" {
}

table "symbols" {
  schema = schema.public
  column "id" {
    null = false
    type = varchar(50)
  }
  column "type" {
    null = false
    type = varchar(50)
  }
  column "name" {
    null = true
    type = varchar(50)
  }
  column "exchange" {
    null = true
    type = varchar(50)
  }
  primary_key {
    columns = [column.id]
  }
}

table "stock_data" {
  schema = schema.public
  column "time" {
    null = false
    type = timestamptz
  }
  column "symbol_id" {
    null = false
    type = varchar(50)
  }
  column "open" {
    null = true
    type = decimal(20,5)
  }
  column "high" {
    null = true
    type = decimal(20,5)
  }
  column "low" {
    null = true
    type = decimal(20,5)
  }
  column "close" {
    null = true
    type = decimal(20,5)
  }
  column "adjusted_close" {
    null = true
    type = decimal(20,5)
  }
  column "volume" {
    null = true
    type = bigint
  }
  column "dividend_amount" {
    null = true
    type = decimal(20,5)
  }
  column "split_coefficient" {
    null = true
    type = double_precision
  }
  foreign_key "symbol_id" {
    columns     = [column.symbol_id]
    ref_columns = [table.symbols.column.id]
  }
  index "time_symbol_idx" {
    columns = [
      column.symbol_id,
      column.time
    ]
    unique = true
  }
}
