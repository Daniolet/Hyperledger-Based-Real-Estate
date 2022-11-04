import requests
import json
resp = requests.post("http://localhost:8000/api/v1/queryRealEstateList", data=json.dumps(
    {
        "accountId": "5feceb66ffc8",
        "proprietor": "4b227777d4dd",
        "totalArea": 211.8,
        "livingSpace": 111.1
    }), json=json.dumps(
    {"accountId": "5feceb66ffc8", "proprietor": "4b227777d4dd", "totalArea": 2.8, "livingSpace": 1.1}))

print(resp.text, resp.status_code)
