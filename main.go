package main

import ("fmt"
	"net/http"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {
	testDB()
	http.HandleFunc("/", handler)
	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))		//  assigned by Heroku as the PORT environment variable
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func testDB(){
	var rows *sql.Rows
	var id int
	var name string

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err);
	}

	_, err = db.Exec(`CREATE TABLE TEST_TABLE(
		ID INT NOT NULL,
		SOME_VALUE INT,
		PRIMARY KEY (ID)
	);`)

	if err != nil {
		log.Println(err);
	}


	_, err = db.Exec(`INSERT INTO TEST_TABLE VALUES (1), (2), (3);`)
	if err != nil {
		log.Println(err);
	}

	rows, err = db.Query(`SELECT * FROM TEST_TABLE`)
	if err != nil {
		log.Println(err);
	} else {
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(id, name)
		}
	}
}