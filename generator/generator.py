import requests
import random
import time


def send_req(computation: float):
    host = "localhost:8888"

    print("Sending request.")
    req = requests.post(f"http://{host}/task/order",
                    data={
                      "computation": f"{computation}"
                    },
                    headers={'Content-Type': 'application/x-www-form-urlencoded'})
    print(f"Res: {req.status_code}, {req.reason}")
    

if __name__ == "__main__":
    print("### Starting generator ###")
    try:
        while True:
            time.sleep(random.expovariate(1/4))
            send_req(random.uniform(0, 5))
    except Exception as e:
        print(e)
