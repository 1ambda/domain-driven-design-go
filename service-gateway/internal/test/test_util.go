package test

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/config"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/order"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/product"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	"github.com/jinzhu/gorm"
	"github.com/ory/dockertest"
)

func PrepareDatabase() (*dockertest.Pool, *dockertest.Resource, string) {
	var db *sql.DB
	var err error
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	env := config.Env

	// pulls an image, creates a container based on it and runs it
	dockerEnvs := []string{
		"MYSQL_ROOT_PASSWORD=" + env.MysqlPassword,
		"MYSQL_DATABASE=" + env.MysqlDatabase,
	}

	resource, err := pool.Run("mysql", "5.7", dockerEnvs)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostPort := resource.GetPort(fmt.Sprintf("%s/tcp", env.MysqlPort))
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		env.MysqlUserName, env.MysqlPassword, env.MysqlHost, hostPort, env.MysqlDatabase)

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("mysql", conn)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return pool, resource, hostPort
}

type DatabaseContainer struct {
	DB       *gorm.DB
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func GetDatabaseContainer() DatabaseContainer {
	pool, resource, dbHostPort := PrepareDatabase()
	config.Env.MysqlPort = dbHostPort
	db := config.GetDatabase(MigrateCallback)

	return DatabaseContainer{db, pool, resource}
}

func (c *DatabaseContainer) Close() {
	c.pool.Purge(c.resource)
}

func GetTestDatabase(debug bool) *gorm.DB {
	config.Env.Mode = "TEST"
	config.Env.EnableDebugSQL = debug

	return config.GetDatabase(MigrateCallback)
}

func MigrateCallback(db *gorm.DB) {
	// Migrate only when using SQLite
	// for MySQL, will use flyway
	db.Set("gorm:table_options", "").AutoMigrate(
		&user.User{},
		&user.AuthIdentity{},
		&product.Category{},
		&product.Image{},
		&product.Product{},
		&order.Order{},
		&order.OrderDetail{},
	)
}
