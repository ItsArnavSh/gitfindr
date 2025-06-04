from sqlalchemy.orm import Session
from src.internal.models.invertedindex import Dictionary, InvertedIndex, FreqTable
from sqlalchemy.dialects.postgresql import insert
from typing import List
#Create Operations
def insert_word(session: Session, word: str) -> int:
    existing = session.query(Dictionary).filter_by(word=word).first()
    if existing:
        return existing.word_id
    entry = Dictionary(word=word)
    session.add(entry)
    session.commit()
    session.refresh(entry)
    return entry.word_id


def increment_dfi(session, word_id: int):
    stmt = insert(InvertedIndex).values(word_id=word_id, dfi=1)
    stmt = stmt.on_conflict_do_update(
        index_elements=['word_id'],
        set_={'dfi': InvertedIndex.dfi + 1}
    )
    session.execute(stmt)
    session.commit()

def insert_freq(session: Session, word_id: int, repo_id: int, freq: int) -> None:
    entry = FreqTable(word_id=word_id, repo_id=repo_id, freq=freq)
    session.add(entry)
    session.commit()

#Read Operations
def word_dfi(session:Session,word_id:int)->int:
    word_data = session.query(InvertedIndex).filter(InvertedIndex.word_id == word_id).first()
    if word_data != None:
        return word_data.dfi
    return 0
def word_freqlist(session:Session,word_id)->List[FreqTable]:
    word_data = session.query(FreqTable).filter(FreqTable.word_id == word_id).all()
    return word_data
