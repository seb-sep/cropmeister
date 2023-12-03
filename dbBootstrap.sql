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
  PRIMARY KEY (FarmID)
);




CREATE TABLE IF NOT EXISTS Crop
(
  CropType            VARCHAR(255),
  PhRange_Weight      DECIMAL(5, 2),
  PhRange_Desired     DECIMAL(5, 2),
  WaterNeeded_Weight  DECIMAL(5, 2),
  WaterNeeded_Desired DECIMAL(5, 2),
  SunRange_Weight     DECIMAL(5, 2),
  SunRange_Desired    DECIMAL(5, 2),
  Banned	      BOOLEAN,
  PRIMARY KEY (CropType)
);




CREATE TABLE IF NOT EXISTS Purchase
(
  PurchaseID       INT AUTO_INCREMENT,
  CropType         VARCHAR(255),
  PurchaseComplete BOOLEAN,
  TotalPrice       DECIMAL(5, 2),
  TotalQuantity    INT,
  PurchaseDate     DATE,
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
  Ph_Base         DECIMAL(5, 2),
  Ph_Fertilized   DECIMAL(5, 2),
  Water_Rain      DECIMAL(5, 2),
  Water_Sprinkler DECIMAL(5, 2),
  Sun             INT,
  Price           DECIMAL(5, 2),
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
