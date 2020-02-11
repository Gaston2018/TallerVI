package main

import (
    "github.com/eaigner/jet"
    "github.com/lib/pq"
    "log"
    "os"
)

func logFatal(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
func main() {
    //Make sure you setup the ELEPHANTSQL_URL to be a uri, e.g. 'postgres://user:pass@host/db?options'
    pgUrl, err := pq.ParseURL(os.Getenv("postgres://hywmgvvs:uqZw6sHPpFvRA-kCtPGtU_fmK10KjL9x@rajje.db.elephantsql.com:5432/hywmgvvs"))
    logFatal(err)
    db, err := jet.Open("postgres", pgUrl)
    logFatal(err)
    var people []*struct {
        Id        int
        FirstName string
        LastName  string
    }
    err = db.Query("SELECT * FROM usuarios").Rows(&people)
    logFatal(err)
    for _, person := range people {
        log.Printf("Id: %v, First Name: %s, Last Name: %s",
            person.Id,
            person.FirstName,
            person.LastName)
    }
}
