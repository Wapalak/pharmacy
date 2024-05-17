package pharma

type Pharmacy interface {
	PharmacyStore
}

type PharmacyStore interface {
	ProductList() ([]Product, error)
	CategorySave(categoryName string) error
	ProductSave(t *Product) error
	SuppliesList() ([]Supplies, error)
	SuppliesSave(supply *Supplies) error
	SuppliersList() ([]Supplier, error)
	GetPharmacies() ([]PharmacyInfo, error)
	GetOrders() ([]Order, error)
	GetShippingData() ([]Shipping, error)
	GetInStockData() ([]InStock, error)
	GetRegions() ([]Region, error)
}

type Product struct {
	ProductId    int    `db:"product_id"`
	Name         string `db:"name"`
	CategoryId   int    `db:"category_id"`
	CategoryName string `db:"categoryname"` // Проверьте, что здесь указано правильное имя столбца
}

type Category struct {
	CategoryId int    `db:"category_id"`
	Name       string `db:"name"`
}

type Order struct {
	OrderId      int    `db:"order_id"`
	ProductId    int    `db:"product_id"`
	Name         string `db:"name"`
	PharmacyId   int    `db:"pharmacy_id"`
	PharmacyName string `db:"pharmacy_name"`
	Quantity     int    `db:"quantity"`
}

type InStock struct {
	StockId   int    `db:"stock_id"`
	ProductId int    `db:"product_id"`
	Name      string `db:"name"`
	Quantity  int    `db:"quantity"`
}

type Shipping struct {
	DeliveryId int    `db:"delivery_id"`
	OrderId    int    `db:"order_id"`
	ProductId  int    `db:"product_id"`
	Name       string `db:"name"`
	StockId    int    `db:"stock_id"`
	Quantity   int    `db:"quantity"`
	Status     string `db:"status"`
}

type Supplies struct {
	SuppliesId   int    `db:"supplies_id"`
	PharmacyId   int    `db:"pharmacy_id"`
	PharmacyName string `db:"pharmacy_name"`
	SupplierId   int    `db:"supplier_id"`
	SupplierName string `db:"supplier_name"`
	ProductId    int    `db:"product_id"`
	ProductName  string `db:"product_name"`
	Quantity     int    `db:"quantity"`
}

type Supplier struct {
	SupplierId int    `db:"supplier_id"`
	Name       string `db:"name"`
}

type PharmacyInfo struct {
	PharmacyId int    `db:"pharmacy_id"`
	Name       string `db:"name"`
	Address    string `db:"address"`
	RegionId   int    `db:"region_id"`
	RegionName string `db:"region_name"`
}

type Region struct {
	RegionId int    `db:"region_id"`
	Name     string `db:"name"`
}

type Employee struct {
	EmployeeId int    `db:"employee_id"`
	Name       string `db:"name"`
	Position   string `db:"position"`
	PharmacyId int    `db:"pharmacy_id"`
}

type Contract struct {
	ContractId int    `db:"contract_id"`
	EmployeeId int    `db:"employee_id"`
	DateOfHire string `db:"date_of_hire"`
	Duration   string `db:"duration"`
	Salary     int    `db:"salary"`
}
