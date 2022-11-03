import requests
import json
resp = requests.post("http://localhost:8000/api/v1/queryAccountList", data=json.dumps(
    {
        "args": [{
            "accountId": "5feceb66ffc8"
        }]

    }), headers={
    "Cookie": "sidebarStatus=1"
})

print(resp.text, resp.status_code)
