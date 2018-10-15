package get

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var count int

type Post struct {
	No   string
	Name string
	H    string
	A    string
	B    string
	C    string
	D    string
	S    string
	Sum  string
}

func Select(text string) (Post, error) {
	var Pokemondata Post
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	//stmtOut, errs := db.Prepare(fmt.Sprintf("SELECT No, Name, H, A, B, C, D, S, Sum FROM value WHERE Name = %s",text))
	if err = db.QueryRow("SELECT * FROM pokemon WHERE Name = ?", text).Scan(&Pokemondata.No, &Pokemondata.Name, &Pokemondata.H, &Pokemondata.A, &Pokemondata.B, &Pokemondata.C, &Pokemondata.D, &Pokemondata.S, &Pokemondata.Sum); err != nil {
		return Pokemondata, errors.New("ポケモンが見つかりませんでした")
	}
	return Pokemondata, nil
}
