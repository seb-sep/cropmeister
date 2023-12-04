-- Get Harvest
SELECT * FROM Harvest;

-- Post Harvest
INSERT INTO Harvest (Quantity, Time_Year, Time_Season, Ph_Base, Ph_Fertilized, Water_Rain, Water_Sprinkler, Sun, Price,
                     CropType, FarmID, Extinct)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- Get Crop {croptype}
SELECT *,
       (BasePrice +
        (Crop.WaterNeeded_Weight - Crop.WaterNeeded_Desired) +
        (Crop.SunRange_Weight - Crop.SunRange_Desired) -
        CASE
            WHEN Restricted THEN BasePrice * 0.5
            ELSE 0
        END) AS CalculatedPrice
FROM Crop
WHERE CropType = ?;

-- Put Crop {croptype}
UPDATE Crop
SET
    BasePrice = ?,
    PhRange_Weight = ?,
    PhRange_Desired = ?,
    WaterNeeded_Weight = ?,
    WaterNeeded_Desired = ?,
    SunRange_Weight = ?,
    SunRange_Desired = ?,
    Banned = ?
WHERE CropType = ?;

-- Delete Crop {croptype}
UPDATE Crop
SET Banned = TRUE
WHERE CropType = ?;

-- Get Farm
SELECT * FROM Farm;

-- Post Farm
INSERT INTO Farm (Name, FarmValue, FarmID, Address_Street, Address_City, Address_State, Address_Zip)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- Get Farm/{FarmID}
SELECT * FROM Farm
WHERE FarmID = ?;

-- Put Farm/{FarmID}
UPDATE Farm
SET
    Name = ?,
    FarmValue = ?,
    Address_Street = ?,
    Address_City = ?,
    Address_State = ?,
    Address_Zip = ?
WHERE FarmID = ?;

-- Delete Farm/{FarmID}
UPDATE Farm
SET Active = FALSE
WHERE FarmID = ?;

-- Get Harvest/{CropType}
SELECT * FROM Harvest
WHERE CropType = ?;

-- Put Harvest/{CropType}
UPDATE Harvest
SET
    Quantity = ?,
    Time_Year = ?,
    Time_Season = ?,
    Ph_Base = ?,
    Ph_Fertilized = ?,
    Water_Rain = ?,
    Water_Sprinkler = ?,
    Sun = ?,
    Price = ?,
    FarmID = ?,
    Extinct = ?
WHERE CropType = ?;

-- Delete Harvest/{CropType}
UPDATE Harvest
SET Extinct = TRUE
WHERE CropType = ?;

-- Get Purchase
SELECT * FROM Purchase;

-- Post Purchase
INSERT INTO Purchase (CropType, PurchaseComplete, TotalPrice, TotalQuantity, PurchaseDate)
VALUES (?, ?, ?, ?, ?);

-- Get Purchase/{PurchaseID}
SELECT * FROM Purchase
WHERE PurchaseID = ?;

-- Put Purchase/{PurchaseID}
UPDATE Purchase
SET
    CropType = ?,
    PurchaseComplete = ?,
    TotalPrice = ?,
    TotalQuantity = ?,
    PurchaseDate = ?
WHERE PurchaseID = ?;

-- Delete Purchase/{PurchaseID}
UPDATE Purchase
SET Canceled = TRUE
WHERE PurchaseID = ?;

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