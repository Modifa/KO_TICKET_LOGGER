package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/Modifa/KO_TICKET_LOGGER/models"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
}

// Register user
// Register user
func (db *DB) RegisterDeveloper(functionnamewithschema string, m interface{}) ([]models.ResponseDB, error) {
	User := []models.ResponseDB{}
	u := ConVertInterface(functionnamewithschema, m)
	ctx := context.Background()
	db1, _ := pgxpool.Connect(ctx, os.Getenv("PostgresConnectionString"))
	defer db1.Close()
	//
	err := pgxscan.Select(ctx, db1, &User, u)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr)         // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}
	return User, err
}

// Convert Interface and return Query string
func ConVertInterface(funcstr string, m interface{}) string {
	q := "select * from " + funcstr + "("

	if m != nil {

		v := reflect.ValueOf(m)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {

			switch typeOfS.Field(i).Type.Name() {
			case "int", "int16", "int32", "int64", "int8":
				str := v.Field(i).Interface().(int64)
				strInt64 := strconv.FormatInt(str, 10)
				q += strInt64 + ","
			case "float64":
				str := v.Field(i).Interface().(float64)
				s := fmt.Sprintf("%f", str)
				q += s + ","
			case "bool":
				q += "'" + strconv.FormatBool(v.Field(i).Interface().(bool)) + "',"
			default:
				if v.Field(i).Interface().(string) == "" {
					q += "null,"
				} else {
					q += "'" + v.Field(i).Interface().(string) + "',"
				}
			}
		}

		q = q[0 : len(q)-len(",")]
	}

	q += ")"

	return q
}
