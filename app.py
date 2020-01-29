from flask import Flask, render_template, request
from pymongo import MongoClient
import json

app = Flask(__name__, static_url_path='/static')
client = MongoClient('localhost', 27017)
db = client.db


@app.route("/")
def greet():
    return render_template("index.html")

@app.route("/person")
def get_personal_info():
    teachers = db.teachers
    uid = request.args.get("uid")
    result = teachers.find_one({"uid": int(uid)})
    return json.dumps(result, default=lambda o: "")


if __name__ == "__main__":
    app.run()
