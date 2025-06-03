from sentence_transformers import SentenceTransformer
from langchain_text_splitters import RecursiveCharacterTextSplitter
from src.internal.crud.embedding import upsert_vector
from src.internal.logger import logger
from src.internal.models.embedding import Embedding
from src.internal.models.repo import Repository
from src.internal.session import SessionLocal
class SemanticHandler:
    def __init__(self) -> None:
        logger.info("Initializing embedding pipeline...")
        self.model = SentenceTransformer('sentence-transformers/all-MiniLM-L6-v2')
    def generate_embeddings(self, text: str) -> list[float]:
        logger.info("Generating embeddings for chunk...")
        token_embeddings = self.model.encode(text)
        return token_embeddings.tolist()
    def chunker(self,text:str) -> list[str]:
        logger.info("Splitting article into chunks...")
        text_splitter = RecursiveCharacterTextSplitter(
            chunk_size=128,
            chunk_overlap=20,
            length_function=len,
            is_separator_regex=False,
        )
        docs = text_splitter.create_documents([text])
        chunks = [doc.page_content for doc in docs]
        logger.debug(f"Created {len(chunks)} chunks")
        return chunks
    def loadText(self,id:str,meta:str,text:str):
        chunks = self.chunker(text)
        if meta!=None or meta!="":
            chunks.append(meta)
        for chunk in chunks:
            upsert_vector( Embedding(id=id,vector= self.generate_embeddings(chunk)))
        logger.info("Chunks Upserted")
        return
