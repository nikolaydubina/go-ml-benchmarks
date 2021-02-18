package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"encoding/json"
	"testing"

	xgb "github.com/nikolaydubina/go-ml-benchmarks/cgo-version/xgb"
)

var sample = Passenger{
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

func BenchmarkXGB_CGo_GoFeatureProcessing_XGB(b *testing.B) {
	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile(os.Getenv("PREPROCESSOR_PATH"))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &fp); err != nil {
		panic(err)
	}

	model, err := xgb.XGBoosterCreate()
	if err != nil {
		panic(err)
	}
	model.LoadModel(os.Getenv("MODEL_PATH"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		features := fp.Transform(&sample)
		featuresFloat32 := make([]float32, len(features))
		for i, v := range features {
			featuresFloat32[i] = float32(v)
		}

		dmatrix, err := xgb.XGDMatrixCreateFromMat(featuresFloat32, 1, len(features), 0)
		if err != nil {
			panic(err)
		}
		prediction, err := model.Predict(dmatrix)
		if err != nil {
			panic(err)
		}
		if len(prediction) != 1 {
			panic(fmt.Errorf("wrong length of prediction: %#v", prediction))
		}
		if prediction[0] == 0 {
			panic("prediction is 0")
		}
	}
}

