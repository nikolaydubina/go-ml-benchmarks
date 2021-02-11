import socket
import os
import struct
import sys
import time
import xgboost as xgb
import numpy as np

N_FEATURES = 12
BUFFER_BYTES = 1024

# https://man7.org/linux/man-pages/man2/socket.2.html
# https://man7.org/linux/man-pages/man7/unix.7.html

clf = xgb.XGBModel(**{'objective':'binary:logistic', 'n_estimators':10})
clf.load_model(os.getenv("MODEL_PATH"))

with socket.socket(socket.AF_UNIX, socket.SOCK_STREAM) as sock:
    sock.bind(os.getenv("SOCKET_PATH"))
    sock.listen()

    while True:
        conn, _ = sock.accept()

        data = conn.recv(BUFFER_BYTES)
        if data == b'':
            raise RuntimeError("socket is broken, recv 0 bytes buffer")

        # f = float32
        # d = float64
        features = struct.unpack("d" * N_FEATURES, data)
        features = np.array(features).reshape((1, N_FEATURES))

        prediction = clf.predict(features)
        
        conn.send(struct.pack("d", prediction[0]))

        conn.close()