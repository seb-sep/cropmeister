package fakedata

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

func CalculatePrice(num1, num2, num3 float64) float64 {
	return (num1 / 30) * (num2 / 60) * num3
}

func generateRandomInt(min, max int) sql.NullInt32 {
	value := rand.Intn(max-min+1) + min
	return sql.NullInt32{Int32: int32(value), Valid: true}
}

func generateRandomPlainFloat(min, max int) float64 {
	value := float64(rand.Intn(max-min+1)+min) + rand.Float64()
	return value
}

func generateRandomFloat(min, max int) sql.NullFloat64 {
	value := float64(rand.Intn(max-min+1)+min) + rand.Float64()
	return sql.NullFloat64{Float64: value, Valid: true}
}

func generateRandomBool() sql.NullBool {
	value := rand.Intn(2) == 1
	return sql.NullBool{Bool: value, Valid: true}
}

func toSqlNullFloat(value float64) sql.NullFloat64 {
	return sql.NullFloat64{Float64: value, Valid: true}
}

func toSqlNullInt(value float64) sql.NullInt32 {
	return sql.NullInt32{Int32: int32(value), Valid: true}
}

func toSqlNullString(value string) sql.NullString {
	return sql.NullString{String: value, Valid: true}
}

func toSqlNullTime(value string) sql.NullTime {
	layout := "2006-01-02" // Specify the layout of the input string
	date, err := time.Parse(layout, value)
	if err != nil {
		fmt.Println(err)
	}
	return sql.NullTime{Time: date, Valid: true}
}
