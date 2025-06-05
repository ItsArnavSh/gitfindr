from src.internal.models.repo import Repository
from src.internal.session import SessionLocal
from dataclasses import asdict
from typing import List
from sqlalchemy.sql import func
from sqlalchemy.exc import SQLAlchemyError
def create_repository(repo_data: Repository) -> Repository | None:
    db = SessionLocal()
    try:
        db.add(repo_data)
        db.commit()
        db.refresh(repo_data)
        return repo_data
    except SQLAlchemyError as e:
        db.rollback()
        print(f"Database error: {e}")  # Replace with proper logging in production
        return None
    finally:
        db.close()
def get_repository_by_id(repo_id: int)->Repository|None:
    db = SessionLocal()
    repo = db.query(Repository).filter(Repository.id == repo_id).first()
    db.close()
    if repo != None:
        return repo
    return None
def list_repos()->List[Repository]:
    db = SessionLocal()
    try:
        repos = db.query(Repository).all()
        return repos
    finally:
        db.close()
def repository_count() -> int:
    db = SessionLocal()
    try:
        count = db.query(Repository).count()
        return count
    finally:
        db.close()
def avg_doc_size() -> float:
    db = SessionLocal()
    try:
        average = db.query(func.avg(Repository.size)).scalar()
        return float(average or 0.0)
    finally:
        db.close()
def doc_size(doc_id:int)->int:
    db = SessionLocal()
    try:
        doc = db.query(Repository).filter(Repository.id==doc_id).first()
        return doc.size  if doc is not None else 0
    finally:
        db.close()
