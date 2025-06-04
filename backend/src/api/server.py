from src.internal.logger import logger
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
        index.centralIndexer(repo,sh)
        return {"message": f"Repository {fullname} registered successfully"}
    @app.post("/query")
    def query(request: QueryRequest):
        repo_list = qengine.query(request.query)
        return repo_list
