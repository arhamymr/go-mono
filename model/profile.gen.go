// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProfile = "Profile"

// Profile mapped from table <Profile>
type Profile struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Bio    string `gorm:"column:bio;not null" json:"bio"`
	UserID int32  `gorm:"column:user_id;not null" json:"user_id"`
}

// TableName Profile's table name
func (*Profile) TableName() string {
	return TableNameProfile
}
