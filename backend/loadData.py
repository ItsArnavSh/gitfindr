import csv
import requests
from queue import Queue
import time

# Setup
CSV_PATH = "repos.csv"
ENDPOINT = "http://0.0.0.0:8000/register"

repo_queue = Queue()

# Step 1: Load CSV and fill the queue
with open(CSV_PATH, newline='', encoding='utf-8') as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
        full_name = f"{row['username']}/{row['repo_name']}"
        repo_queue.put(full_name)

# Step 2: Send POST requests sequentially
while not repo_queue.empty():
    fullname = repo_queue.get()
    payload = {"fullname": fullname}

    try:
        print(f"Sending: {fullname}")
        response = requests.post(ENDPOINT, json=payload)
        response.raise_for_status()
        print(f"✅ Success: {fullname}")
    except requests.RequestException as e:
        print(f"❌ Error sending {fullname}: {e}")

    # Optional: small delay if needed to be polite
    # time.sleep(0.1)
