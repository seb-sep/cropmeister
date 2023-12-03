package main

import (
	"fmt"

	"time"

	"github.com/seb-sep/cropmeister/db"

	"math/rand"

	"database/sql"

	"github.com/go-faker/faker/v4"
)

func generateRandomInt(min, max int) sql.NullInt32 {
	value := rand.Intn(max-min+1) + min
	return sql.NullInt32{Int32: int32(value), Valid: true}
}

func generateRandomFloat(min, max int) sql.NullFloat64 {
	value := float64(rand.Intn(max-min+1)+min) + rand.Float64()
	return sql.NullFloat64{Float64: value, Valid: true}
}

func generateRandomBool() sql.NullBool {
	value := rand.Intn(2) == 1
	return sql.NullBool{Bool: value, Valid: true}
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

func generateFarm(farmid sql.NullInt32) db.Farm {
	farm := db.Farm{
		Name:          faker.Name(),
		FarmValue:     generateRandomInt(1, 100),
		FarmID:        farmid,
		AddressStreet: toSqlNullString(faker.GetRealAddress().Address),
		AddressCity:   toSqlNullString(faker.GetRealAddress().City),
		AddressState:  toSqlNullString(faker.GetRealAddress().State),
		AddressZip:    toSqlNullString(faker.GetRealAddress().PostalCode),
	}
	return farm
}

func generateCrop(cropType sql.NullString) db.Crop {
	crop := db.Crop{
		CropType:           cropType,
		PhRangeWeight:      generateRandomFloat(0, 1),
		PhRangeDesired:     generateRandomFloat(1, 100),
		WaterNeededWeight:  generateRandomFloat(0, 1),
		WaterNeededDesired: generateRandomFloat(1, 100),
		SunRangeWeight:     generateRandomFloat(0, 1),
		SunRangeDesired:    generateRandomFloat(1, 100),
		Banned:             generateRandomBool(),
	}
	return crop
}

func generatePurchase(purchaseID sql.NullInt32, cropType sql.NullString) db.Purchase {
	purchase := db.Purchase{
		PurchaseID:       purchaseID,
		CropType:         cropType,
		PurchaseComplete: generateRandomBool(),
		TotalPrice:       generateRandomFloat(1, 50000),
		TotalQuantity:    generateRandomInt(1, 100),
		PurchaseDate:     toSqlNullTime(faker.Date()),
	}
	return purchase
}

func generateCropInvestor(name string) db.CropInvestor {
	cropInvestor := db.CropInvestor{
		Name:            name,
		BuyPrice:        generateRandomInt(1, 50000),
		InvestibleMoney: generateRandomInt(1, 1000000),
		SellPrice:       generateRandomInt(1, 60000),
	}
	return cropInvestor
}

func generateCropInvestigator(name string) db.CropInvestigator {
	cropInvestigator := db.CropInvestigator{
		Name:   name,
		Usdaid: generateRandomInt(1, 100),
	}
	return cropInvestigator
}

func generateDistrictCode(codeID sql.NullInt32, cropType sql.NullString) db.DistrictCode {
	districtCode := db.DistrictCode{
		MaxWater: generateRandomInt(1, 100),
		MaxFert:  generateRandomInt(1, 100),
		CropType: cropType,
		CodeID:   codeID,
	}
	return districtCode
}

func generateCropBuyer(name sql.NullString, cropType sql.NullString) db.CropBuyer {
	cropBuyer := db.CropBuyer{
		Name:               name,
		QuantitiesRequired: generateRandomInt(1, 10),
		CropType:           cropType,
		TargetPrice:        generateRandomInt(1, 100),
	}
	return cropBuyer
}
