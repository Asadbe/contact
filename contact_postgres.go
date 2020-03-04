package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "asadbek"
	password = "1"
	dbname   = "mydb"
)

type sqlxDB struct {
	connectDB *sqlx.DB
}

var err error

//NewAnimalList ...
func NewContactList() (ContactListint, error) {
	cm := sqlxDB{}
	psqlInfo := fmt.Sprintf(` user=%s dbname=%s password=%s`, user, dbname, password)
	cm.connectDB, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &cm, nil
}
func (s *sqlxDB) Add(a ContactList) error {
	insertionQuery := `insert into conntacts (name, email, number,age) values ($1, $2, $3,$4);`

	_, err := s.connectDB.Exec(insertionQuery, a.name, a.email, a.number,a.age)

	if err != nil {
		return err
	}

	return nil
}
func (s *sqlxDB) Update(name string, email string, number int, age int,id int) error {
	updatingQuery := `update conacts set name=$1,email=$2, number=$3,age=$4
	where id =$5`

	_, err := s.connectDB.Exec(updatingQuery, name, email, number, age ,id )

	if err != nil {
		fmt.Println("Can't update")
		return err
	}

	return nil
}
func (s *sqlxDB) Delete(id int) error {
	deletingQuery := `delete from contacts where id=$1;`

	_, err := s.connectDB.Exec(deletingQuery, id)

	if err != nil {
		fmt.Println("Can't delete")
		return err
	}
	return nil
}
func (s *sqlxDB) GetContact(id int) error {
	listContactQuery := `select id,name,email,number from contacts where id=$1;`

	rows, err := s.connectDB.Queryx(listContactQuery, id)

	if err != nil {
		fmt.Println("Can't print task list")
		return err
	}
	var ts ContactList
	for rows.Next() {
		err = rows.Scan(&ts.id, &ts.name, &ts.email, &ts.number,&ts.age)
		if err != nil {
			fmt.Println("Can't scan struct")
			return err
		}
	}
	fmt.Println(ts)
	return nil

}
