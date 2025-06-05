import csv
from queue import Queue
from src.service.storage.populate import store_link
from src.api.entity import QueryRequest, RegisterRequest
from src.internal.session import SessionLocal
from src.service.query.query import QueryEngine
from src.service.indexing import index
from src.service.indexing.semantic.embedding import SemanticHandler

# Initialize queue
repo_queue = Queue()

# Read CSV and enqueue "username/repo_name"
with open('repos.csv', newline='', encoding='utf-8') as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
        full_name = f"{row['username']}/{row['repo_name']}"
        repo_queue.put(full_name)

sh = SemanticHandler()
db = SessionLocal()
# Feed elements from the queue
while not repo_queue.empty():
    repo = repo_queue.get()
    if repo:
        store_link(fullname=repo)
        index.centralIndexer(repo,sh)
