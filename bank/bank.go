package bank

import (
	"context"
	"encoding/json"
	"github.com/Ad3bay0c/mortgage_app/db"
	_ "github.com/Ad3bay0c/mortgage_app/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

type Bank struct {
	ID			string	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name		string	`json:"name,omitempty" bson:"name,omitempty"`
	Interest	float64	`json:"interest,omitempty" bson:"interest,omitempty"`
	MaxLoan		float64	`json:"max_loan,omitempty" bson:"max_loan,omitempty"`
	MinDown		float64	`json:"min_down,omitempty" bson:"min_down,omitempty"`
	LoanTerm	float64	`json:"loan_term,omitempty" bson:"loan_term,omitempty"`
	CreatedAt	int64	`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt	int64	`json:"update_at,omitempty" bson:"update_at,omitempty"`
}

type Message struct {
	Message		string		`json:"message,omitempty" bson:"message,omitempty"`
	Data		interface{}	`json:"data,omitempty" bson:"data,omitempty"`
}

var collection = db.Client.Database("ContactKeeper").Collection("bank")

func Create(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json")

	 var bank *Bank

	 _ = json.NewDecoder(r.Body).Decode(&bank)
	bank.CreatedAt = time.Now().Unix()
	bank.UpdateAt = time.Now().Unix()

	 ctx, cancelFunc := context.WithTimeout(context.Background(), 10 * time.Minute)
	 cancelFunc()
	 result, err := collection.InsertOne(ctx, bank)
	 if err != nil {
		 w.WriteHeader(http.StatusInternalServerError)
		 log.Printf("%v", err.Error())
		 json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		 return
	 }

	 w.WriteHeader(http.StatusOK)
	 json.NewEncoder(w).Encode(Message{Message: "Successful", Data: result.InsertedID})

}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var banks []Bank

	ctx, _:= context.WithTimeout(context.Background(), 10 * time.Minute)
	result, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}

	for result.Next(ctx) {
		var bank Bank

		_ = result.Decode(&bank)
		banks = append(banks, bank)
	}

	w.WriteHeader(http.StatusOK)
	if len(banks) == 0 {
		json.NewEncoder(w).Encode(Message{Message: "Empty List", Data: nil})
		return
	}
	json.NewEncoder(w).Encode(Message{Message: "Success", Data: banks})
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Bank"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Bank"))
}