package model

func NewUser(id int64) *User {
	u := &User{}
	u.Id = id
	u.sharding = &ModHashSharding{
		ShardingVal:   id,
		ShardingDb:    1,
		ShardingTable: 5,
	}
	return u
}

type User struct {
	Id       int64  `gorm:"column:id;primary_key" json:"userId"`
	Name     string `gorm:"column:name" json:"name"`
	sharding Sharding
}

func (u *User) dbPrefix() string {
	return "db_user_"
}

func (u *User) dbSuffix() string {
	return u.sharding.GetDbSuffix()
}

func (u *User) Db() string {
	return u.dbPrefix() + u.dbSuffix()
}

func (u *User) tablePrefix() string {
	return "user_"
}

func (u *User) tableSuffix() string {
	return u.sharding.GetTableSuffix()
}

func (u *User) TableName() string {
	return u.tablePrefix() + u.tableSuffix()
}

func NewUserLog(id int64) *UserLog {
	ul := &UserLog{}
	ul.Id = id
	ul.sharding = &MonthSharding{}
	return ul
}

type UserLog struct {
	Id       int64  `gorm:"column:id;primary_key" json:"id"`
	Log      string `gorm:"column:log" json:"log"`
	sharding Sharding
}

func (ul *UserLog) dbPrefix() string {
	return "user_log"
}

func (ul *UserLog) dbSuffix() string {
	return ul.sharding.GetDbSuffix()
}

func (ul *UserLog) Db() string {
	return ul.dbPrefix() + ul.dbSuffix()
}

func (u *UserLog) tablePrefix() string {
	return "user_log_"
}

func (u *UserLog) tableSuffix() string {
	return u.sharding.GetTableSuffix()
}

func (u *UserLog) TableName() string {
	return u.tablePrefix() + u.tableSuffix()
}
