from sqlalchemy import Column,String,Float,Integer,PrimaryKeyConstraint
from sqlalchemy.schema import MetaData
from src.internal.base import Base
from sqlalchemy.orm import Mapped, mapped_column


class InvertedIndex(Base):
    __tablename__ = "invertedindex"
    word_id: Mapped[int] = mapped_column(primary_key=True)
    dfi: Mapped[int] = mapped_column()
class Dictionary(Base):
    __tablename__ = "dictionary"
    word_id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)
    word: Mapped[str] = mapped_column(index=True, unique=True, nullable=False)
class FreqTable(Base):
    __tablename__ = "freqtable"

    word_id: Mapped[int] = mapped_column()
    repo_id: Mapped[int] = mapped_column()
    freq: Mapped[int] = mapped_column(nullable=False)

    __table_args__ = (
        PrimaryKeyConstraint('word_id', 'repo_id'),
    )
