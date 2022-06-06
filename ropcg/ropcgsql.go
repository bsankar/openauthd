/*
 Copyright 2022 Balaje Sankar

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package ropcg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/microsoft/go-mssqldb"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {
	dbusername := os.Getenv("DBUSERNAME")
	dbpassword := os.Getenv("DBPASSWORD")
	var connstring = `sqlserver://` + dbusername + `:` + dbpassword + `@localhost/SQLExpress?database=auth&connection+timeout=30`
	var err error
	db, err = sql.Open("sqlserver", connstring)
	if err != nil {
		log.Fatal("Error connection ", err.Error())
		db = nil
	}
}

func ReadPrincipalRecord(username string) (outputdata []string, err error) {
	outputdata = make([]string, 3)
	ctx := context.TODO()
	err = db.PingContext(ctx)
	if err != nil {
		log.Println("Connection not available ", err.Error())
		return nil, err
	}

	query := `select principalId,principalName,principalSecret from dbo.principals where principalName=@principalName`
	row := db.QueryRowContext(ctx, query, sql.Named("principalName", username))
	err = row.Scan(&outputdata[0], &outputdata[1], &outputdata[2])
	if err != nil {
		log.Println("Error Returning data from DB ", err.Error())
		return nil, err
	}
	return outputdata, nil
}

func CreatePrincipalRecord(userName string, password string) (created bool, err error) {
	created = false
	ctx := context.TODO()
	err = db.PingContext(ctx)
	if err != nil {
		log.Println("Connection not available", err.Error())
		return
	}

	query := `insert into dbo.principals(principalID,principalName,principalSecret) values(@principalId,@principalName,@principalSecret)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error creating db statement", err.Error())

		return
	}

	defer stmt.Close()
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	_, err = stmt.ExecContext(ctx, sql.Named("principalId", uuid.New()), sql.Named("principalName", userName), sql.Named("principalSecret", string(bytes)))
	if err != nil {
		log.Println("Error createing user ", err.Error())
		return
	}

	return true, nil

}

func main() {
	data, err := ReadPrincipalRecord("user1")
	if err != nil {
		log.Fatal("Error")
	}

	for _, out := range data {
		fmt.Println(out)
	}
}
