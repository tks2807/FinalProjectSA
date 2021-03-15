package main

import (
	"Catalog_Admin/pkg"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connected!")
		return nil, err
	}
	return pool, nil
}

func main() {
	dsn := flag.String("dsn", "postgresql://localhost/catalog?user=postgres&password=admin", "PostGreSQL")
	flag.Parse()
	var err error
	pkg.Conn, err = openDB(*dsn)
	if err != nil{
		log.Fatalf("ERROR!", err)
	}
	router  := gin.Default()
	router.Run(":4001")
}
