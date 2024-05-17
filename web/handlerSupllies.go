package web

import (
	"html/template"
	"net/http"
	"pharma"
	"strconv"
)

func (h *Handler) SuppliesList() http.HandlerFunc {
	type Data struct {
		Supplies []pharma.Supplies
	}

	tmpl := template.Must(template.New("suppliesList.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"pharma\\web\\templates\\suppliesList.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.product.SuppliesList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "suppliesList.html", Data{Supplies: tt})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) NewSupplies() http.HandlerFunc {
	tmpl := template.Must(template.New("suppliesNew.html").ParseFiles("C:\\Users\\User\\" +
		"GolandProjects\\pharma\\web\\templates\\suppliesNew.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) SuppliesSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		pharmacyIdStr := r.FormValue("pharmacy_id")
		supplierIdStr := r.FormValue("supplier_id")
		productIdStr := r.FormValue("product_id")
		quantityStr := r.FormValue("quantity")

		pharmacyId, err := strconv.Atoi(pharmacyIdStr)
		if err != nil {
			http.Error(w, "Invalid pharmacy ID", http.StatusBadRequest)
			return
		}

		supplierId, err := strconv.Atoi(supplierIdStr)
		if err != nil {
			http.Error(w, "Invalid supplier ID", http.StatusBadRequest)
			return
		}

		productId, err := strconv.Atoi(productIdStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		supply := &pharma.Supplies{
			PharmacyId: pharmacyId,
			SupplierId: supplierId,
			ProductId:  productId,
			Quantity:   quantity,
		}

		if err := h.product.SuppliesSave(supply); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/supplies/", http.StatusFound)
	}
}

func (h *Handler) SuppliersList() http.HandlerFunc {
	type Data struct {
		Supplier []pharma.Supplier
	}

	tmpl := template.Must(template.New("supplierList.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"pharma\\web\\templates\\supplierList.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		supplier, err := h.product.SuppliersList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "supplierList.html", Data{Supplier: supplier})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
