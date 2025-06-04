
from src.internal.crud.embedding import get_top_k_chunk_ids
from src.internal.crud.invertedindex import insert_word, word_dfi, word_freqlist
from src.internal.crud.repo import avg_doc_size, doc_size, get_repository_by_id, repository_count
from src.internal.models.repo import Repository
from src.service.indexing.semantic.embedding import SemanticHandler
from src.service.indexing.util.util import clean_text, keyword_extraction
from typing import List
from collections import defaultdict
import math
class QueryEngine:
    def __init__(self,semanticHandler:SemanticHandler,dbHandler):
        self.sh = semanticHandler
        self.db = dbHandler
    def query(self,query:str)->List[Repository]:
        query = clean_text(query)#For Semantic
        keywords = keyword_extraction(query)#For BM25

        limit = 20
        embedding = self.sh.generate_embeddings(query)
        semanticRankings = get_top_k_chunk_ids(embedding,limit)
        keywordRankings = self.bm25(keywords)
        final_rankings =  self.rrf_merge([semanticRankings,keywordRankings])
        repos:List[Repository] = []
        for repo_id in final_rankings:
            repo = get_repository_by_id(repo_id)
            if repo!=None:
                repos.append(repo)
        return repos
    def bm25(self,keywords:List[str])->List[int]:
        B = 0.75
        K = 1.5
        BM25 = defaultdict(float)
        avgdl = avg_doc_size() or 1
        for keyword in keywords:
            word_id = insert_word(self.db,keyword)
            dfi = word_dfi(self.db,word_id)#How many documents the word is in
            freq_list = word_freqlist(self.db,word_id)# Which all repos the word is in and its frequency there
            no_repo = repository_count()
            #idf calc
            idf = math.log((no_repo-dfi+0.5)/(dfi+0.5)+1)
            for freq_data in freq_list:
                freq = freq_data.freq
                repoid = freq_data.repo_id
                #So in repoid x the word_id has occured freq times
                size = doc_size(repoid)
                curr_bm = freq*(K+1)/(freq+K*(1-B+B*size/avgdl))*idf
                #so bm score has to be added
                BM25[repoid] += curr_bm
        # Now sort and return the repo names
        sorted_keys = sorted(BM25, key=lambda k: BM25[k], reverse=True)
        return sorted_keys

    def rrf_merge(self,rankings: list[list[int]], k: int = 60, top_n: int = 100) -> list[int]:
        scores = defaultdict(float)

        for ranking in rankings:
            for rank, doc_id in enumerate(ranking, start=1):  # rank starts at 1
                scores[doc_id] += 1 / (k + rank)
        # Sort by score in descending order
        sorted_docs = sorted(scores.items(), key=lambda x: x[1], reverse=True)

        return [doc_id for doc_id, _ in sorted_docs[:top_n]]
