// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCategoryToPost = "_CategoryToPost"

// CategoryToPost mapped from table <_CategoryToPost>
type CategoryToPost struct {
	A int32 `gorm:"column:A;primaryKey" json:"A"`
	B int32 `gorm:"column:B;primaryKey" json:"B"`
}

// TableName CategoryToPost's table name
func (*CategoryToPost) TableName() string {
	return TableNameCategoryToPost
}
