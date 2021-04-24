package main
import(
	"fmt"
	"errors"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

//答案：应该 Wrap 这个错误，因为不同的DAO查询都可能返回相同的ErrNoRows，所以应该Wrap这个错误，包含查询的语句和参数。
func main() {
	name, err:= getUserNameDao(1)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println(err);
	}
	fmt.Println("User Name ", name);
}

func getUserNameDao(id int) (name string, err error) {
	db, err := sql.Open("mysql", "root:1234567@tcp(172.17.0.3:3306)/homework")
	
	defer db.Close()

	err = db.Ping()
    if err != nil {
		return "", fmt.Errorf("DB Connect Error: %w", err);
    }

	var row_name string;
	var queryString = "select row_name from homework_1 where id = ?";

	err = db.QueryRow(queryString, id).Scan(&row_name);

	switch {
    	case err == sql.ErrNoRows:
			return "", fmt.Errorf("No Data With Query:%v, Param:%v: %w", queryString, id, err);
		case err != nil:
			return "", fmt.Errorf("Unknown Err With Query:%v, Param:%v: %w", queryString, id, err);	
	}

	return row_name, err
}