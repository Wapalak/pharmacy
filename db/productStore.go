package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"pharma"
)

type ProductStore struct {
	*sqlx.DB
}

func (s *ProductStore) ProductList() ([]pharma.Product, error) {
	var products []pharma.Product
	err := s.Select(&products, `SELECT p.Product_ID, p.Name, c.Name AS CategoryName 
		FROM Product p JOIN Category c ON p.Category_ID = c.Category_ID;`)

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductStore) CategorySave(categoryName string) error {
	stmt, err := s.Prepare(`INSERT INTO Category (Name) VALUES ($1);`)
	if err != nil {
		log.Fatal("can't prepare sql statement", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(categoryName)
	if err != nil {
		log.Print("error saving data", err)
		return err
	}
	return nil
}

func (s *ProductStore) ProductSave(t *pharma.Product) error {
	stmt, err := s.Prepare("INSERT INTO Product (Name, Category_ID) VALUES ($1, $2);")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Name, t.CategoryId)
	if err != nil {
		log.Print("error saving data", err)
		return err
	}
	return nil
}

//func (s *ProductStore) DeleteProduct(productID int) error {
//	// Определите SQL-запрос для удаления продукта по его ID
//	query := "DELETE FROM Product WHERE product_id = $1"
//
//	// Выполните SQL-запрос для удаления продукта
//	_, err := s.DB.Exec(query, productID)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

func (s *ProductStore) SuppliesList() ([]pharma.Supplies, error) {
	var supplies []pharma.Supplies
	query := `
		SELECT 
			s.supplies_id,
			p.pharmacy_id,
			p.name AS pharmacy_name,
			su.supplier_id,
			su.name AS supplier_name,
			pr.product_id,
			pr.name AS product_name,
			s.quantity
		FROM 
			Supplies s
			INNER JOIN Pharmacy p ON s.pharmacy_id = p.pharmacy_id
			INNER JOIN Supplier su ON s.supplier_id = su.supplier_id
			INNER JOIN Product pr ON s.product_id = pr.product_id;
	`
	err := s.Select(&supplies, query)
	if err != nil {
		return nil, err
	}
	return supplies, nil
}

func (s *ProductStore) SuppliesSave(supply *pharma.Supplies) error {
	query := `
		INSERT INTO Supplies (pharmacy_id, supplier_id, product_id, quantity)
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.Exec(query, supply.PharmacyId, supply.SupplierId, supply.ProductId, supply.Quantity)
	if err != nil {
		return err
	}
	return nil
}

// ProductStore.go
func (s *ProductStore) SuppliersList() ([]pharma.Supplier, error) {
	var suppliers []pharma.Supplier
	query := `
		SELECT supplier_id, name
		FROM Supplier
	`
	if err := s.Select(&suppliers, query); err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *ProductStore) GetPharmacies() ([]pharma.PharmacyInfo, error) {
	var pharmacies []pharma.PharmacyInfo
	query := `
        SELECT p.pharmacy_id, p.name, p.address, p.region_id, r.name AS region_name
        FROM pharmacy p
        JOIN region r ON p.region_id = r.region_id
    `
	if err := s.Select(&pharmacies, query); err != nil {
		return nil, err
	}
	return pharmacies, nil
}

func (s *ProductStore) GetOrders() ([]pharma.Order, error) {
	var orders []pharma.Order
	query := `
        SELECT o.order_id, o.product_id, p.name as product_name, o.pharmacy_id, ph.name as pharmacy_name, o.quantity
		FROM "Order" o
		JOIN Product p ON o.product_id = p.product_id
		JOIN Pharmacy ph ON o.pharmacy_id = ph.pharmacy_id;
    `
	if err := s.DB.Select(&orders, query); err != nil {
		return nil, err
	}
	return orders, nil
}
func (s *ProductStore) SaveOrders(order *pharma.Order) error {
	query := `
        INSERT INTO "Order" (pharmacy_id, product_id, quantity)
        VALUES ($1, $2, $3)
    `
	_, err := s.DB.Exec(query, order.PharmacyId, order.ProductId, order.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductStore) GetShippingData() ([]pharma.Shipping, error) {
	var shipping []pharma.Shipping
	query := `
       	SELECT s.delivery_id, s.order_id, s.stock_id, s.status
		FROM Shipping s
		JOIN "Order" o ON s.order_id = o.order_id
    `
	if err := s.DB.Select(&shipping, query); err != nil {
		return nil, err
	}
	return shipping, nil
}

func (s *ProductStore) GetInStockData() ([]pharma.InStock, error) {
	var inStock []pharma.InStock
	query := `
        SELECT i.stock_id, i.product_id, p.name AS name, i.quantity
        FROM in_stock i
        JOIN product p ON i.product_id = p.product_id
    `
	if err := s.DB.Select(&inStock, query); err != nil {
		return nil, err
	}
	return inStock, nil
}

func (s *ProductStore) GetRegions() ([]pharma.Region, error) {
	var region []pharma.Region
	query := `SELECT * from Region`

	if err := s.DB.Select(&region, query); err != nil {
		return nil, err
	}
	return region, nil
}
