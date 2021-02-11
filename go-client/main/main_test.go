package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"testing"

	"github.com/dmitryikh/leaves"
)

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

// predictRawBytes uses raw floats encoding
func predictRawBytes(w io.Writer, r io.Reader, features []float64) (float64, error) {
	if err := binary.Write(w, binary.LittleEndian, features); err != nil {
		return 0, err
	}

	var prediction float64
	if err := binary.Read(r, binary.LittleEndian, &prediction); err != nil {
		return 0, err
	}

	return prediction, nil
}

func benchmarkUDSRawBytesNewConn(b *testing.B, socketpathin string) {
	// PassengerId,Pclass,Name,Sex,Age,SibSp,Parch,Ticket,Fare,Cabin,Embarked
	// 904,1,"Snyder, Mrs. John Pillsbury (Nelle Stevenson)",female,23,1,0,21228,82.2667,B45,S
	sample := Passenger{
		PassengerID: 904,
		PClass:      1,
		Name:        "Snyder, Mrs. John Pillsbury (Nelle Stevenson)",
		Sex:         "female",
		Age:         23,
		SibSp:       1,
		Parch:       0,
		Ticket:      "A/B 21228",
		Fare:        82.2667,
		Cabin:       "B45",
		Embarked:    "S",
	}

	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile("../../data/models/go-featureprocessor.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &fp); err != nil {
		panic(err)
	}

	addr, err := net.ResolveUnixAddr("unix", socketpathin)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c, err := net.DialUnix("unix", nil, addr)
		if err != nil {
			panic(err)
		}

		features := fp.Transform(&sample)

		result, err := predictRawBytes(c, c, features)
		if result == 0 || err != nil {
			panic(err)
		}
		c.Close()
	}
}

func BenchmarkXGB_Python_UDS_RawBytes_NewConnection(b *testing.B) {
	benchmarkUDSRawBytesNewConn(b, path.Join(os.Getenv("PROJECT_PATH"), "sc"))
}

func BenchmarkXGB_GoFeatureProcessing_GoLeaves(b *testing.B) {
	// PassengerId,Pclass,Name,Sex,Age,SibSp,Parch,Ticket,Fare,Cabin,Embarked
	// 904,1,"Snyder, Mrs. John Pillsbury (Nelle Stevenson)",female,23,1,0,21228,82.2667,B45,S
	sample := Passenger{
		PassengerID: 904,
		PClass:      1,
		Name:        "Snyder, Mrs. John Pillsbury (Nelle Stevenson)",
		Sex:         "female",
		Age:         23,
		SibSp:       1,
		Parch:       0,
		Ticket:      "A/B 21228",
		Fare:        82.2667,
		Cabin:       "B45",
		Embarked:    "S",
	}

	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile(path.Join(os.Getenv("PROJECT_PATH"), "data", "models", "go-featureprocessor.json"))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &fp); err != nil {
		panic(err)
	}

	model, err := leaves.XGEnsembleFromFile(path.Join(os.Getenv("PROJECT_PATH"), "data", "models", "titanic_v090.xgb"), false)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		features := fp.Transform(&sample)
		model.PredictSingle(features, 0)
	}
}

func benchmarkRestJSON(b *testing.B) {
	// PassengerId,Pclass,Name,Sex,Age,SibSp,Parch,Ticket,Fare,Cabin,Embarked
	// 904,1,"Snyder, Mrs. John Pillsbury (Nelle Stevenson)",female,23,1,0,21228,82.2667,B45,S
	sample := Passenger{
		PassengerID: 904,
		PClass:      1,
		Name:        "Snyder, Mrs. John Pillsbury (Nelle Stevenson)",
		Sex:         "female",
		Age:         23,
		SibSp:       1,
		Parch:       0,
		Ticket:      "A/B 21228",
		Fare:        82.2667,
		Cabin:       "B45",
		Embarked:    "S",
	}

	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile("../../data/models/go-featureprocessor.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &fp); err != nil {
		panic(err)
	}

	type Response struct {
		Prediction float64 `json:"prediction,string"`
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		features, err := json.Marshal(sample)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post("http://0.0.0.0:80/predict", "application/json", bytes.NewReader(features))
		if resp == nil || resp.Body == nil || err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			panic("non 200 response")
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		response := Response{}
		if err := json.Unmarshal(bodyBytes, &response); err != nil {
			panic(fmt.Errorf("body %#v : %w", string(bodyBytes), err))
		}
		if response.Prediction == 0 {
			panic("prediction is 0")
		}
	}
}

func BenchmarkXGB_Python_JSON_Gunicorn_Flask_sklearn_xgb(b *testing.B) {
	benchmarkRestJSON(b)
}
