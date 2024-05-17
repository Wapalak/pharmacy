package web

import (
	"html/template"
	"net/http"
	"pharma"
	"strconv"
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

func (h *Handler) ShippingList() http.HandlerFunc {
	type Data struct {
		Shipping []pharma.Shipping
	}
	tmpl := template.Must(template.New("shippingList.html").ParseFiles("C:\\Users\\User\\GolandProjects\\" +
		"pharma\\web\\templates\\shippingList.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.product.GetShippingData()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.ExecuteTemplate(w, "shippingList.html", Data{Shipping: tt})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) NewOrder() http.HandlerFunc {
	tmpl := template.Must(template.New("newOrder.html").ParseFiles("C:\\Users\\User\\" +
		"GolandProjects\\pharma\\web\\templates\\newOrder.html"))
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func (h *Handler) OrderSave() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Получаем данные заказа из формы
		productIDStr := r.FormValue("product_id")
		pharmacyIDStr := r.FormValue("pharmacy_id")
		quantityStr := r.FormValue("quantity")

		// Преобразуем данные в числа
		productID, err := strconv.Atoi(productIDStr)
		if err != nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		pharmacyID, err := strconv.Atoi(pharmacyIDStr)
		if err != nil {
			http.Error(w, "Invalid pharmacy ID", http.StatusBadRequest)
			return
		}

		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			http.Error(w, "Invalid quantity", http.StatusBadRequest)
			return
		}

		// Создаем объект заказа
		order := &pharma.Order{
			ProductId:  productID,
			PharmacyId: pharmacyID,
			Quantity:   quantity,
		}

		// Сохраняем заказ
		if err := h.product.SaveOrders(order); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Перенаправляем пользователя на страницу со списком продуктов
		http.Redirect(w, r, "/orders", http.StatusFound)
	}
}
