package main
import (
	"database/sql" //这包一定要引用，是底层的sql驱动
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv" //这个是为了把int转换为string
)
func main() {  //main函数
	db, err := sql.Open("mysql", "ormgo:ormgo@tcp(localhost:3306)/ormgo?charset=utf8")
	//数据库连接字符串，别告诉我看不懂。端口一定要写/
	if err != nil {  //连接成功 err一定是nil否则就是报错
		panic(err.Error()) //抛出异常
		fmt.Println(err.Error())//仅仅是显示异常
	}
	defer db.Close()  //只有在前面用了 panic 这时defer才能起作用，如果链接数据的时候出问题，他会往err写数据

	rows, err := db.Query("select id,name from user where id = ?",3)
	//判断err是否有错误的数据，有err数据就显示panic的数据
	if err != nil {
		panic(err.Error())
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()
	var id int  //定义一个id 变量
	var name string //定义lvs 变量
	for rows.Next() { //开始循环
		rerr := rows.Scan(&id, &name)  //数据指针，会把得到的数据，往刚才id 和 lvs引入
		if rerr == nil {
			fmt.Println("id号是",strconv.Itoa(id) + "     lvs name是"+ name) //输出来而已，看看
		}
	}
	insert_sql := "INSERT INTO user(name) VALUES(?)"
	_, e4 := db.Exec(insert_sql,"nima")
	fmt.Println(e4)
	db.Close() //关闭数据库
}
