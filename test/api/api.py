from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/api', methods=['POST'])
def receive_data():
    # Get Json data
    data = request.get_json()
    # Print or save data
    print(data)
    # Return response
    return jsonify({'status': 'success', 'received': data}), 200

if __name__ == '__main__':
    app.run(debug=True)
