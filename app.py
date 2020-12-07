from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/')
def hello():
    #Prepare the json to return
    return jsonify({'hello': 'world'})

if __name__ == '__main__':
    app.run()