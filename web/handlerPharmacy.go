package web

import (
	"html/template"
	"net/http"
	"pharma"
)

func (h *Handler) GetPharmacies() http.HandlerFunc {
	type Data struct {
		Pharmacies []pharma.PharmacyInfo
		Regions    []pharma.Region
	}

	tmpl := template.Must(template.New("pharmacyList.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"pharma\\web\\templates\\pharmacyList.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.product.GetPharmacies()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		regions, err := h.product.GetRegions()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "pharmacyList.html", Data{Pharmacies: tt, Regions: regions})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
