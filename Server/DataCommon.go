package Server

import (
	"database/sql"
	"errors"
	"reflect"
)

func ScanToStruct(row *sql.Rows, dest interface{}) error {
	//取得資料庫colums名稱
	col_names, err := row.Columns()
	er.CheckErr(err)

	v := reflect.ValueOf(dest)
	if v.Elem().Type().Kind() != reflect.Struct {
		return errors.New("it is not Struct")
	}
	//interface slice
	scan_dest := []interface{}{}
	// 建立一個string, interface{}的map
	addr_by_col_name := map[string]interface{}{}

	for i := 0; i < v.Elem().NumField(); i++ {
		propertyName := v.Elem().Field(i)
		col_name := v.Elem().Type().Field(i).Name
		if col_name == "" {
			if v.Elem().Field(i).CanInterface() == false {
				continue
			}
			col_name = propertyName.Type().Name()
		}
		// Addr() 返回該屬性的記憶體位置的指針
		// Interface() 返回該屬性真正的值, 這裡還是存著位置
		addr_by_col_name[col_name] = propertyName.Addr().Interface()
	}
	// 把實際各成員屬性的位置, 給加到scan_dest中
	for _, col_name := range col_names {
		scan_dest = append(scan_dest, addr_by_col_name[col_name])
	}
	// 執行Scan
	return row.Scan(scan_dest...)
}
