package fun

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	basic "uminer/common/basic"
	"uminer/common/sql"

	commapi "uminer/common/api/v1"

	"gorm.io/gorm"
)

type Binds []*Bind

type Bind struct {
	Platform string `json:"platform"`
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func (r Binds) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Binds) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

type ResourcePools []string

func (r ResourcePools) Value() (driver.Value, error) {
	return sql.Value(r)
}

func (r *ResourcePools) Scan(input interface{}) error {
	return sql.Scan(r, input)
}

type Permission commapi.UserPermission

func (r Permission) Value() (driver.Value, error) {
	return sql.Value(r)
}

func (r *Permission) Scan(input interface{}) error {
	return sql.Scan(r, input)
}

type User struct {
	basic.Model
	Id            string        `gorm:"type:varchar(100);not null;primaryKey;comment:'用户ID'"`
	FullName      string        `gorm:"type:varchar(100);not null;default:'';index;comment:'姓名'"`
	Gender        int32         `gorm:"type:int;not null;default:0;comment:'性别：1.男,2.女'"`
	Email         string        `gorm:"type:varchar(100);not null;default:'';index;comment:'用户邮箱'"`
	Phone         string        `gorm:"type:varchar(100);not null;default:'';index;comment:'电话号码'"`
	Password      string        `gorm:"type:varchar(100);not null;default:'';comment:'密码'"`
	Status        int32         `gorm:"type:int;not null;default:0;comment:'性别：1.冻结,2.正常'"`
	Bind          Binds         `gorm:"type:json;comment:'第三方账号绑定信息'"`
	ResourcePools ResourcePools `gorm:"type:json;comment:'资源池'"`
	Desc          string        `gorm:"type:varchar(100);default:'';index;comment:'备注'"`
	Permission    *Permission   `gorm:"type:json;comment:权限"`
	MinioUserName string        `gorm:"type:varchar(100);uniqueIndex:minioUserName;comment:'minio用户名'"`
	Buckets       sql.Strings   `gorm:"type:json;comment:'允许访问的buckets'"`
}

func (User) TableName() string {
	return "user"
}

// ************** params ****************

type UserList struct {
	SortBy    string
	OrderBy   string
	PageIndex uint32
	PageSize  uint32
	FullName  string
	Gender    int32
	Email     string
	Phone     string
	SearchKey string
	Status    int32
	Bind      Binds
	Desc      string
}

func (u UserList) Where(db *gorm.DB) *gorm.DB {
	querySql := "1 = 1"
	params := make([]interface{}, 0)

	if u.FullName != "" {
		querySql += " and full_name = ? "
		params = append(params, u.FullName)
	}

	if u.Email != "" {
		querySql += " and email = ? "
		params = append(params, u.Email)
	}

	if u.Phone != "" {
		querySql += " and phone = ? "
		params = append(params, u.Phone)
	}

	if u.Gender != 0 {
		querySql += " and gender = ? "
		params = append(params, u.Gender)
	}

	if u.Status != 0 {
		querySql += " and status = ? "
		params = append(params, u.Status)
	}

	return db.Where(querySql, params...)
}

func (u UserList) Or(db *gorm.DB) *gorm.DB {
	if u.SearchKey != "" {
		searchKeyLike := "%" + u.SearchKey + "%"
		db = db.Where("full_name like ? or email like ? or bind->'$[*].userName' like ? or `desc` like ?", searchKeyLike, searchKeyLike, searchKeyLike, searchKeyLike)
	}
	return db
}

func (u UserList) Order(db *gorm.DB) *gorm.DB {
	var orderBy, sortBy string
	if u.OrderBy != "" {
		orderBy = u.OrderBy
	} else {
		orderBy = "desc"
	}
	if u.SortBy != "" {
		sortBy = u.SortBy
	} else {
		sortBy = "created_at"
	}

	db = db.Order(fmt.Sprintf("%v %v", sortBy, orderBy))
	return db
}

func (u UserList) Pagination(db *gorm.DB) *gorm.DB {
	var pageIndex, pageSize int
	if u.PageIndex <= 0 {
		pageIndex = 1
	} else {
		pageIndex = int(u.PageIndex)
	}
	if u.PageSize <= 0 {
		pageSize = 10
	} else {
		pageSize = int(u.PageSize)
	}
	db = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	return db
}

type UserQuery struct {
	Id    string
	Email string
	Phone string
	Bind  *Bind
	Desc  string
}

type UserAdd struct {
	Id            string
	FullName      string
	Gender        int32
	Email         string
	Phone         string
	Password      string
	Status        int32
	Bind          *Bind
	ResourcePools []string
	Desc          string
}

type UserUpdate struct {
	FullName      string
	Gender        int32
	Email         string
	Phone         string
	Password      string
	Status        int32
	Bind          Binds
	ResourcePools []string
	Desc          string
}

type UserUpdateCond struct {
	Id    string
	Email string
	Phone string
}

type UserListIn struct {
	Ids []string
}

type UserConfig struct {
	basic.Model
	UserId string  `gorm:"primaryKey;type:varchar(100);not null;default:'';comment:用户id"`
	Config sql.Map `gorm:"type:json;comment:配置"`
}

func (UserConfig) TableName() string {
	return "user_config"
}
