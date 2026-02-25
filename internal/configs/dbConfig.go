package configs
import(
	"database/sql"
	_"github.com/lib/pq"
	"log"
)
var(
	DB= InitDB(NewConfig().DatabaseURL)
)
func InitDB(databaseURL string)*sql.DB{
	db,err:=sql.Open("postgres",databaseURL)
	if err!=nil{
		log.Fatal(err)
	}
	if err :=db.Ping();err!=nil{
		log.Fatal(err)
	}
	return db
}

