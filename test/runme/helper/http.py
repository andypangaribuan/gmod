import requests
import json
from rich import print_json, print


def get(url: str, style: int = 0, headers: dict[str, str] = None):
    r = requests.get(url, headers=headers)
    __show(r, style)


def post(url: str, style: int = 0, headers: dict[str, str] | None = None, json: any = None, files: any = None, params: any = None):
    r = requests.post(url, headers=headers, json=json, files=files, params=params)
    __show(r, style)


def __show(r: requests.Response, style: int):
    match style:
        case 0:
            print(f'status: {r.status_code}')
            print_json(json.dumps(dict(r.headers)))

            print('\n\n')
            try:
                print_json(r.text)
            except:
                print(r.text)

        case 1:
            try:
                print_json(r.text)
            except:
                print(r.text)
