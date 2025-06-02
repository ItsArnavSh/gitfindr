from src.internal.logger import logger
from typing import Union
from fastapi import FastAPI
from src.entity.entity import RegisterRequest
from src.storage.populate import store_link
def StartServer(app:FastAPI):
    @app.post("/register")
    def register_repo(request: RegisterRequest):
        fullname = request.fullname
        store_link(fullname=fullname)
        return {"message": f"Repository {fullname} registered successfully"}
