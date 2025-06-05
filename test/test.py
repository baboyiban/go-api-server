import requests

BASE_URL = 'http://localhost:8080/api'

def test_api():
    response = requests.get(f'{BASE_URL}/zones')
    print(f'response:{response}')

    if response.status_code == 200:
        # 성공
        print(response.json())
    else:
        # 실패
        print(response.text())

if __name__ == "__main__":
    test_api()
