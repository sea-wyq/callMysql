package data

import (
	"context"
	"fmt"

	"callMysql/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type MysqlController struct {
	data *Data
	log  *log.Helper
}

// NewDataController .
func NewMysqlController(data *Data, logger log.Logger) biz.DataController {
	return &MysqlController{
		data: data,
		log:  log.NewHelper(logger),
	}

}

func (r *MysqlController) FindByID(ctx context.Context, num int64) (*biz.Mess, error) {
	rows, err := r.data.db.Query("SELECT * FROM test WHERE id = ?", num)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// 处理查询结果
	var id int64
	var name string
	result := ""
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		result += fmt.Sprintf("ID: %d, Name: %s\n", id, name)
	}
	return &biz.Mess{Info: result}, nil
}
