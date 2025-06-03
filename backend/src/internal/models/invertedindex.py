from sqlalchemy import Column,String,Float,Integer,JSON
from sqlalchemy.schema import MetaData
from src.internal.base import Base
class InvertedIndex(Base):
    __tablename__ = "invertedindex"
    word = Column(String, primary_key=True, index=True)
    dfi = Column(Integer)
    idf = Column(Float)
    data = Column(JSON)
