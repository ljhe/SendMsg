package db

const (
	maxConn     = 25
	maxIdle     = 10
	maxLifeTime = 600
	maxIdleTime = 300
)

var dbConfig = []struct {
	host         string
	userName     string
	passWord     string
	port         int
	dataBaseName string
}{
	{
		host:         "127.0.0.1",
		userName:     "root",
		passWord:     "123456",
		port:         3306,
		dataBaseName: "test1",
	}, {
		host:         "127.0.0.1",
		userName:     "root",
		passWord:     "123456",
		port:         3306,
		dataBaseName: "test2",
	},
}
