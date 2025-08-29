package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-demo/internal/bondProductApi/biz"
	"kratos-demo/internal/bondProductApi/data/ent"
	"kratos-demo/internal/bondProductApi/data/grom/repo"
	"kratos-demo/internal/config"
	gormtool "kratos-demo/pkg/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRepoRegistry,
)

var GormDataProviderSet = wire.NewSet(
	repo.NewGormProductRepo, // 实现 biz.ProductRepo
	NewGormDB,
	NewGormSess,
	NewGormTxManager,
)

var EntDataProviderSet = wire.NewSet(
	ent.NewEntProductRepo, // 实现 biz.ProductRepo
)

func NewGormSess(orm *gormtool.Orm) gormtool.Sess {
	return orm
}

func NewGormTxManager(orm *gormtool.Orm) biz.TxManager {
	return orm
}

// Data .
type Data struct {
}

// NewData .
// 统一管理 db redis mq的close
func NewData(gormDB *gormtool.Orm, logger log.Logger) (*Data, func(), error) {
	// 聚合 cleanup
	cleanup := func() {
		log.Info("closing data resources")
		/*for _, c := range cleanupList {
			c()
		}*/
		if gormDB != nil {
			gormDB.Tx()
		}
	}

	return &Data{}, cleanup, nil
}

func NewGormDB(cfg *config.Data) *gormtool.Orm {
	db, err := gormtool.InitMySQLDB(&gormtool.MySQLCfg{
		User:     "cfg.Database.User",
		Password: "cfg.Database.User",
		Host:     "cfg.Database.User",
		Port:     3366,
		Schema:   "cfg.Database.User",
	}, &gormtool.DBConfig{
		MaxIdleConn:     6,
		MaxOpenConn:     20,
		ConnMaxLifetime: 60,
	})
	if err != nil {
		panic(err)
	}

	return gormtool.NewORM(db)
}
