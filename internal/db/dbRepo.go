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

	log.Println("Beginning query now.")

	// query := `insert into kraken_meta_order.meta_order (meta_order_id, meta_order_type, status)
	// values (nextval('kraken_meta_order.meta_order_id'),  '` + metaOrder.MetaOrderType + `', 'ACTIVE') RETURNING *`

	query := `insert into kraken_meta_order.meta_order (meta_order_id, meta_order_type, status, exchange) 
	values (nextval('kraken_meta_order.meta_order_id'), '` + metaOrder.MetaOrderType + `', 'ACTIVE', '` + metaOrder.Exchange + `') RETURNING *`

	log.Println(query)

	m := &models.MetaOrder{}

	err = conn.QueryRow(query).Scan(&m.MetaOrderId, &m.MetaOrderType, &m.Status, &m.Exchange, &m.CreateDateTime, &m.CreateUserName, &m.LastUpdateDateTime, &m.LastUpdateUserName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	fmt.Println("Success.")
	return m, nil
}
