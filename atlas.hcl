// Define an environment named "local"
env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["multi.hcl", "file.hcl"].
  src = "./db/schema.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = "postgres://itracker:stony_cyclable_adequacy@localhost:5430/itracker?search_path=public&sslmode=disable"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "docker://postgres/15"
}

variable "db_url" {
  type = string
  default = ""
}

env "fly" {
  src = "./db/schema.hcl"

  url = "${var.db_url}"
}