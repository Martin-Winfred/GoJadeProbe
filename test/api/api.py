"""
This is a python program that creates a simple API using Flask.
reciving the data from the api in this case a POST request and returning a response.
"""
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/api', methods=['POST'])
def receive_data():
    # Get Json data
    data = request.get_data(as_text=True)
    # Print or save data
    print(str(data))
    # Return response
    return jsonify({'status': 'success', 'received': data}), 200

if __name__ == '__main__':
    app.run(debug=True)
