from sqlalchemy import Column, String, Integer, Boolean
from sqlalchemy.dialects.postgresql import ARRAY
from src.internal.base import Base  # Make sure the import path is correct
from sqlalchemy.orm import Mapped, mapped_column

class Repository(Base):
    __tablename__ = "repositories"
    id: Mapped[int] = mapped_column(primary_key=True, autoincrement=True)
    url = Column(String, unique=True, nullable=False)
    readme_content = Column(String)
    name = Column(String, nullable=False)
    fullname = Column(String, nullable=False)
    description = Column(String)
    topics = Column(String)
    language = Column(String)
    stars = Column(Integer, default=0)
    forks = Column(Integer, default=0)
    issues = Column(Integer, default=0)
    watchers = Column(Integer, default=0)
    archived = Column(Boolean, default=False)
    forked = Column(Boolean, default=False)
