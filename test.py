import requests as rq
import math

def main():
    q = rq.get("http://localhost:8080/algo/gcd")
    id = q.headers["entry"]
    q = q.json()
    n1 = q["n1"]
    n2 = q["n2"]

    print(id)
    print(q)

    gcd = math.gcd(n1, n2)

    rsp = rq.post("http://localhost:8080/algo/gcd", json={"ans": gcd}, headers={"entry": id})

while True:
    main()