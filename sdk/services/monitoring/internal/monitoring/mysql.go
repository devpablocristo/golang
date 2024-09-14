package monitoring

import (
	"context"

	portsmysql "github.com/devpablocristo/golang/sdk/pkg/databases/sql/mysql/go-sql-driver/ports"
	ports "github.com/devpablocristo/golang/sdk/services/monitoring/internal/monitoring/ports"
)

type mysqlRepository struct {
	mysql portsmysql.Repository
}

func NewMySqlRepository(db portsmysql.Repository) ports.Repository {
	return &mysqlRepository{
		mysql: db,
	}
}

func (r *mysqlRepository) CheckDbConn(ctx context.Context) error {
	return r.mysql.DB().PingContext(ctx)
}
