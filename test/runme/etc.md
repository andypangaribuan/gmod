---
runme:
  id: 01HW4NBJW922AFCW0YC0R3HV5D
  version: v3
---

```sh {"id":"01HW4NBMDKFAF9VYT8ZVSY68NX"}
curl http://balancer.netdania.com/StreamingServer/StreamingServer?xml=price&
```

```python {"id":"01HW4NDEWR0SBAW3VQ0P2653J1"}
from helper.http import get

url = 'http://balancer.netdania.com/StreamingServer/StreamingServer'
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