package service

/*
func TestNewHistoryServer(t *testing.T) {
	assert.Nil(t, godotenv.Load("../../../.env"))
	c, err := config.NewConfig("../../../config/config.json")
	assert.Nil(t, err)

	prov := postgres.NewProvider(c.GetSQL())
	assert.Nil(t, prov.Open())

	db := prov.GetConn()

	t1, err := time.Parse("2006-02-01 15:04:05", "2022-01-01 12:49:34")
	t2, err := time.Parse("2006-02-01 15:04:05", "2023-06-12 12:49:34.00")
	pair := "USDRUB"
	rows, err := db.Query(query, pair, t1, t2)
	assert.Nil(t, err)

	for rows.Next() {
		var coast float32
		var timeReq time.Time
		assert.Nil(t, rows.Scan(&timeReq, &coast))
		log.Println(timeReq, coast)
	}

}
*/
