from typing import List

from src.internal.crud.invertedindex import increment_dfi, insert_word,insert_freq
from src.internal.session import SessionLocal
def insert_keyword_list(repo_id:int,keywords:list[tuple[str, int]]):
    db = SessionLocal()
    for keyword in keywords:
        word_id = insert_word(db,keyword[0])
        increment_dfi(db,word_id)
        insert_freq(db,word_id,repo_id,keyword[1])
