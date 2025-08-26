package db

import (
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrorTrans = errors.New("transaction error")
)

type Orm struct {
	tx *gorm.DB
	db *gorm.DB
}

type Sess interface {
	Tx() *gorm.DB
	SetContext(ctx context.Context)
	Transaction(f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error
}

func NewORM(db *gorm.DB) *Orm {
	return &Orm{
		db: db,
	}
}

func NewORMWithLog(db *gorm.DB) *Orm {
	return &Orm{
		db: db,
	}
}

func (t *Orm) BeginTx(opts ...*sql.TxOptions) error {
	t.tx = t.db.Begin(opts...)
	if t.tx.Error != nil {
		//t.Errorf("begin transaction error:%v", t.tx.Error)
	}
	return t.tx.Error
}

func (t *Orm) CommitTx() error {
	if t.tx == nil {
		return ErrorTrans
	}
	er := t.tx.Commit().Error
	if er != nil {
		//t.Errorf("commit transaction error:%v", er)
	} else {
		t.tx = nil
	}
	return er
}

func (t *Orm) Tx() *gorm.DB {
	if t.tx != nil {
		return t.tx
	}
	return t.db
}

func (t *Orm) RollbackTx() {
	if t.tx == nil {
		return
	}
	t.tx.Rollback()
	t.tx = nil
}

func (t *Orm) SavePoint(sp string) error {
	if t.tx == nil {
		return nil
	}
	er := t.tx.SavePoint(sp).Error
	if er != nil {
		//t.Errorf("save point error:%v", er)
	}
	return er
}

func (t *Orm) Transaction(f func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	if t.tx != nil {
		return ErrorTrans
	}
	er := t.db.Transaction(f, opts...)
	if er != nil {
		//t.Errorf("transaction error:%v", er)
	}
	return er
}

// should call before start transaction
func (t *Orm) SetContext(ctx context.Context) {
	if t.tx != nil {
		t.tx = t.tx.WithContext(ctx)
		return
	}
	t.db = t.db.WithContext(ctx)
}
