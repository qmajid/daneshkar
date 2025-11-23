from flask import Flask, jsonify, request
import jwt, datetime, os

app = Flask(__name__)
SECRET = os.getenv("JWT_SECRET", "supersecret")

@app.route("/token", methods=["POST"])
def generate_token():
    data = request.json or {}
    user = data.get("username", "anonymous")
    password = data.get("password", "")

    payload = {
        "user": user,
        "exp": datetime.datetime.utcnow() + datetime.timedelta(minutes=30)
    }

    token = jwt.encode(payload, SECRET, algorithm="HS256")
    return jsonify({"token": token})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)