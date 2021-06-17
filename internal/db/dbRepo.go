package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"

	"github.com/TheOtherDavid/kraken-meta-order/internal/models"
)

func CreateMetaOrder(metaOrder models.MetaOrder) (*models.MetaOrder, error) {
	//Do stuff here
	log.Println("Trying to connect to DB.")
	host := "database"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "postgres"
	//conn, err := pgx.Connect(context.Background(), "postgresql://postgres:postgres@postgres:5432/postgres")
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", dsn)

	//I guess we're going to have to open a connection, and then connect to the right table, and then execute a query.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	log.Println("Connected to DB.")

	defer conn.Close()

	var name string
	var weight int64
	log.Println("Beginning query.")

	err = conn.QueryRow("insert into kraken_meta_order.meta_order (meta_order_type) values ("+metaOrder.MetaOrderType+")", 42).Scan(&name, &weight)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	fmt.Println("Success.")
	return nil, nil
}
