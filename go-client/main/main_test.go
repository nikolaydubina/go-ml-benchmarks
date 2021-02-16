package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/dmitryikh/leaves"
	"google.golang.org/grpc"

	pb "github.com/nikolaydubina/go-ml-benchmarks/go-client/proto"
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

func BenchmarkXGB_GoFeatureProcessing_GoLeaves(b *testing.B) {
	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile(os.Getenv("PREPROCESSOR_PATH"))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &fp); err != nil {
		panic(err)
	}

	model, err := leaves.XGEnsembleFromFile(os.Getenv("MODEL_PATH"), false)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		features := fp.Transform(&sample)
		model.PredictSingle(features, 0)
	}
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

func BenchmarkXGB_GoFeatureProcessing_UDS_RawBytes_Python_XGB(b *testing.B) {
	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile(os.Getenv("PREPROCESSOR_PATH"))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &fp); err != nil {
		panic(err)
	}

	addr, err := net.ResolveUnixAddr("unix", os.Getenv("SOCKET_PATH"))
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

func benchmark_UDS_gRPC_Processed_XGB(b *testing.B) {
	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile(os.Getenv("PREPROCESSOR_PATH"))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &fp); err != nil {
		panic(err)
	}

	conn, err := grpc.Dial("unix:///tmp/test.sock", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewPredictorClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		features := fp.Transform(&sample)

		response, err := client.PredictProcessed(context.Background(), &pb.PredictProcessedRequest{
			Features: features,
		})
		if response == nil || err != nil {
			panic(err)
		}
		if response.Prediction == 0 {
			panic("prediction is 0")
		}
	}
}

func benchmark_UDS_gRPC_Struct_XGB(b *testing.B) {
	conn, err := grpc.Dial("unix:///tmp/test.sock", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewPredictorClient(conn)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		response, err := client.Predict(context.Background(), &pb.PredictRequest{
			Survived:    int32(sample.Survived),
			PassengerId: int32(sample.PassengerID),
			Name:        sample.Name,
			Pclass:      sample.PClass,
			Sex:         sample.Sex,
			Age:         sample.Age,
			SibSp:       sample.SibSp,
			Parch:       sample.Parch,
			Ticket:      sample.Ticket,
			Fare:        sample.Fare,
			Cabin:       sample.Cabin,
			Embarked:    sample.Embarked,
		})
		if response == nil || err != nil {
			panic(err)
		}
		if response.Prediction == 0 {
			panic("prediction is 0")
		}
	}
}

func BenchmarkXGB_UDS_gRPC_Python_sklearn_XGB(b *testing.B) {
	benchmark_UDS_gRPC_Struct_XGB(b)
}

func BenchmarkXGB_GoFeatureProcessing_UDS_gRPC_Python_XGB(b *testing.B) {
	benchmark_UDS_gRPC_Processed_XGB(b)
}

func BenchmarkXGB_GoFeatureProcessing_UDS_gRPC_CPP_XGB(b *testing.B) {
	benchmark_UDS_gRPC_Struct_XGB(b)
}

func BenchmarkXGB_HTTP_JSON_Python_Gunicorn_Flask_sklearn_XGB(b *testing.B) {
	var fp PassengerFeatureTransformer
	config, err := ioutil.ReadFile(os.Getenv("PREPROCESSOR_PATH"))
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
