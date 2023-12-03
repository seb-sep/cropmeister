-- name: AddCropBuyer :exec
INSERT INTO Crop_Buyer (Name, Quantities_Required, Crop_Type, Target_Price)
VALUES (?, ?, ?, ?);

-- name: AddFarmer :exec
INSERT INTO Farmer (Name, Budget, Net_Worth, Farm_ID, Purchase_ID)
VALUES (?, ?, ?, ?, ?);

-- name: AddFarm :exec
INSERT INTO Farm (Name, Farm_Value, Address_Street, Address_City, Address_State, Address_Zip)
VALUES (?, ?, ?, ?, ?, ?);

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
VALUES (?,?,?,?,?,?,?,?,?,?,?,?);

-- name: AddMonitorsBuyer :exec
INSERT INTO Monitors_Buyer (Name, Crop_Type)
VALUES (?,?);

-- name: AddInvestsIn :exec
INSERT INTO Invests_In (Name, Farm_ID)
VALUES (?,?);

-- name: AddEnforces :exec
INSERT INTO Enforces (USDAID, Code_ID)
VALUES (?,?);

-- name: AddInvestMonitor :exec
INSERT INTO Monitors_Investments (Name, Crop_Type)
VALUES (?,?);
