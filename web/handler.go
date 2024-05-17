package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"pharma"
)

type Handler struct {
	*chi.Mux

	product      pharma.Pharmacy
	category     pharma.Category
	supplies     pharma.Supplies
	pharmacyinfo pharma.PharmacyInfo
	orders       pharma.Order
}

func NewHandler(
	product pharma.Pharmacy, category pharma.Category, supplies pharma.Supplies,
	pharmacyinfo pharma.PharmacyInfo, orders pharma.Order) *Handler {
	h := &Handler{Mux: chi.NewMux(),
		product:      product,
		category:     category,
		supplies:     supplies,
		pharmacyinfo: pharmacyinfo,
		orders:       orders,
	}

	h.Use(middleware.Logger)
	h.Route("/products", func(r chi.Router) {
		r.Get("/", h.HelloPage())
		r.Get("/list", h.ProductsPage())
		r.Get("/list/newCategory", h.NewCategory())
		// Маршруты для создания новой категории
		r.Post("/list/newCategory", h.CategorySave())
		// Маршруты для создания нового продукта
		r.Get("/list/newProduct", h.NewProduct())
		r.Get("/list/inStock", h.inStockList())
		r.Post("/list/newProduct", h.ProductSave())
	})

	h.Route("/supplies", func(r chi.Router) {
		r.Get("/", h.SuppliesList())
		r.Get("/newSupplies", h.NewSupplies())
		r.Post("/", h.SuppliesSave())
		r.Get("/suppliers", h.SuppliersList())
	})

	h.Route("/pharmacy", func(r chi.Router) {
		r.Get("/", h.GetPharmacies())
		r.Get("/newSupplies", h.NewSupplies())
		r.Post("/", h.SuppliesSave())
		r.Get("/suppliers", h.SuppliersList())
	})

	h.Route("/orders", func(r chi.Router) {
		r.Get("/", h.OrdersList())
		r.Get("/shippings", h.ShippingList())
		r.Post("/", h.SuppliesSave())
	})

	return h
}
