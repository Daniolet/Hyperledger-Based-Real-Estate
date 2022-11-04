import requests
import json
resp = requests.post("http://localhost:8000/api/v1/updateSelling", data=json.dumps(
    {
        "buyer": "4e07408562be",
        "objectOfSale": "a4f093107de3fdbc",
        "seller": "4b227777d4dd",
        "status": "cancelled"
    }))

print(resp.text, resp.status_code)
