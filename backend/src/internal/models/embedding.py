from sqlalchemy import Column, String,Integer
from pgvector.sqlalchemy import Vector
from src.internal.base import Base
from config import settings
class Embedding(Base):
    __tablename__ = "embeddings"
    prim_id = Column(Integer, primary_key=True, autoincrement=True)
    id = Column(Integer)
    vector = Column(Vector(settings.VECTOR_DIMENSIONS))
