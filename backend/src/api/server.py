from src.internal.logger import logger
import csv
from queue import Queue
from typing import Union
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from src.api.entity import QueryRequest, RegisterRequest
from src.internal.session import SessionLocal
from src.service.indexing import index
from src.service.indexing.semantic.embedding import SemanticHandler
from src.service.query.query import QueryEngine
from src.service.storage.populate import store_link
def StartServer(app:FastAPI):
    sh = SemanticHandler()
    db = SessionLocal()
    qengine = QueryEngine(sh,db)
    origins = [
        "http://localhost:5173",  # React/Vite dev server
        "http://127.0.0.1:5173",
        "http://localhost",       # optional
    ]
    app.add_middleware(
        CORSMiddleware,
        allow_origins=origins,              # or ["*"] to allow all
        allow_credentials=True,
        allow_methods=["*"],                # Allow all methods (GET, POST, etc)
        allow_headers=["*"],                # Allow all headers
    )
    @app.post("/register")
    def register_repo(request: RegisterRequest):
        fullname = request.fullname
        repo = store_link(fullname=fullname)
        if repo:
            index.centralIndexer(repo,sh)
            return {"message": f"Repository {fullname} registered successfully"}
    @app.post("/query")
    def query(request: QueryRequest):
        repo_list = qengine.query(request.query)
        return repo_list
    @app.post("/load")
    def load():
        repo_queue = Queue()
        with open('repos.csv', newline='', encoding='utf-8') as csvfile:
            reader = csv.DictReader(csvfile)
            for row in reader:
                full_name = f"{row['username']}/{row['repo_name']}"
                repo_queue.put(full_name)
        while not repo_queue.empty():
            repo = repo_queue.get()
            if repo:
                store_link(fullname=repo)
                index.centralIndexer(repo,sh)
