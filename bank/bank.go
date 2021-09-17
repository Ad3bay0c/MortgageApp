package bank

import (
	"context"
	"encoding/json"
	"github.com/Ad3bay0c/mortgage_app/db"
	_ "github.com/Ad3bay0c/mortgage_app/db"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type Bank struct {
	ID			primitive.ObjectID	`json:"_id,omitempty" bson:"_id,omitempty"`
	Name		string				`json:"name,omitempty" bson:"name,omitempty"`
	Interest	float64				`json:"interest" bson:"interest"`
	MaxLoan		float64				`json:"max_loan" bson:"max_loan"`
	MinDown		float64				`json:"min_down" bson:"min_down"`
	LoanTerm	int64				`json:"loan_term" bson:"loan_term"`
	CreatedAt	int64				`json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdateAt	int64				`json:"update_at,omitempty" bson:"update_at,omitempty"`
}

type Message struct {
	Message		string		`json:"message,omitempty" bson:"message,omitempty"`
	Data		interface{}	`json:"data,omitempty" bson:"data,omitempty"`
}

var collection = db.Client.Database("ContactKeeper").Collection("bank")

func Create(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json")

	 var bank Bank

	 _ = json.NewDecoder(r.Body).Decode(&bank)
	bank.CreatedAt = time.Now().Unix()
	bank.UpdateAt = time.Now().Unix()

	 ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	 //cancelFunc()
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

func GetBank(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank Bank

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	ctx, cancelFunc := context.WithTimeout(context.Background(), 10 * time.Second)

	result := collection.FindOne(ctx, bson.M{"_id": id})
	cancelFunc()

	err := result.Decode(&bank)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Message{Message: "ID Does Not Exist"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bank)
}

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var banks []Bank
	Opt := options.Find().SetSort(bson.D{{"created_at", -1}})

	ctx, _:= context.WithTimeout(context.Background(), 10 * time.Second)

	result, err := collection.Find(ctx, bson.M{}, Opt)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}

	err = result.All(ctx, &banks)
	if err != nil {
		log.Printf("Error: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}

	w.WriteHeader(http.StatusOK)
	if len(banks) == 0 {
		json.NewEncoder(w).Encode(Message{Message: "Empty List", Data: nil})
		return
	}
	json.NewEncoder(w).Encode(Message{Message: "Success", Data: banks})
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var bank, result Bank

	err := json.NewDecoder(r.Body).Decode(&bank)
	if err != nil {
		log.Printf("Server Error : %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Message{Message: "Server Error"})
		return
	}

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.D{{"$set", bank}}).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Message{Message: "ID does not exist"})
		return
	}

	json.NewEncoder(w).Encode(Message{Message: "Updated Successfully", Data: result.ID})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var deletedBank Bank
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err := collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&deletedBank)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Message{Message: "ID Does Not Exist"})
		return
	}
	json.NewEncoder(w).Encode(Message{Message: "Deleted Successfully", Data: deletedBank.ID})
}