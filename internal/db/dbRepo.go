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
	log.Println("Trying to connect to DB.")
	host := "database"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "postgres"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", dsn)

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

func GetMetaOrder(metaOrderId string) (*models.MetaOrder, error) {
	conn, err := initializeDb()

	log.Println("Beginning get meta order now.")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	defer conn.Close()

	log.Println("Beginning query now.")

	query := `select meta_order_id, meta_order_type, status, exchange, crt_dtm, crt_usr_nm, last_udt_dtm, last_udt_usr_nm from kraken_meta_order.meta_order where meta_order_id = ` + metaOrderId

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

func FindMetaOrders(searchCriteria models.SearchCriteria) ([]*models.MetaOrder, error) {
	conn, err := initializeDb()

	log.Println("Beginning find meta order now.")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	defer conn.Close()

	query := "SELECT meta_order_id, meta_order_type, status, exchange, crt_dtm, crt_usr_nm, last_udt_dtm, last_udt_usr_nm FROM kraken_meta_order.meta_order WHERE status = $1"

	stmt, err := conn.Prepare(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating query: %v\n", err)
		return nil, err
	}

	log.Println(stmt)

	result := []*models.MetaOrder{}

	log.Println("Querying DB now.")

	rows, err := stmt.Query(searchCriteria.Status)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		m := &models.MetaOrder{}
		err = rows.Scan(&m.MetaOrderId, &m.MetaOrderType, &m.Status, &m.Exchange, &m.CreateDateTime, &m.CreateUserName, &m.LastUpdateDateTime, &m.LastUpdateUserName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse result failed: %v\n", err)
			return nil, err
		}
		result = append(result, m)
	}

	fmt.Println("Success.")
	return result, nil
}

func DeleteMetaOrder(metaOrderId string) (*models.MetaOrder, error) {
	log.Println("Beginning delete meta order now.")

	conn, err := initializeDb()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	defer conn.Close()

	log.Println("Beginning query now.")

	query := `update kraken_meta_order.meta_order set status='CANCELLED' where meta_order_id = ` + metaOrderId + ` returning *`

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

func initializeDb() (*sql.DB, error) {
	log.Println("Trying to connect to DB.")
	host := "database"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "postgres"

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}

	log.Println("Connected to DB.")

	return conn, nil
}
