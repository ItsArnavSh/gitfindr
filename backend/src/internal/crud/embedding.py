from src.internal.models.embedding import Embedding
from src.internal.session import SessionLocal
from typing import List
from sqlalchemy import select
from sqlalchemy.orm import Session

def upsert_vector(chunk:Embedding):
    db = SessionLocal()
    db.add(chunk)
    db.commit()
    db.refresh(chunk)
    db.close()
    return

def get_top_k_chunk_ids(query_vector: List[float], k: int) -> List[int]:
    db: Session = SessionLocal()
    stmt = (
        select(Embedding.id)
        .order_by(Embedding.vector.cosine_distance(query_vector))
        .limit(k)
    )
    results = db.execute(stmt).scalars().all()
    db.close()
    return list(results)
