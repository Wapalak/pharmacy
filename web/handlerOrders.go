package web

import (
	"html/template"
	"net/http"
	"pharma"
)

func (h *Handler) OrdersList() http.HandlerFunc {
	type Data struct {
		Orders []pharma.Order
	}

	tmpl := template.Must(template.New("ordersList.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"pharma\\web\\templates\\ordersList.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.product.GetOrders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "ordersList.html", Data{Orders: tt})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
