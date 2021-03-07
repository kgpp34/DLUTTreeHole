package main

import (
	"TreeHole/internal/handler"
	"TreeHole/internal/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/hako/branca"
	_ "github.com/jackc/pgx/v4/stdlib"
)


const (
	databaseURL = "postgresql://root@V_SXIA-NB2:26257//treehole?sslmode=disable"
	port        = 8083
)

func main() {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("could not open db connection: %v\n", err)
		return
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("could not ping to db: %v\n", err)
		return 
	}

	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")
	codec.SetTTL(uint32(service.TokenLifeSpan.Seconds()))
	s := service.New(db, codec)
	h := handler.New(s)
	addr := fmt.Sprintf(":%d", port)
	log.Printf("accepting connect on port %d", port)
	if err = http.ListenAndServe(addr, h); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}

