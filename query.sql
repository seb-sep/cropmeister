-- name: AddCropBuyer :exec
INSERT INTO Crop_Buyer (Name, Quantities_Required, Crop_Type, Target_Price)
VALUES ($1, $2, $3, $4);

-- name: AddFarmer :exec
INSERT INTO Farmer (Name, Budget, Net_Worth, Farm_ID, Purchase_ID)
VALUES ($1, $2, $3, $4, $5);

-- name: AddFarm :exec
INSERT INTO Farm (Name, Farm_Value, Address_Street, Address_City, Address_State, Address_Zip)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: AddHarvest :exec
INSERT INTO Harvest (
  Quantity,
  Time_Year,
  Time_Season,
  Ph_Base,
  Ph_Fertilized,
  Water_Rain,
  Water_Sprinkler,
  Sun,
  Price,
  Crop_Type,
  Farm_ID,
  Extinct
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);

-- name: AddMonitorsBuyer :exec
INSERT INTO Monitors_Buyer (Name, Crop_Type)
VALUES ($1, $2);

-- name: AddInvestsIn :exec
INSERT INTO Invests_In (Name, Farm_ID)
VALUES ($1, $2);

-- name: AddEnforces :exec
INSERT INTO Enforces (USDAID, Code_ID)
VALUES ($1, $2);

-- name: AddInvestMonitor :exec
INSERT INTO Monitors_Investments (Name, Crop_Type)
VALUES ($1, $2);
