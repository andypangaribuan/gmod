---
runme:
  id: 01HW2RZW1NMW2XMM1BNN2PQ3HH
  version: v3
---

<br/> 
fetch all user

```python {"id":"01HW2XM5RKHM2WJHFPTTKDQ057","interactive":"false","mimeType":"text/x-json"}
from helper.http import get

url = 'http://localhost:3321/fetch-1'
headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'X-From-SvcName': 'x-service',
    'X-From-SvcVersion': '1.0.0',
    'X-Version': '1.0',
    'X-Source': 'android'
}

get(url, style=0, headers=headers)
```

## 

insert a user

```python {"id":"01HW39Y545VB3R2QSTE1X29X9G","interactive":"false","mimeType":"text/x-json"}
from helper.http import get, post

url = 'http://localhost:3321/insert-1'
headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'X-From-SvcName': 'x-service',
    'X-From-SvcVersion': '1.0.0',
    'X-Version': '1.0',
    'X-Source': 'ios'
}
body = {
    'name': 'andy',
    'address': 'tangerang',
    'height': 10,
    'gold_amount': 10000.12345678901234
}

post(url=url, style=0, headers=headers, json=body)
```

## 

delete user with name

```python {"id":"01HW3FFB9CTGS1CD3VJS9MM6V2","interactive":"false","mimeType":"text/x-json"}
from helper.http import get, post

url = 'http://localhost:3321/delete-1'
headers = {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'X-From-SvcName': 'x-service',
    'X-From-SvcVersion': '1.0.0',
    'X-Version': '1.0',
    'X-Source': 'gmod-test'
}
params = {
    'name': 'andy'
}

post(url=url, style=0, headers=headers, params=params)
```

## 

test multipart form data

```python {"id":"01HW3BASFFKS3GSH3YJ6KEVG71","interactive":"false","mimeType":"text/x-json"}
from helper.http import get, post

url = "http://localhost:3321/form-1"
files = [
    ('req', (None, 'one')),
    ('md-file', open('main.md', 'rb')),
    ('md-file', open('prereq.md', 'rb')),
]
headers = {
    'X-From-SvcName': 'x-service',
    'X-From-SvcVersion': '1.0.0',
    'X-Version': '1.0',
    'X-Source': 'gmod-test'
}

post(url=url, style=0, headers=headers, files=files)
```