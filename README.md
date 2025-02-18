# 🚀 GitFindr

### 🔍 The Problem with Current GitHub Search
GitHub's search functionality has several limitations that make it difficult for developers to discover relevant repositories:

1. **Over-Reliance on Metadata**  
   - Search results depend heavily on repository names, languages, and descriptions.
   - Repositories with vague or incomplete descriptions are often overlooked.

2. **Lack of Support for Synonyms & Acronyms**  
   - Searching for *"JS tools"* vs. *"JavaScript tools"* can yield different results.
   - This inconsistency makes searching frustrating and inefficient.

---

## ✅ Our Solution
We are building an advanced search engine that enhances GitHub's search capabilities by:

### 🏗 **Workflow Overview**
Our system consists of two main components:

1. **🌐 FastAPI Python Server**  
   - Accepts repository links from users.  
   - Downloads and extracts README.md content.  
   - Processes README into a **list of words** for further analysis.

2. **🚀 Go Backend Engine**  
   - Converts the processed words into a **list of indexed terms** for synonym matching.  
   - Utilizes **Redis for caching** synonym lookups.  
   - Incorporates GitHub statistics like **languages, stars, and forks** to enhance ranking.  
   - Implements a **modified BM25 algorithm** for weighted frequency, considering additional repository metadata beyond just keywords.  
   - Stores an **inverted index in SQLite** for efficient searching.

---

## 📦 Building a Repository Database
We collect repositories using three methods:

1. **🌐 Web Crawling**  
   - Crawl developer websites for repository links.  
   - Recursively follow discovered GitHub links to index repositories.

2. **📝 Manual Submission**  
   - Allow users to submit repositories directly to our search database.

3. **🔗 GitHub Search API**  
   - Use the GitHub API to identify and add repositories.
   - The API is rate-limited, but it provides a valuable starting point.

---

## ⚙️ Preprocessing Repositories
After collecting repositories, we process the data using the FastAPI Python server:

1. **📖 Normalization & Tokenization**  
   - Extract text from README files.
   - Convert content into a normalized list of words by stemming and removing stopwords.

2. **🚫 Duplicate Detection**  
   - Compute a **SHA-256 hash** of repository content to detect and prevent duplicate indexing.

3. **✅ Typo Correction**  
   - Perform spell-checking and autocorrection for better search accuracy.

---

## 🔄 Handling Synonyms & Acronyms with Redis
To enhance search flexibility, we use Redis for synonym clustering:

1. **🗂️ Synonym Clustering**  
   - Map words like *"js"* and *"javascript"* to the same ID.
   - If a word is missing, fetch synonyms via an external API.

2. **⚡ Efficient Caching**  
   - Cache API responses in Redis for quick lookups.
   - Store queries as **SHA-256 hashed keys** to avoid redundant API calls.

---

## 📊 Building the Inverted Index
To enable fast and efficient searches, we construct an **inverted index**:

- Stored in **SQLite** with three columns:
  - **Word Index** → A unique identifier (e.g., Redis ID)
  - **Frequency** → Number of times the word appears in a repository
  - **Repositories** → A list of repositories containing the word

---

## 📈 Search Ranking & Results
Search results are ranked using a **modified BM25 algorithm** that considers:

1. **🔎 Weighted Frequency**  
   - Matches user queries with indexed words, adjusted for their significance.

2. **⭐ Repository Popularity**  
   - Factors in GitHub stats like **stars, forks, and contributors** as ranking signals.

3. **📅 Recent Activity & Additional Metrics**  
   - Prioritizes active repositories and relevant language usage.

---

## 🎯 Why This Matters
By solving GitHub's search limitations, we provide developers with:
✅ **More accurate results** – Even when descriptions are missing.  
✅ **Support for synonyms & typos** – Making search more flexible.  
✅ **Efficient ranking** – Prioritizing quality repositories.

---

## 📌 Future Development
- 🔄 Implementing **real-time updates** for repository changes.
- 🔍 Adding **semantic search** for deeper understanding of queries.
- 🌐 Providing **a web interface** for easy access.

🚀 **Stay tuned for updates!**
