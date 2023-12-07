-- Get Harvest
SELECT * FROM Harvest;

-- name: GetCropData :many
SELECT *,
       (Base_Price +
        (Crop.Water_Needed_Weight - Crop.WaterNeeded_Desired) +
        (Crop.Sun_Range_Weight - Crop.SunRange_Desired) -
        CASE
            WHEN Restricted THEN Base_Price * 0.5
            ELSE 0
        END) AS CalculatedPrice
FROM Crop
WHERE Crop_Type = ?;

-- name: UpdateCrop :execresult
UPDATE Crop
SET
    Base_Price = ?,
    Ph_Range_Weight = ?,
    Ph_Range_Desired = ?,
    Water_Needed_Weight = ?,
    Water_Needed_Desired = ?,
    Sun_Range_Weight = ?,
    Sun_Range_Desired = ?,
    Banned = ?
WHERE Crop_Type = ?;

-- name: DeleteCrop :execresult
DELETE FROM Crop
WHERE Crop_Type = ?;

-- name: BanCrop :execresult
UPDATE Crop
SET Banned = TRUE
WHERE Crop_Type = ?;

-- name: GetFarms :many
SELECT * FROM Farm;

-- name: AddFarm :execresult
INSERT INTO Farm (Name, Farm_Value, Address_Street, Address_City, Address_State, Address_Zip)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GeTFarm :one
SELECT * FROM Farm
WHERE Farm_ID = ?;

-- name: UpdateFarm :execresult
UPDATE Farm
SET
    Name = ?,
    Farm_Value = ?,
    Address_Street = ?,
    Address_City = ?,
    Address_State = ?,
    Address_Zip = ?
WHERE Farm_ID = ?;

-- name: DeleteFarm :execresult
UPDATE Farm
SET Active = FALSE
WHERE Farm_ID = ?;

-- name: GetHarvests :many
SELECT * FROM Harvest
WHERE Crop_Type = ?;

-- name: UpdateHarvest :execresult
UPDATE Harvest
SET
    Quantity = ?,
    Harvest_Date = ?,
    Ph_Base = ?,
    Ph_Fertilized = ?,
    Water_Rain = ?,
    Water_Sprinkler = ?,
    Sun = ?,
    Price = ?,
    Farm_ID = ?,
    Extinct = ?
WHERE Crop_Type = ?;

-- name: DeleteHarvest :execresult
UPDATE Harvest
SET Extinct = TRUE
WHERE Crop_Type = ?;

-- name: GetPurchases :many
SELECT * FROM Purchase;

-- name: AddPurchase :execresult
INSERT INTO Purchase (Crop_Type, Farm_ID, Purchase_Complete, Total_Price, Total_Quantity, Purchase_Date)
VALUES (?, ?, ?, ?, ?, ?);

-- name: GetPurchase :one
SELECT * FROM Purchase
WHERE Purchase_ID = ?;

-- name: UpdatePurchase :execresult
UPDATE Purchase
SET
    Crop_Type = ?, 
    Farm_ID = ?,
    Purchase_Complete = ?,
    Total_Price = ?,
    Total_Quantity = ?,
    Purchase_Date = ?
WHERE Purchase_ID = ?;

-- name: DeletePurchase :execresult
UPDATE Purchase
SET Canceled = TRUE
WHERE Purchase_ID = ?;

-- name: AddCropBuyer :execresult
INSERT INTO Crop_Buyer (Name, Quantities_Required, Crop_Type, Target_Price)
VALUES (?, ?, ?, ?);

-- name: AddHarvest :execresult
INSERT INTO Harvest (
  Quantity,
  Harvest_Date,
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
VALUES (?,?,?,?,?,?,?,?,?,?,?);

-- name: AddMonitorsBuyer :execresult
INSERT INTO Monitors_Buyer (Name, Crop_Type)
VALUES (?,?);

-- name: AddInvestsIn :execresult
INSERT INTO Invests_In (Name, Farm_ID)
VALUES (?,?);

-- name: AddEnforces :execresult
INSERT INTO Enforces (USDAID, Code_ID)
VALUES (?,?);

-- name: AddInvestMonitor :execresult
INSERT INTO Monitors_Investments (Name, Crop_Type)
VALUES (?,?);

-- name: GetFarmer :one
SELECT *
FROM Farmer
WHERE Farm_ID = ?;

-- name: AddFarmer :execresult
INSERT INTO Farmer (Name, Budget, Net_Worth, Farm_ID)
VALUES (?, ?, ?, ?);

-- name: GetCropInspectors :many
SELECT *
FROM Crop_Inspector;

-- name: AddCropInvestigator :execresult
INSERT INTO Crop_Inspector (Name, USDAID)
VALUES (?, ?);

-- name: GetDistrictsWithCrop :many
SELECT *
FROM District_Code
WHERE Crop_Type = ?;

-- name: AddDistrictCode :execresult
INSERT INTO District_Code (Max_Water, Max_Fert, Crop_Type, Code_ID)   
VALUES (?, ?, ?, ?);

-- name: GetDistrictsForInspector :many
SELECT * 
FROM Crop_Inspector
JOIN District_Code
ON Crop_Inspector.USDAID = District_Code.Code_ID
WHERE USDAID = ?;


-- name: GetInspectorForDistrict :one
SELECT *
FROM District_Code
JOIN Crop_Inspector 
ON District_Code.Code_ID = Crop_Inspector.Code_ID
WHERE Code_ID = ?;

-- name: GetCropInspector :one
SELECT *
FROM Crop_Inspector
WHERE USDAID = ?;

-- name: UpdateCropInspector :execresult
UPDATE Crop_Inspector
SET Name = ?
WHERE USDAID = ?;


-- name: DeleteCropInspector :execresult
DELETE FROM Crop_Inspector 
WHERE USDAID = ?;

-- name: GetDistrictCodes :many
SELECT *
FROM District_Code;


-- name: GetDistrictCode :one
SELECT *
FROM District_Code
WHERE Code_ID = ?;

-- name: UpdateDistrictCode :execresult
UPDATE District_Code
SET Max_Water = ?, Max_Fert = ?,
    Crop_Type = ?, Code_ID = ?
WHERE Code_ID = ?;

-- name: DeleteDistrictCode :execresult
DELETE FROM District_Code
WHERE Code_ID = ?;