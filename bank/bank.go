package bank

import (
	_ "github.com/Ad3bay0c/mortgage_app/db"
	"net/http"
)

type Bank struct {
	Name		string	`json:"name,omitempty" bson:"name,omitempty"`
	Interest	float64	`json:"interest,omitempty" bson:"interest,omitempty"`
	MaxLoan		float64	`json:"max_loan,omitempty" bson:"max_loan,omitempty"`
	MinDown		float64	`json:"min_down,omitempty" bson:"min_down,omitempty"`
	LoanTerm	float64	`json:"loan_term,omitempty" bson:"loan_term,omitempty"`
	CreatedAt	int		`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt	int		`json:"update_at,omitempty" bson:"update_at,omitempty"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Endpoint Connected"))
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bank List Endpoint Created"))
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Bank"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Bank"))
}