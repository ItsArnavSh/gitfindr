# 🚀 GitFindr

### 🔎 The Google for GitHub

## 📌 Introduction

GitFindr is an advanced search tool that enhances repository discovery by utilizing an improved **BM25-based ranking algorithm**. It integrates repository statistics like **⭐ stars, 🍴 forks, and 👀 clicks** to refine search results and provide more relevant rankings.

## 🤔 Why GitFindr?

GitHub's built-in search often fails to provide the most relevant repositories due to its reliance on basic keyword matching. Many great projects remain undiscovered because:

- **README files and descriptions aren’t fully analyzed**, causing keyword-dependent searches.
- **Synonyms aren’t considered**, meaning searches miss related terms.
- **Ranking is inconsistent**, prioritizing older or more forked repositories regardless of recent relevance.

GitFindr fixes these issues by introducing **README scanning** and **synonym matching**, ensuring that even loosely related terms surface the right repositories. By leveraging an optimized BM25 ranking algorithm, GitFindr provides more precise and meaningful search results.

## ⚙️ How It Works

GitFindr is built mainly in **Golang** for speed and reliability.

### 🔍 Keyword Extraction

1. **Processing Repository Links**: When a repository link is provided, the **Python backend (FastAPI)** processes it.
2. **Extracting Meaningful Keywords**: Using **spaCy** 🧠 and **regex** ✍️, the backend extracts a list of relevant keywords from the repository description, README, and codebase.

### 📖 Indexing & Search Optimization

1. **Creating Inverted Indexes**: The **Go backend** processes the extracted keywords and maintains two separate **inverted indexes** ([Wikipedia](https://en.wikipedia.org/wiki/Inverted_index)):
   - **🔄 Synonym-based Indexing**: Allows for **vague** term matching.
   - **✅ Exact Term Indexing**: Ensures precision in results.
2. **Handling Synonyms**:
   - A **synonym API** fetches related words (e.g., "hi" and "hello" share an index).
   - Each word maps to an **index**, but since words can have **multiple meanings**, they may link to **multiple indexes**.
   - To **avoid repeated calls**, once synonyms are fetched, the indexes are mapped to **Redis** as a cache.
   - Initially, this caused **overly vague** search results.
3. **Introducing Exact Term Matching**:
   - A separate **exact term table** ignores synonyms, improving accuracy.
   - **BM25 ranking** is computed for both tables, with **higher weight** assigned to exact matches.
   - This hybrid approach ensures **optimal search precision** across various queries.

## 📊 BM25 Calculation

GitFindr's ranking mechanism enhances **BM25** by incorporating weights and repository interaction metrics. Below is the **BM25S formula** used:

\({ BM25S(D, Q) = \sum_{i=1}^{|Q|} IDF(q_i) \cdot \frac{f(q_i, D) \cdot (K+1)}{f(q_i, D) + K \cdot (1 - b + b \cdot \frac{|D|}{avgD})} \cdot alt }\)

Where:

- **📖 Inverse Document Frequency (IDF):**

  \({ IDF(q_i) = \log \left( \frac{N - df_i + 0.5}{df_i + 0.5} + 1 \right) }\)

  - `N` = Total number of documents (repositories)
  - `df_i` = Number of documents containing term `i`

- **📈 Term Frequency Weighting:**

  \({ f(q_i, D) = \sum_{b} v_b \cdot qd_i^b }\)

  - `v_b` = Frequency weight of field `b`
  - `qd_i^b` = Total occurrences of `q_i` in field `b` of document `D`

- **⚖️ Scaling Factor (****`K`****):**

  \({ K = k_1 \cdot \frac{\text{avg term freq in dataset}}{\text{avg term freq in dataset after weighting}} }\)

  - `k_1` is a tunable parameter (`k_1 ∈ [1.2, 2.0]`)

- **📊 Additional Weighting (****`alt`****):**

  \({ alt = (1 + \sum_{i} \alpha_i \log (1 + x_i)) }\)

  - `x_i` represents repository statistics (**⭐ stars, 🍴 forks, 👀 clicks**)
  - `α_i` is a tuning constant

## 🛠️ Installation & Usage

### 📂 Project Structure

GitFindr consists of two main folders:

- **Frontend**: The UI for searching repositories.
- **Backend**: Handles indexing, searching, and processing.

### 🚀 Frontend Setup

1. Navigate to the frontend directory:
   ```sh
   cd frontend
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Start the frontend:
   ```sh
   npm run dev
   ```

### 🔧 Backend Setup

The backend consists of three folders:

#### 1️⃣ `pyProcess` (Python API)

- **Dockerized**, just run the backend container.

#### 2️⃣ `indexer` (Go Indexer)

- Navigate to the folder:
  ```sh
  cd backend/indexer
  ```
- Install dependencies and run:
  ```sh
  go get
  go run main.go
  ```

#### 3️⃣ `redisCloner` (Redis Cache Loader)

- Contains Python scripts to **preload Redis** with synonyms for faster setup.

### ⚡ Redis Setup

- Ensure **Redis** is installed.
- Install **Redis CLI** ([Download here](https://redis.io/docs/getting-started/)).

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**.
2. **Create a feature branch** (`git checkout -b feature-branch`).
3. **Commit changes** (`git commit -m 'Add new feature'`).
4. **Push to your branch** (`git push origin feature-branch`).
5. **Open a pull request**.

## 🎯 Vision

GitFindr envisions **a community where people can discover and submit their ideas**, ensuring that no idea gets buried and every project gets a fair chance to be seen. 🚀

