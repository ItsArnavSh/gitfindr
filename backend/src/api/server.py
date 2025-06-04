from src.internal.logger import logger
from typing import Union
from fastapi import FastAPI
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
    @app.post("/register")
    def register_repo(request: RegisterRequest):
        fullname = request.fullname
        repo = store_link(fullname=fullname)
        index.centralIndexer(repo,sh)
        return {"message": f"Repository {fullname} registered successfully"}
    @app.get("/query")
    def query(request: QueryRequest):
        repo_list = qengine.query(QueryRequest.query)
        for repo in repo_list:
            print(repo.fullname)
