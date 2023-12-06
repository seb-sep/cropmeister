package fakedata

import (
	"github.com/seb-sep/cropmeister/db"

	"database/sql"

	"github.com/go-faker/faker/v4"
)

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

func generateFarmer(farmid sql.NullInt32) db.Farmer {
	farmer := db.Farmer{
		Name:     faker.Name(),
		FarmID:   farmid,
		Budget:   generateRandomInt(1, 100000),
		NetWorth: generateRandomInt(1, 1000000),
	}
	return farmer
}

func generateHarvest(farmID sql.NullInt32, cropType sql.NullString) db.Harvest {

	phb := generateRandomPlainFloat(0, 10)
	phf := generateRandomPlainFloat(0, 5)
	wr := generateRandomPlainFloat(0, 10)
	ws := generateRandomPlainFloat(0, 5)
	s := generateRandomPlainFloat(0, 10)
	p := CalculatePrice(phb+phf, wr+ws, s)

	harvest := db.Harvest{
		FarmID:         farmID,
		CropType:       cropType,
		Quantity:       generateRandomInt(1, 100),
		Extinct:        generateRandomBool(),
		HarvestDate:    toSqlNullTime(faker.Date()),
		PhBase:         toSqlNullFloat(s),
		Sun:            toSqlNullInt(s),
		PhFertilized:   toSqlNullFloat(phf),
		WaterRain:      toSqlNullFloat(wr),
		WaterSprinkler: toSqlNullFloat(ws),
		Price:          toSqlNullFloat(p),
	}
	return harvest
}

func generateCrop(cropType string) db.Crop {
	crop := db.Crop{
		CropType:           cropType,
		PhRangeWeight:      generateRandomFloat(0, 1),
		PhRangeDesired:     generateRandomFloat(1, 100),
		WaterNeededWeight:  generateRandomFloat(0, 1),
		WaterNeededDesired: generateRandomFloat(1, 100),
		SunRangeWeight:     generateRandomFloat(0, 1),
		SunRangeDesired:    generateRandomFloat(1, 100),
		BasePrice:          generateRandomFloat(1, 10),
		Banned:             generateRandomBool(),
	}
	return crop
}

func generatePurchase(purchaseID sql.NullInt32, CropType sql.NullString, FarmID sql.NullInt32) db.Purchase {
	purchase := db.Purchase{
		PurchaseID:       purchaseID,
		CropType:         CropType,
		FarmID:           FarmID,
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
