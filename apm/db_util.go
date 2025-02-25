package apm

import "database/sql"

type dBUtil struct{}

// DBUtil 是 dBUtil 类型的全局实例，提供数据库查询工具函数
var DBUtil = &dBUtil{}

// Query 方法用于从数据库查询结果中提取数据并返回结果。
// rows: 代表查询结果的 sql.Rows 对象。
// err: 查询过程中发生的错误，如果为 nil，则表示查询成功。
// 返回：一个包含多个记录的切片，每个记录是一个 map，key 是列名，value 是列的值。
// 如果查询过程中出现错误或没有数据，则返回 nil 或空切片。
func (d *dBUtil) Query(rows *sql.Rows, err interface{}) []map[string]interface{} {
	// 如果查询发生错误，则直接返回 nil
	if err != nil {
		return nil
	}
	// 如果 rows 为 nil，表示没有结果，返回空切片
	if rows == nil {
		return []map[string]interface{}{}
	}
	defer rows.Close() // 确保在函数结束时关闭 rows，释放数据库连接

	// 获取列名
	columns, _ := rows.Columns()

	// 创建切片和扫描参数，用于逐行处理数据库查询结果
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	// 将每一列的值都映射到 scanArgs 中
	for j := range values {
		scanArgs[j] = &values[j]
	}

	// 存储查询结果
	res := make([]map[string]interface{}, 0, 5)

	// 遍历查询结果的每一行
	for rows.Next() {
		record := make(map[string]interface{}) // 存储每一行数据的 map
		// 扫描当前行的所有列的值
		rows.Scan(scanArgs...)
		// 将每列的值与对应列名一起存储到 record 中
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []byte:
					// 如果列的值是字节数组，转换为字符串
					record[columns[i]] = string(col.([]byte))
				default:
					// 其他类型直接存储
					record[columns[i]] = col
				}
			}
		}
		// 将记录添加到结果切片中
		res = append(res, record)
	}

	// 返回所有记录
	return res
}

// QueryFirst 方法用于获取查询结果中的第一条记录。
// rows: 代表查询结果的 sql.Rows 对象。
// err: 查询过程中发生的错误，如果为 nil，则表示查询成功。
// 返回：查询结果中的第一条记录，如果没有记录，则返回 nil。
func (d *dBUtil) QueryFirst(rows *sql.Rows, err interface{}) map[string]interface{} {
	// 调用 Query 方法获取所有结果
	res := d.Query(rows, err)
	// 如果有结果，则返回第一条记录
	if len(res) > 0 {
		return res[0]
	}
	// 如果没有结果，返回 nil
	return nil
}
