package main

import ("fmt"
	"net/http"
	"log"
	"os"
	"database/sql"
	"github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))		//  assigned by Heroku as the PORT environment variable
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func testDB(){
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err);
	}

	db.Exec(`CREATE TABLE TEST_TABLE(
		ID INT NOT NULL,
		NAME TEST_FIELD
		PRIMARY KEY (ID)
	);`)

	db.Exec(`INSERT INTO TEST_TABLE
		VALUES(2, 4, 8)
	);`)

	db.Exec(`SELECT * FROM TEST_TABLE`)
}