package requester

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"service_history/internal/app/contract"
	"strings"
	"time"
)

var (
	requestCurrency = "select * from currency_pair"
	insertStart     = "insert into history_deal (pair_id, deal_time, coast) values"
	insertPattern   = "(%d, '%s', %f)"
	insertEnd       = "on conflict do nothing;"
)

type requester struct {
	period time.Duration
}

func New(p int) contract.IRequester {
	r := &requester{
		period: time.Duration(p) * time.Second,
	}
	if p <= 0 {
		return nil
	}
	return r
}

func (r *requester) Start(db *sql.DB) (err error) {
	go func() {
		{
			for {
				<-time.After(r.period)
				if err = checkConnection(); err != nil {
					return
				}
				pairValues, err := getPair(db)
				if err != nil {
					return
				}
				if err = distributor(pairValues, db); err != nil {
					return
				}
			}
		}
	}()
	return err
}

// checkConnection with generate_service
func checkConnection() error {
	addr := strings.Join([]string{"http://", os.Getenv("NETWORK_NAME"), os.Getenv("GENERATE_PORT"), "/health"}, "")
	_, err := http.Get(addr)
	if err != nil {
		log.Println("Check connection err:", err)
	}
	return err
}

type currencyPair struct {
	id   int
	name string
}

// getPair retrieves currency name valutes from the database
func getPair(db *sql.DB) ([]currencyPair, error) {
	rows, err := db.Query(requestCurrency)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var result []currencyPair
	for rows.Next() {
		var temp currencyPair
		err = rows.Scan(&temp.id, &temp.name)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, temp)
	}

	return result, err
}

// distributor starts work for each currency name
func distributor(pairs []currencyPair, db *sql.DB) (err error) {
	for _, p := range pairs {
		go func(curr currencyPair) {
			err = worker(curr, db)
		}(p)
	}
	return err
}

// worker GET-request to generate_service and unmarshall response
func worker(c currencyPair, db *sql.DB) error {
	addr := strings.Join([]string{"http://", os.Getenv("NETWORK_NAME"), os.Getenv("GENERATE_PORT"), "/values?target=", c.name}, "")
	resp, err := http.Get(addr)
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	interResult := map[string]map[string]map[string]float32{}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(body, &interResult)
	if err != nil {
		log.Println(err)
		return err
	}

	result := interResult[c.name]["Rate"]
	values100 := mapToString(c.id, result)

	query := strings.Join([]string{insertStart, values100, insertEnd}, " ")

	insertDB(query, db)

	return err
}

// mapToString all values concatenates in one string
func mapToString(id int, val map[string]float32) string {
	var temp []string
	for date, value := range val {
		temp = append(temp, fmt.Sprintf(insertPattern, id, date, value))
	}
	return strings.Join(temp, ",")
}

// insertDB push to DB all values with transaction
func insertDB(req string, db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return
	}

	_, err = tx.Exec(req)
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		_ = tx.Rollback()
	}

}
