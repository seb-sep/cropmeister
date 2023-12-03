CREATE TABLE IF NOT EXISTS Farm
(
  Name           VARCHAR(255) NOT NULL,
  Farm_Value      INT,
  Farm_ID         INT,
  Address_Street VARCHAR(255),
  Address_City   VARCHAR(255),
  Address_State  VARCHAR(255),
  Address_Zip    VARCHAR(10),
  PRIMARY KEY (Farm_ID)
);

CREATE TABLE IF NOT EXISTS Crop
(
  Crop_Type            VARCHAR(255),
  Ph_Range_Weight      DECIMAL(5, 2),
  Ph_Range_Desired     DECIMAL(5, 2),
  Water_Needed_Weight  DECIMAL(5, 2),
  Water_Needed_Desired DECIMAL(5, 2),
  Sun_Range_Weight     DECIMAL(5, 2),
  SunRange_Desired    DECIMAL(5, 2),
  Banned	      BOOLEAN,
  PRIMARY KEY (Crop_Type)
);

CREATE TABLE IF NOT EXISTS Purchase
(
  Purchase_ID       INT AUTO_INCREMENT,
  Crop_Type         VARCHAR(255),
  Purchase_Complete BOOLEAN,
  Total_Price       DECIMAL(5, 2),
  Total_Quantity    INT,
  Purchase_Date     DATE,
  PRIMARY KEY (Purchase_ID),
  FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
);

CREATE TABLE IF NOT EXISTS Crop_Investor
(
  Name            VARCHAR(255) NOT NULL,
  Buy_Price        INT,
  Investible_Money INT,
  Sell_Price       INT,
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
  Max_Water INT,
  Max_Fert  INT,
  Crop_Type VARCHAR(255),
  Code_ID   INT,
  PRIMARY KEY (Code_ID),
  FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
);

CREATE TABLE IF NOT EXISTS Crop_Buyer
(
  Name                VARCHAR(255),
  Quantities_Required INT,
  Crop_Type            VARCHAR(255),
  Target_Price         INT,
  PRIMARY KEY (Name),
  FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
);

CREATE TABLE IF NOT EXISTS Farmer
(
  Name          VARCHAR(255) NOT NULL,
  Budget        INT,
  Net_Worth      INT,
  Farm_ID        INT,
  Purchase_ID    INT,
  -- PRIMARY KEY (Name),
  FOREIGN KEY (Farm_ID) REFERENCES Farm (Farm_ID),
  FOREIGN KEY (Purchase_ID) REFERENCES Purchase (Purchase_ID)
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
  Crop_Type        VARCHAR(255),
  Farm_ID          INT,
  Extinct         BOOLEAN,
  FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type),
  FOREIGN KEY (Farm_ID) REFERENCES Farm (Farm_ID)
);

CREATE TABLE IF NOT EXISTS Monitors_Buyer
(
  Name     VARCHAR(255) NOT NULL,
  Crop_Type VARCHAR(255) NOT NULL,
  FOREIGN KEY (Name) REFERENCES Crop_Buyer (Name),
  FOREIGN KEY (Crop_Type) REFERENCES Crop (Crop_Type)
);

CREATE TABLE IF NOT EXISTS Invests_In
(
  Name   VARCHAR(255),
  Farm_ID INT,
  -- PRIMARY KEY (Name),
  FOREIGN KEY (Name) REFERENCES Crop_Investor (Name),
  FOREIGN KEY (Farm_ID) REFERENCES Farm(Farm_ID)
);

CREATE TABLE IF NOT EXISTS Enforces
(
  USDAID INT NOT NULL,
  Code_ID INT NOT NULL,
  -- PRIMARY KEY (USDAID, CodeID),
  FOREIGN KEY (USDAID) REFERENCES Crop_Investigator (USDAID),
  FOREIGN KEY (Code_ID) REFERENCES District_Code (Code_ID)
);

CREATE TABLE IF NOT EXISTS Monitors_Investments
(
  Name     VARCHAR(255),
  Crop_Type VARCHAR(255),
  -- PRIMARY KEY (Name),
  FOREIGN KEY (Name) REFERENCES Crop_Investor (Name),
  FOREIGN KEY (Crop_Type) REFERENCES Harvest (Crop_Type)
);