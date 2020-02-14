package driver

import (
  "database/sql"
  "log"
  "os"
  "github.com/lib/pq"
)


func ConnectDB() *sqlDB{
pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
logFatal(err)

db, err = sql.Open("postgres", pgUrl)
logFatal(err)

err = db.Ping()
logFatal(err)

return db
}
