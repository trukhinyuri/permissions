package model

import (
	"time"

	"github.com/go-pg/pg/orm"
	"github.com/pkg/errors"
)

type Resource struct {
	ID          string     `sql:"id,pk,type:uuid,default:uuid_generate_v4()"`
	CreateTime  time.Time  `sql:"create_time,default:now(),type:TIMESTAMPTZ,notnull"`
	Deleted     bool       `sql:"deleted,notnull"`
	DeleteTime  *time.Time `sql:"delete_time,type:TIMESTAMPTZ"`
	TariffID    string     `sql:"tariff_id,type:uuid,notnull"`
	OwnerUserID string     `sql:"owner_user_id,type:uuid,notnull"`
	Label       string     `sql:"label,notnull"`
}

func (r *Resource) BeforeDelete(db orm.DB) error {
	// do not allow delete from app
	return errors.New("record delete not allowed, use update set deleted = true")
}

func (r *Resource) BeforeUpdate(db orm.DB) error {
	if r.Deleted {
		now := time.Now()
		r.DeleteTime = &now
	}
	return nil
}
