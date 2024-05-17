package web

import (
	"html/template"
	"net/http"
	"pharma"
	"strconv"
)

func (h *Handler) HelloPage() http.HandlerFunc {
	tmpl := template.Must(template.New("helloPage.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"pharma\\web\\templates\\helloPage.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := tmpl.ExecuteTemplate(w, "helloPage.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handler) ProductsPage() http.HandlerFunc {
	type Data struct {
		Products []pharma.Product
	}

	tmpl := template.Must(template.New("productList.html").ParseFiles("C:\\Users\\User\\" +
		"GolandProjects\\pharma\\web\\templates\\productList.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.product.ProductList()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Проходим по каждому продукту и присваиваем полю CategoryName значение из структуры Product
		//for i := range tt {
		//	tt[i].CategoryName = tt[i].Category.Name
		//}

		err = tmpl.ExecuteTemplate(w, "productList.html", Data{Products: tt})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) NewCategory() http.HandlerFunc {
	tmpl := template.Must(template.New("newCategory.html").ParseFiles("C:\\Users\\User\\" +
		"GolandProjects\\pharma\\web\\templates\\newCategory.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) CategorySave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := r.FormValue("category")
		if err := h.product.CategorySave(category); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/products/list", http.StatusFound)
	}
}

func (h *Handler) NewProduct() http.HandlerFunc {
	tmpl := template.Must(template.New("newProduct.html").ParseFiles("C:\\Users\\User\\" +
		"GolandProjects\\pharma\\web\\templates\\newProduct.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) ProductSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		category_idSTR := r.FormValue("category_id")
		category_id, err := strconv.Atoi(category_idSTR)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}
		if err := h.product.ProductSave(&pharma.Product{
			Name:       name,
			CategoryId: category_id,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/products/list", http.StatusFound)
	}
}
