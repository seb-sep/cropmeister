DROP DATABASE IF EXISTS farm;
CREATE DATABASE IF NOT EXISTS farm;
USE farm;




CREATE TABLE IF NOT EXISTS Farm
(
  Name           VARCHAR(255) NOT NULL,
  FarmValue      INT,
  FarmID         INT,
  Address_Street VARCHAR(255),
  Address_City   VARCHAR(255),
  Address_State  VARCHAR(255),
  Address_Zip    VARCHAR(10),
  Active         BOOLEAN,
  PRIMARY KEY (FarmID)
);




CREATE TABLE IF NOT EXISTS Crop
(
  BasePrice           Real,
  CropType            VARCHAR(255),
  PhRange_Weight      REAL,
  PhRange_Desired     Real,
  WaterNeeded_Weight  Real,
  WaterNeeded_Desired Real,
  SunRange_Weight     Real,
  SunRange_Desired    Real,
  Restricted	      BOOLEAN,
  Restricted_Penalty  REAL,
  PRIMARY KEY (CropType)
);




CREATE TABLE IF NOT EXISTS Purchase
(
  PurchaseID       INT AUTO_INCREMENT,
  CropType         VARCHAR(255),
  PurchaseComplete BOOLEAN,
  TotalPrice       Real,
  TotalQuantity    INT,
  PurchaseDate     DATE,
  Canceled         BOOLEAN, 
  PRIMARY KEY (PurchaseID),
  FOREIGN KEY (CropType) REFERENCES Crop (CropType)
);




CREATE TABLE IF NOT EXISTS Crop_Investor
(
  Name            VARCHAR(255) NOT NULL,
  BuyPrice        INT,
  InvestibleMoney INT,
  SellPrice       INT,
  PRIMARY KEY (Name)
);




CREATE TABLE IF NOT EXISTS Crop_Investigator
(
  Name   VARCHAR(255) NOT NULL,
  USDAID INT,
  PRIMARY KEY (USDAID)
);




CREATE TABLE IF NOT EXISTS District_Code
(
  MaxWater INT,
  MaxFert  INT,
  CropType VARCHAR(255),
  CodeID   INT,
  PRIMARY KEY (CodeID),
  FOREIGN KEY (CropType) REFERENCES Crop (CropType)
);




CREATE TABLE IF NOT EXISTS Crop_Buyer
(
  Name                VARCHAR(255),
  Quantities_Required INT,
  CropType            VARCHAR(255),
  TargetPrice         INT,
  PRIMARY KEY (Name),
  FOREIGN KEY (CropType) REFERENCES Crop (CropType)
);




CREATE TABLE IF NOT EXISTS Farmer
(
  Name          VARCHAR(255) NOT NULL,
  Budget        INT,
  NetWorth      INT,
  FarmID        INT,
  PurchaseID    INT,
  -- PRIMARY KEY (Name),
  FOREIGN KEY (FarmID) REFERENCES Farm (FarmID),
  FOREIGN KEY (PurchaseID) REFERENCES Purchase (PurchaseID)
);




CREATE TABLE IF NOT EXISTS Harvest
(
  Quantity        INT,
  Time_Year       YEAR,
  Time_Season     VARCHAR(255),
  Ph_Base         Real,
  Ph_Fertilized   Real,
  Water_Rain      Real,
  Water_Sprinkler Real,
  Sun             INT,
  Price           Real,
  CropType        VARCHAR(255),
  FarmID          INT,
  Extinct         BOOLEAN,
  -- PRIMARY KEY (CropType, FarmID),
  FOREIGN KEY (CropType) REFERENCES Crop (CropType),
  FOREIGN KEY (FarmID) REFERENCES Farm (FarmID)
);




CREATE TABLE IF NOT EXISTS Monitors_Buyer
(
  Name     VARCHAR(255) NOT NULL,
  CropType VARCHAR(255) NOT NULL,
  -- PRIMARY KEY (Name, CropType),
  FOREIGN KEY (Name) REFERENCES Crop_Buyer (Name),
  FOREIGN KEY (CropType) REFERENCES Crop (CropType)
);




CREATE TABLE IF NOT EXISTS Invests_In
(
  Name   VARCHAR(255),
  FarmID INT,
  -- PRIMARY KEY (Name),
  FOREIGN KEY (Name) REFERENCES Crop_Investor (Name),
  FOREIGN KEY (FarmID) REFERENCES Farm(FarmID)
);




CREATE TABLE IF NOT EXISTS Enforces
(
  USDAID INT NOT NULL,
  CodeID INT NOT NULL,
  -- PRIMARY KEY (USDAID, CodeID),
  FOREIGN KEY (USDAID) REFERENCES Crop_Investigator (USDAID),
  FOREIGN KEY (CodeID) REFERENCES District_Code (CodeID)
);




CREATE TABLE IF NOT EXISTS Monitors_Investments
(
  Name     VARCHAR(255),
  CropType VARCHAR(255),
  -- PRIMARY KEY (Name),
  FOREIGN KEY (Name) REFERENCES Crop_Investor (Name),
  FOREIGN KEY (CropType) REFERENCES Harvest (CropType)
);
