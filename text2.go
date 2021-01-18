package main

import (
	"database/sql"
	"fmt"
	"strings"
)

import (
	_ "github.com/mattn/go-adodb"
)

type Mssql struct {
	*sql.DB
	dataSource string
	database   string
	windows    bool
	sa         SA
}

type SA struct {
	user   string
	passwd string
}

func (m *Mssql) Open() (err error) {
	var conf []string
	conf = append(conf, "Provider=SQLOLEDB")
	conf = append(conf, "Data Source="+m.dataSource)
	if m.windows {
		// Integrated Security=SSPI 这个表示以当前WINDOWS系统用户身去登录SQL SERVER服务器(需要在安装sqlserver时候设置)，

		//相关推荐：SQL Server超连接查询

		//超连接查询也是连接查询，所以必须有两张或两张以上的表，超连接查询共包括四种，分别是内连接查询、左连接查询、右连接查询、全连接查询。 1、内连接查询 内连接查询也可以在on后面带有条件，如： select 姓名,城市 from 仓库 inner join 职工 on 职工.仓库


			// 如果SQL SERVER服务器不支持这种方式登录时，就会出错。
			conf = append(conf, "integrated security=SSPI")
		}
		conf = append(conf, "Initial Catalog="+m.database)
		conf = append(conf, "user id="+m.sa.user)
		conf = append(conf, "password="+m.sa.passwd)

		m.DB, err = sql.Open("adodb", strings.Join(conf, ";"))
		if err != nil {
			return err
		}
		return nil
	}

	func main() {
		db := Mssql{
			dataSource: "127.0.0.1",
			database:   "testOne",
			// windwos: true 为windows身份验证，false 必须设置sa账号和密码
			windows: true,
			sa: SA{
				user:   "sa",
				passwd: "chen360219",
			},
		}
		// 连接数据库
		err := db.Open()
		if err != nil {
			fmt.Println("sql open:", err)
			return
		}
		defer db.Close()

		// 执行SQL语句
		rows, err := db.Query("select * from info")
		if err != nil {
			fmt.Println("query: ", err)
			return
		}
		for rows.Next() {
			var name string
			var number int
			rows.Scan(&name, &number)
			fmt.Printf("Name: %s \t Number: %d\n", name, number)
		}
	}