from sqlalchemy import Column,String,Float,Integer,PrimaryKeyConstraint
from sqlalchemy.schema import MetaData
from src.internal.base import Base
from sqlalchemy.orm import Mapped, mapped_column


class InvertedIndex(Base):
    __tablename__ = "invertedindex"
    word_id = Column(Integer, primary_key=True, index=True)
    dfi = Column(Integer)

class Dictionary(Base):
    __tablename__ = "dictionary"
    word_id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)
    word: Mapped[str] = mapped_column(index=True, unique=True, nullable=False)
class FreqTable(Base):
    __tablename__ = "freqtable"
    word_id = Column(Integer)
    repo_id = Column(Integer)
    freq = Column(Integer, nullable=False)

    __table_args__ = (
        PrimaryKeyConstraint('word_id', 'repo_id'),
    )
