package requester

/*

	NEED CONNECT TO DB

*/

/*


import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"service_history/pkg/config"
	"service_history/pkg/store/postgres"
	"testing"
	"time"
)

type DBMock struct {
}

func (D DBMock) Open() error {
	return nil
}

func (D DBMock) GetConn() *sql.DB {
	return nil
}

func (D DBMock) Close() error {
	fmt.Println("close db")
	return nil
}

func TestNew(t *testing.T) {
	assert.NotNil(t, New(14))
	assert.Nil(t, New(0))
	assert.Nil(t, New(-1))
}

func TestConnect(t *testing.T) {
	assert.Nil(t, godotenv.Load("../../../.env"))
	req := requester{period: 5 * time.Second}

	assert.Nil(t, req.Start(DBMock{}.GetConn()))

}

func Test_GetPair(t *testing.T) {
	assert.Nil(t, godotenv.Load("../../../.env"))
	assert.Nil(t, os.Setenv("DB_HOST", "localhost"))
	conf, err := config.NewConfig("../../../config/config.json")
	assert.Nil(t, os.Setenv("GENERATE_PORT", ":8080"))
	assert.Nil(t, os.Setenv("NETWORK_NAME", "localhost"))

	assert.Nil(t, err)
	prov := postgres.NewProvider(conf.GetSQL())
	assert.Nil(t, prov.Open())

	conn := prov.GetConn()
	defer conn.Close()
	pair, err := getPair(conn)
	assert.Nil(t, err)
	log.Println(pair)
	assert.Nil(t, distributor(pair, conn))

	time.Sleep(5 * time.Second)
}
*/
