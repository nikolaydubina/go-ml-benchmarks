import os
import sys
import joblib
import xgboost as xgb
import pandas as pd
import logging
import grpc
import numpy as np

import proto.predictor_pb2_grpc
import proto.predictor_pb2

from concurrent import futures

class Predictor(proto.predictor_pb2_grpc.PredictorServicer):

    def __init__(self):
        self.preprocessor = joblib.load(os.getenv('PREPROCESSOR_PATH'))
        self.clf = xgb.XGBModel(**{'objective':'binary:logistic', 'n_estimators':10})
        self.clf.load_model(os.getenv('MODEL_PATH'))

    def Predict(self, request, context):
        features = pd.DataFrame({
            'PassengerId': pd.Series([], dtype='int64'),
            'Survived': pd.Series([], dtype='int64'),
            'Pclass': pd.Series([], dtype='int64'),
            'Name': pd.Series([], dtype='str'),
            'Sex': pd.Series([], dtype='str'),
            'Age': pd.Series([], dtype='float64'),
            'SibSp': pd.Series([], dtype='int64'),
            'Parch': pd.Series([], dtype='int64'),
            'Ticket': pd.Series([], dtype='str'),
            'Fare': pd.Series([], dtype='float64'),
            'Cabin': pd.Series([], dtype='str'),
            'Embarked': pd.Series([], dtype='str'),
        })
        features = features.append({
            'PassengerId': request.PassengerId,
            'Survived': request.Survived,
            'Pclass': request.Pclass,
            'Name': request.Name,
            'Sex': request.Sex,
            'Age': request.Age,
            'SibSp': request.SibSp,
            'Parch': request.Parch,
            'Ticket': request.Ticket,
            'Fare': request.Fare,
            'Cabin': request.Cabin,
            'Embarked': request.Embarked,
        }, ignore_index=True)

        prediction = self.clf.predict(self.preprocessor.transform(features))
        return proto.predictor_pb2.PredictResponse(Prediction=prediction[0])

    def PredictProcessed(self, request, context):
        features = np.array(request.Features).reshape((1, len(request.Features)))
        prediction = self.clf.predict(features)
        return proto.predictor_pb2.PredictResponse(Prediction=prediction[0])

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    proto.predictor_pb2_grpc.add_PredictorServicer_to_server(Predictor(), server)
    server.add_insecure_port(f'unix:///tmp/test.sock')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()