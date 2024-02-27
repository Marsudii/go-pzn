package database

var connection = ""

func init() {
	connection = "MYsql"
}

func GetDatabase() string {
	return connection
}
