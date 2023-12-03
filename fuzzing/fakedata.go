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

func generateFarm() db.Farm {
	farm := db.Farm{
		Name:          faker.Name(),
		FarmValue:     generateRandomInt(1, 100),
		FarmID:        generateRandomInt(1, 100),
		AddressStreet: toSqlNullString(faker.GetRealAddress().Address),
		AddressCity:   toSqlNullString(faker.GetRealAddress().City),
		AddressState:  toSqlNullString(faker.GetRealAddress().State),
		AddressZip:    toSqlNullString(faker.GetRealAddress().PostalCode),
	}
	return farm
}

func generateCrop() db.Crop {
	crop := db.Crop{
		CropType:           toSqlNullString(faker.Word()),
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

func generatePurchase() db.Purchase {
	purchase := db.Purchase{
		PurchaseID:       generateRandomInt(1, 100),
		CropType:         toSqlNullString(faker.Word()),
		PurchaseComplete: generateRandomBool(),
		TotalPrice:       generateRandomFloat(1, 50000),
		TotalQuantity:    generateRandomInt(1, 100),
		PurchaseDate:     toSqlNullTime(faker.Date()),
	}
	return purchase
}

func generateCropInvestor() db.CropInvestor {
	cropInvestor := db.CropInvestor{
		Name:            faker.Name(),
		BuyPrice:        generateRandomInt(1, 50000),
		InvestibleMoney: generateRandomInt(1, 1000000),
		SellPrice:       generateRandomInt(1, 60000),
	}
	return cropInvestor
}

func generateCropInvestigator() db.CropInvestigator {
	cropInvestigator := db.CropInvestigator{
		Name:   faker.Name(),
		Usdaid: generateRandomInt(1, 100),
	}
	return cropInvestigator
}

func generateDistrictCode() db.DistrictCode {
	districtCode := db.DistrictCode{
		MaxWater: generateRandomInt(1, 100),
		MaxFert:  generateRandomInt(1, 100),
		CropType: toSqlNullString(faker.Word()),
		CodeID:   generateRandomInt(1, 100),
	}
	return districtCode
}

func generateCropBuyer() db.CropBuyer {
	cropBuyer := db.CropBuyer{
		Name:               toSqlNullString(faker.Name()),
		QuantitiesRequired: generateRandomInt(1, 10),
		CropType:           toSqlNullString(faker.Name()),
		TargetPrice:        generateRandomInt(1, 100),
	}
	return cropBuyer
}

func main() {
	farm := generateFarm()
	fmt.Println(farm)

	crop := generateCrop()
	fmt.Println(crop)

	purchase := generatePurchase()
	fmt.Println(purchase)

	cropInvestor := generateCropInvestor()
	fmt.Println(cropInvestor)

	cropInvestigator := generateCropInvestigator()
	fmt.Println(cropInvestigator)

	districtCode := generateDistrictCode()
	fmt.Println(districtCode)

	cropBuyer := generateCropBuyer()
	fmt.Println(cropBuyer)
}
