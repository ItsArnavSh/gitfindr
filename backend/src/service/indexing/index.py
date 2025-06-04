from src.internal.models.repo import Repository
from src.service.indexing.lexical.keywords import insert_keyword_list
from src.service.indexing.semantic.embedding import SemanticHandler
from src.service.indexing.util.util import clean_text, keyword_extraction
from collections import Counter
from config import settings
def centralIndexer(repo:Repository,sh:SemanticHandler):
    metaWeight = 3
    tagWeight = 2
    readmeText = clean_text(str(repo.readme_content))
    id = repo.id
    metadata = str(repo.metadata)
    #Semantic Indexing
    sh.loadText(id,metadata,readmeText)
    language = str(repo.language)

    keywords = keyword_extraction(readmeText)+metaWeight*keyword_extraction(metadata.lower())+3*[language]
    word_freq = Counter(keywords)
    word_list = list(word_freq.items())
    insert_keyword_list(id,word_list)
