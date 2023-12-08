create database farms;
CREATE TABLE IF NOT EXISTS Farm
(
  Name           VARCHAR(255) NOT NULL,
  Farm_Value      INT,
  Farm_ID         INT PRIMARY KEY AUTO_INCREMENT,
  Address_Street VARCHAR(255),
  Address_City   VARCHAR(255),
  Address_State  VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS Crop
(
  Crop_Type            VARCHAR(255) PRIMARY KEY,
  Ph_Range_Weight      REAL,
  Ph_Range_Desired     REAL,
  Water_Needed_Weight  REAL,
  Water_Needed_Desired REAL,
  Sun_Range_Weight     REAL,
  Sun_Range_Desired    REAL,
  Base_Price           REAL,
  Banned	      BOOLEAN
);

CREATE TABLE IF NOT EXISTS Purchase
(
  Purchase_ID       INT PRIMARY KEY AUTO_INCREMENT,
  Crop_Type         VARCHAR(255),
  Farm_ID           INT,
  Purchase_Complete BOOLEAN,
  Total_Price       REAL,
  Total_Quantity    INT,
  Purchase_Date     DATE,
  Farmer_Name VARCHAR(255) NOT NULL,
  CONSTRAINT purchase_crop 
    FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
    ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT purchase_farm 
    FOREIGN KEY (Farm_ID, Farmer_Name) REFERENCES Farmer (Farm_ID, Name)
    ON DELETE NO ACTION ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Crop_Investor
(
  Name            VARCHAR(255) PRIMARY KEY,
  Buy_Price        REAL,
  Investible_Money REAL,
  Sell_Price       REAL,
  Unemployed       BOOLEAN
);

CREATE TABLE IF NOT EXISTS Crop_Inspector
(
  Name   VARCHAR(255) NOT NULL,
  USDAID INT PRIMARY KEY AUTO_INCREMENT
);

CREATE TABLE IF NOT EXISTS District_Code
(
  Code_ID   INT PRIMARY KEY AUTO_INCREMENT,
  Max_Water REAL,
  Max_Fert  REAL,
  Crop_Type VARCHAR(255),
  CONSTRAINT district_crop 
    FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
    ON DELETE NO ACTION ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Crop_Buyer
(
  Name                VARCHAR(255) PRIMARY KEY,
  Quantities_Required INT,
  Crop_Type            VARCHAR(255),
  Target_Price         REAL,
  CONSTRAINT buyer_crop 
    FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
    ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Farmer
(
  Name          VARCHAR(255) NOT NULL,
  Budget        REAL,
  Net_Worth      REAL,
  Farm_ID        INT NOT NULL,
  PRIMARY KEY (Farm_ID, Name),
  CONSTRAINT farmer_crop 
    FOREIGN KEY (Farm_ID) REFERENCES Farm (Farm_ID)
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Harvest
(
  Quantity        INT,
  Harvest_Year    YEAR NOT NULL,
  Ph_Base         REAL,
  Ph_Fertilized   REAL,
  Water_Rain      REAL,
  Water_Sprinkler REAL,
  Sun             INT,
  Price           REAL,
  Crop_Type       VARCHAR(255) NOT NULL,
  Farm_ID         INT NOT NULL,
  Extinct         BOOLEAN,
  PRIMARY KEY (Farm_ID, Crop_Type, Harvest_Year),
  CONSTRAINT harvest_crop 
    FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
    ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT harvest_farm 
    FOREIGN KEY (Farm_ID) REFERENCES Farm (Farm_ID)
    ON DELETE NO ACTION ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Monitors_Buyer
(
  Name     VARCHAR(255) NOT NULL,
  Crop_Type VARCHAR(255) NOT NULL,
  CONSTRAINT buyer 
    FOREIGN KEY (Name) REFERENCES Crop_Buyer (Name)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT crop 
    FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Invests_In
(
  Name   VARCHAR(255),
  Farm_ID INT,
  CONSTRAINT investor 
    FOREIGN KEY (Name) REFERENCES Crop_Investor (Name)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT farm
    FOREIGN KEY (Farm_ID) REFERENCES Farm(Farm_ID)
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Enforces
(
  USDAID INT NOT NULL,
  Code_ID INT NOT NULL,
  CONSTRAINT inspector
    FOREIGN KEY (USDAID) REFERENCES Crop_Inspector (USDAID)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT district
    FOREIGN KEY (Code_ID) REFERENCES District_Code (Code_ID)
    ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS Monitors_Investments
(
  Name     VARCHAR(255),
  Farm_ID INT,
  Crop_Type VARCHAR(255),
  Harvest_Year YEAR,
  CONSTRAINT investor
    FOREIGN KEY (Name) REFERENCES Crop_Investor (Name)
    ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT crop
    FOREIGN KEY (Farm_ID, Crop_Type, Harvest_Year) REFERENCES Harvest (Farm_ID, Crop_Type, Harvest_Year)
    ON DELETE CASCADE ON UPDATE CASCADE
);