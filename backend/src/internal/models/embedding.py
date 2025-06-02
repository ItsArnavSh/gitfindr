from sqlalchemy import Column, String
from pgvector.sqlalchemy import Vector
from internal.base import Base
from config import settings
class Embedding(Base):
    __tablename__ = "embeddings"
    id = Column(String)
    vector = Column(Vector(settings.VECTOR_DIMENSIONS))
