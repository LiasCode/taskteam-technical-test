package container

import (
	"company-microservice/api/database"
	"database/sql"
)

func (c *Container) GetPostgresqlDB() *sql.DB {
	if c.postgresqlDB != nil {
		return c.postgresqlDB
	}

	db, err := database.NewPostgresqlConnection()

	if err != nil {
		panic(err)
	}

	c.postgresqlDB = db

	return c.postgresqlDB
}
