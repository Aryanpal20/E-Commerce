package routings

import (
	ct "e-Commerce/Admin/adminchangescategoryvalue"
	st "e-Commerce/Admin/adminchangestorevalue"
	cate "e-Commerce/Admin/adminfixcategory"
	sto "e-Commerce/Admin/adminfixstore"
	login "e-Commerce/Auth/login"
	register "e-Commerce/Auth/register"
	cust1 "e-Commerce/Customer/customerchange"
	cust "e-Commerce/Customer/customerselect"
	prod "e-Commerce/Vendor/vendorchangesproductvalue"
	pro "e-Commerce/Vendor/vendorfixproduct"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	r := mux.NewRouter()

	r.HandleFunc("/register", register.Register).Methods("POST")
	r.HandleFunc("/login", login.Login).Methods("POST")
	r.HandleFunc("/post/product", pro.ProductVendorAccess).Methods("POST")
	r.HandleFunc("/post/store", sto.StoreAminAccess).Methods("POST")
	r.HandleFunc("/post/category", cate.CategoryAminAccess).Methods("POST")
	r.HandleFunc("/update/store/{vendor_id}", st.StoreUpdateByAdmin).Methods("PUT")
	r.HandleFunc("/update/category/{store_id}", ct.CategoryUpdateByAdmin).Methods("PUT")
	r.HandleFunc("/update/product/{category_id}", prod.ProductUpdateByVendor).Methods("PUT")
	r.HandleFunc("/post/customer", cust.Customer_Select).Methods("POST")
	r.HandleFunc("/update/customer", cust1.CustomerChange).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))

}
