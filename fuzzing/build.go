package fakedata

import (
	"fmt"

	"github.com/go-faker/faker/v4"
	"github.com/seb-sep/cropmeister/db"
)

func GenerateList(generator func() string, count int) []string {
	list := make([]string, count)
	for i := 0; i < count; i++ {
		list[i] = generator()
	}
	return list
}

func GenerateFakeInstances() {
	// Generate 15 or more fake instances using each function
	for i := 0; i < 15; i++ {
		farmId := generateRandomInt(1, 100)
		farmerId := generateRandomInt(1, 100)

		investorName := faker.Name()
		investigatorName := faker.Name()
		buyerName := faker.Name()
		cropInvestor := generateCropInvestor(investorName)
		cropInvestigator := generateCropInvestigator(investigatorName)
		farmer := generateFarmer(farmerId)
		farm := generateFarm(farmId)

		fakeWord := func() string {
			return faker.Word()
		}
		cropNames := GenerateList(fakeWord, 5)

		cropBuyers := make([]db.CropBuyer, len(cropNames))
		crops := make([]db.Crop, len(cropNames))
		harvests := make([]db.Harvest, len(cropNames)*5)
		purchases := make([]db.Purchase, len(cropNames)*5)
		districts := make([]db.DistrictCode, len(cropNames)*5)
		for i := 0; i < len(cropNames); i++ {
			cropBuyers[i] = generateCropBuyer(toSqlNullString(buyerName), toSqlNullString(cropNames[i]))
			crops[i] = generateCrop(toSqlNullString(cropNames[i]))
			districts[i] = generateDistrictCode(generateRandomInt(1, 100), toSqlNullString(cropNames[i]))

			for j := 0; j < 5; j++ {
				hrv := generateHarvest(farmId, toSqlNullString(cropNames[i]))
				harvests[j+(i*5)] = hrv
				purchases[j+(i*5)] = generatePurchase(generateRandomInt(1, 100), hrv.CropType, farm.FarmID)
			}
		}

		//TODO: push schema to table
		fmt.Println(farmer, cropInvestor, cropInvestigator)
	}
}
