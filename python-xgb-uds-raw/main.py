import socket
import os
import struct
import sys
import time
import xgboost as xgb
import numpy as np

SOCKET_PATH = sys.argv[1]
MODEL_PATH = sys.argv[2]
N_FEATURES = 13
BUFFER_BYTES = 512

if os.path.exists(SOCKET_PATH):
    os.remove(SOCKET_PATH)

# https://man7.org/linux/man-pages/man2/socket.2.html
# https://man7.org/linux/man-pages/man7/unix.7.html

bst = xgb.Booster({'nthread': 4})
bst.load_model(MODEL_PATH)

with socket.socket(socket.AF_UNIX, socket.SOCK_STREAM) as sock:
    sock.bind(SOCKET_PATH)
    sock.listen()

    while True:
        conn, _ = sock.accept()

        data = conn.recv(BUFFER_BYTES)
        if data == b'':
            raise RuntimeError("socket is broken, recv 0 bytes buffer")

        features = struct.unpack("f" * N_FEATURES, data)
        features = np.array(features).reshape((1, N_FEATURES))

        dmatrix = xgb.DMatrix(features)

        prediction = bst.predict(dmatrix)
        conn.send(struct.pack("f", prediction[0]))

        conn.close()