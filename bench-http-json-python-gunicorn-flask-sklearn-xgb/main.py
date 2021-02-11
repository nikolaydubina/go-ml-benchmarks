import os
import sys
import joblib
import xgboost as xgb
import pandas as pd

from flask import Flask
from flask import request
from flask.logging import default_handler

app = Flask(__name__)

preprocessor = joblib.load(os.getenv('PREPROCESSOR_PATH'))
clf = xgb.XGBModel(**{'objective':'binary:logistic', 'n_estimators':10})
clf.load_model(os.getenv('MODEL_PATH'))

@app.route('/predict',  methods=['POST'])
def predict():
    requestJSON = request.get_json(force=True, cache=False)

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
    features = features.append(requestJSON, ignore_index=True)

    prediction = clf.predict(preprocessor.transform(features))
    return {"prediction": str(prediction[0])}