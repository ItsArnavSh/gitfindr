from src.internal.models.embedding import Embedding
from src.internal.session import SessionLocal
from typing import List

def upsert_vector(chunk:Embedding):
    db = SessionLocal()
    db.add(chunk)
    db.commit()
    db.refresh(chunk)
    db.close()
    return
