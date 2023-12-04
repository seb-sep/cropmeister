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

-- /Farmer/{FarmID}
-- Get
SELECT *
FROM Farmer
WHERE FarmID = ?;
-- Post
INSERT INTO Farmer (Name, Budget, NetWorth, FarmID, PurchaseID)
SELECT *
FROM Farmer
WHERE FarmID = ?;

-- /CropInspector
-- Get
SELECT *
FROM Crop_Investigator;
-- Post
INSERT INTO Crop_Investigator (Name, USDAID)
VALUES (?, ?);

-- /DistinctCode/Crop/{CropType}
-- Get
SELECT *
FROM District_Code
WHERE CropType = ?;
-- Post
INSERT INTO District_Code (MaxWater, MaxFert, CropType, CodeID)
SELECT *
FROM District_Code
WHERE CropType = ?;

-- /DistrictCode/Inspector/{USDAID}
-- Get
SELECT *
FROM District_Code
WHERE farm.Crop_Investigator.USDAID = ?;
-- Post
INSERT INTO District_Code (MaxWater, MaxFert, CropType, CodeID)
SELECT *
FROM District_Code
WHERE farm.Crop_Investigator.USDAID = ?;

-- /CropInspector/Code/{CodeID}
-- Get
SELECT *
FROM Crop_Investigator
WHERE farm.District_Code.CodeID = ?;
-- Post
INSERT INTO Crop_Investigator (Name, USDAID)
SELECT *
FROM Crop_Investigator
WHERE farm.District_Code.CodeID = ?;

-- /CodeInspector/{USDAID}
-- Get
SELECT *
FROM Crop_Investigator
WHERE USDAID = ?;
-- Put
UPDATE Crop_Investigator
SET Name = ?
WHERE USDAID = ?;
-- Delete
DELETE FROM Crop_Investigator
WHERE USDAID = ?;

-- /DistrictCode
-- Get
SELECT *
FROM District_Code;
-- Post
INSERT INTO District_Code (MaxWater, MaxFert, CropType, CodeID)
VALUES (?, ?, ?, ?);

-- DistrictCode/Code/{CodeID}
-- Get
SELECT *
FROM District_Code
WHERE CodeID = ?;
-- Put
UPDATE District_Code
SET MaxWater = ?, MaxFert = ?,
    CropType = ?, CodeID = ?
WHERE CodeID = ?;
-- Delete
DELETE FROM District_Code
WHERE CodeID = ?;