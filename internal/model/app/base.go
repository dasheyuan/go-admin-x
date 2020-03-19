package app

import (
	"fmt"
	"go-admin-x/internal/model/base"
)

func TableName(name string) string {
	return fmt.Sprintf("%s%s%s", base.GetTablePrefix(), "app_", name)
}
