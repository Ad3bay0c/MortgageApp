package bank

import (
	_ "github.com/Ad3bay0c/mortgage_app/db"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Endpoint Connected"))
}

//func BankList(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Bank List Endpoint Created"))
//}
//
//func Update(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Update Bank"))
//}
//
//func Delete(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Delete Bank"))
//}
