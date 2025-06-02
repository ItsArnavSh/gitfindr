from models.repo import Repository
from internal.session import SessionLocal
from dataclasses import asdict
from typing import List
def create_repository(repo_data: Repository):
    db = SessionLocal()
    db.add(repo_data)
    db.commit()
    db.refresh(repo_data)
    db.close()
    return
def get_repository_by_id(repo_id: str)->Repository|None:
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
