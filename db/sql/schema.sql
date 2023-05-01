-- Create "symbols" table
CREATE TABLE "symbols" ("id" character varying(50) NOT NULL, "type" character varying(50) NOT NULL, "name" character varying(50) NULL, "exchange" character varying(50) NULL, PRIMARY KEY ("id"));
-- Create "stock_data" table
CREATE TABLE "stock_data" ("time" timestamptz NOT NULL, "symbol_id" character varying(50) NOT NULL, "open" numeric(20,5) NULL, "high" numeric(20,5) NULL, "low" numeric(20,5) NULL, "close" numeric(20,5) NULL, "adjusted_close" numeric(20,5) NULL, "volume" bigint NULL, "dividend_amount" numeric(20,5) NULL, "split_coefficient" double precision NULL, CONSTRAINT "symbol_id" FOREIGN KEY ("symbol_id") REFERENCES "symbols" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create index "time_symbol_idx" to table: "stock_data"
CREATE UNIQUE INDEX "time_symbol_idx" ON "stock_data" ("symbol_id", "time");
