package aggregate

import (
	"go_dispatcher/entity"
)

// Customer 聚合包含了代表一个客户所需的所有实体
type Customer struct {
	// Person是客户的根实体
	// person.id是聚合的主标识符
	Person *entity.Person
}
