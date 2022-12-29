// GraphQLのスキーマとO/Rマッパーの間を取り持つGo構造体
package entity

type User struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}



