-- Get Harvest
SELECT * FROM Harvest;

-- Post Harvest
INSERT INTO Harvest (Quantity, Harvest_Date, Ph_Base, Ph_Fertilized, Water_Rain, Water_Sprinkler, Sun, Price,
                     Crop_Type, Farm_ID, Extinct)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- Get Crop {Crop_Type}
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

-- Put Crop {Crop_Type}
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

-- Delete Crop {Crop+Type}
UPDATE Crop
SET Banned = TRUE
WHERE Crop_Type = ?;

-- Get Farm
SELECT * FROM Farm;

-- Post Farm
INSERT INTO Farm (Name, Farm_Value, Farm_ID, Address_Street, Address_City, Address_State, Address_Zip)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- Get Farm/{Farm_ID}
SELECT * FROM Farm
WHERE Farm_ID = ?;

-- Put Farm/{Farm_ID}
UPDATE Farm
SET
    Name = ?,
    Farm_Value = ?,
    Address_Street = ?,
    Address_City = ?,
    Address_State = ?,
    Address_Zip = ?
WHERE Farm_ID = ?;

-- Delete Farm/{Farm_ID}
UPDATE Farm
SET Active = FALSE
WHERE Farm_ID = ?;

-- Get Harvest/{Crop_Type}
SELECT * FROM Harvest
WHERE Crop_Type = ?;

-- Put Harvest/{Crop_Type}
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

-- Delete Harvest/{Crop_Type}
UPDATE Harvest
SET Extinct = TRUE
WHERE Crop_Type = ?;

-- Get Purchase
SELECT * FROM Purchase;

-- Post Purchase
INSERT INTO Purchase (Crop_Type, Farm_ID, Purchase_Complete, Total_Price, Total_Quantity, Purchase_Date)
VALUES (?, ?, ?, ?, ?, ?);

-- Get Purchase/{Purchase_ID}
SELECT * FROM Purchase
WHERE Purchase_ID = ?;

-- Put Purchase/{Purchase_ID}
UPDATE Purchase
SET
    Crop_Type = ?, 
    Farm_ID = ?,
    Purchase_Complete = ?,
    Total_Price = ?,
    Total_Quantity = ?,
    Purchase_Date = ?
WHERE Purchase_ID = ?;

-- Delete Purchase/{Purchase_ID}
UPDATE Purchase
SET Canceled = TRUE
WHERE Purchase_ID = ?;

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

-- /Farmer/{Farm_ID}
-- Get
SELECT *
FROM Farmer
WHERE Farm_ID = ?;

-- Post
INSERT INTO Farmer (Name, Budget, Net_Worth, Farm_ID)
VALUES (?, ?, ?, ?);

-- /CropInspector
-- Get
SELECT *
FROM Crop_Investigator;
-- Post
INSERT INTO Crop_Investigator (Name, USDAID)
VALUES (?, ?);

-- /DistinctCode/Crop/{Crop_Type}
-- Get
SELECT *
FROM District_Code
WHERE Crop_Type = ?;

-- Post
INSERT INTO District_Code (Max_Water, Max_Fert, Crop_Type, Code_ID)   
VALUES (?, ?, ?, ?);

-- /DistrictCode/Inspector/{USDAID}
-- Get
SELECT *
FROM District_Code   
WHERE (SELECT * 
FROM Crop_Investigator
WHERE USDAID = ?);
                 
-- Post
INSERT INTO District_Code (Max_Water, Max_Fert, Crop_Type, Code_ID)
VALUES (?, ?, ?, ?);

-- /CropInspector/Code/{Code_ID}
-- Get
SELECT *
FROM Crop_Investigator
WHERE farm.District_Code.Code_ID = ?;
-- Post
INSERT INTO Crop_Investigator (Name, USDAID) 
VALUES (?, ?);

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
INSERT INTO District_Code (Max_Water, Max_Fert, Crop_Type, Code_ID)
VALUES (?, ?, ?, ?);

-- DistrictCode/Code/{Code_ID}
-- Get
SELECT *
FROM District_Code
WHERE Code_ID = ?;
-- Put
UPDATE District_Code
SET Max_Water = ?, Max_Fert = ?,
    Crop_Type = ?, Code_ID = ?
WHERE Code_ID = ?;
-- Delete
DELETE FROM District_Code
WHERE Code_ID = ?;