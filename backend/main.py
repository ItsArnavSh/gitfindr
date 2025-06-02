from fastapi import FastAPI
from src.api.server import StartServer
from config import settings

app = FastAPI()
StartServer(app)
