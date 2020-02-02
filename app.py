from flask import Flask, render_template, request, send_from_directory, send_file
from pymongo import MongoClient
import json

app = Flask(__name__, static_url_path='/static')
client = MongoClient('localhost', 27017)
db = client.db


def trunc_filename(name):
    j = 0
    for i in range(len(name), 0):
        if name[i] == '/' or name[i] == '\\':
            j = i
    return name[j:]


@app.route("/")
def greet():
    return send_file("/Users/jernicozz/WebstormProjects/wikireact/public/index.html")


@app.route("/person")
def get_personal_info():
    teachers = db.teachers
    uid = request.args.get("uid")
    result = teachers.find_one({"uid": int(uid)}, {"text": 0})
    return json.dumps(result, default=lambda o: "")


@app.route("/data/<filename>")
def send_data(filename):
    return send_from_directory("data", filename)


@app.route("/allPersons")
def send_all_persons():
    return json.dumps([{"persName": "Кумсков", "persStatus": "Профессор", "persAvatar": "http://localhost:5000/data/vD6kgDeBImc.jpg", "uid": 1}])


@app.route("/newReview", methods=["POST"])
def register_review():
    file = request.files["ava"]
    name = trunc_filename(file.filename)
    file.save(f"data/{name}")
    text = request.form["text"]
    pName = request.form["name"]
    status = request.form["status"]
    jsobj = {"name": pName, "text": text, "status": status, "avatar": name}
    db.pendingReviews.insert(jsobj)
    return "200"


if __name__ == "__main__":
    app.run()
