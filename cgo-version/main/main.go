package main

//go:generate go run github.com/nikolaydubina/go-featureprocessing/cmd/generate -struct=Passenger

type Passenger struct {
	Survived    int     `json:"Survived"`
	PassengerID int     `json:"PassengerId"`
	Name        string  `json:"Name"`
	PClass      float64 `json:"Pclass" feature:"identity"`
	Sex         string  `json:"Sex" feature:"onehot"`
	Age         float64 `json:"Age" feature:"minmax"`
	SibSp       float64 `json:"SibSp" feature:"quantile"`
	Parch       float64 `json:"Parch" feature:"identity"`
	Ticket      string  `json:"Ticket"`
	Fare        float64 `json:"Fare" feature:"standard"`
	Cabin       string  `json:"Cabin" feature:"ordinal"`
	Embarked    string  `json:"Embarked" feature:"onehot"`
}

func main() {}
