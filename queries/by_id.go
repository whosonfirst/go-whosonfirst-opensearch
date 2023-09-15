package queries

import (
	"fmt"
)

func GetById(id int64) string {
	return fmt.Sprintf(`{"query": { "ids": { "values": [ %d ] } } }`, id)
}
