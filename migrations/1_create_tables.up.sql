CREATE TABLE Region
(
    Region_ID SERIAL PRIMARY KEY,
    Name      TEXT NOT NULL
);

CREATE TABLE Pharmacy
(
    Pharmacy_ID SERIAL PRIMARY KEY,
    Name        TEXT    NOT NULL,
    Address     TEXT    NOT NULL,
    Region_ID   INTEGER NOT NULL REFERENCES Region (Region_ID)
);

CREATE TABLE Category
(
    Category_ID SERIAL PRIMARY KEY,
    Name        TEXT NOT NULL
);

CREATE TABLE Product
(
    Product_ID  SERIAL PRIMARY KEY,
    Name        TEXT    NOT NULL,
    Category_ID INTEGER NOT NULL REFERENCES Category (Category_ID)
);

CREATE TABLE Employee
(
    Employee_ID SERIAL PRIMARY KEY,
    Name        TEXT    NOT NULL,
    Position    TEXT    NOT NULL,
    Pharmacy_ID INTEGER NOT NULL REFERENCES Pharmacy (Pharmacy_ID)
);

CREATE TABLE Contract
(
    Contract_ID  SERIAL PRIMARY KEY,
    Employee_ID  INTEGER NOT NULL REFERENCES Employee (Employee_ID),
    Date_of_hire DATE    NOT NULL,
    Duration INTERVAL,
    Salary       REAL    NOT NULL
);

CREATE TABLE Customer
(
    Customer_ID SERIAL PRIMARY KEY,
    Name        TEXT NOT NULL,
    PhoneNumber TEXT NOT NULL,
    Email       TEXT NOT NULL
);

CREATE TABLE In_Stock
(
    Stock_ID   SERIAL PRIMARY KEY,
    Product_ID INTEGER NOT NULL REFERENCES Product (Product_ID),
    Quantity   INTEGER NOT NULL
);

CREATE TABLE "Order"
(
    Order_ID    SERIAL PRIMARY KEY,
    Pharmacy_ID INTEGER NOT NULL REFERENCES Pharmacy (Pharmacy_ID),
    Product_ID  INTEGER NOT NULL REFERENCES Product (Product_ID),
    Quantity    INTEGER NOT NULL
);

CREATE TABLE Shipping
(
    Delivery_ID SERIAL PRIMARY KEY,
    Order_ID    INTEGER NOT NULL REFERENCES "Order" (Order_ID),
    Stock_ID    INTEGER NOT NULL REFERENCES In_Stock (Stock_ID),
    Status      TEXT    NOT NULL
);

CREATE TABLE Supplier
(
    Supplier_ID SERIAL PRIMARY KEY,
    Name        TEXT NOT NULL
);

CREATE TABLE Supplies
(
    Supplies_ID SERIAL PRIMARY KEY,
    Pharmacy_ID INTEGER NOT NULL REFERENCES Pharmacy (Pharmacy_ID),
    Supplier_ID INTEGER NOT NULL REFERENCES Supplier (Supplier_ID),
    Product_ID  INTEGER NOT NULL REFERENCES Product (Product_ID),
    Quantity    INTEGER NOT NULL
);
