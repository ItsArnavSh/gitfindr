# 🚀 GitFindr

### 🔎 The Google for GitHub

## 📌 Introduction

GitFindr is an advanced search tool that enhances repository discovery by combining **BM25-based ranking** with **semantic search** using transformer-based embeddings. It refines search results using repository statistics like **⭐ stars, 🍴 forks, and 👀 clicks**, providing highly relevant rankings.

## 🤔 Why GitFindr?

GitHub's built-in search often fails to provide the most relevant repositories due to its reliance on basic keyword matching. Many great projects remain undiscovered because:

* README files and descriptions aren’t fully analyzed, causing keyword-dependent searches.
* Synonyms aren’t considered, meaning searches miss related terms.
* Ranking is inconsistent, prioritizing older or more forked repositories regardless of recent relevance.

GitFindr fixes these issues by introducing README scanning and semantic matching, ensuring that even loosely related terms surface the right repositories. By leveraging a hybrid **BM25 + embedding search**, GitFindr delivers more precise and meaningful results.

## ⚙️ How It Works

GitFindr is now built fully in **Python**, utilizing **FastAPI**, **PostgreSQL with pgvector**, and **sentence-transformers** for semantic indexing.

### 🔍 Indexing & Search Optimization

1. **Embedding-Based Semantic Search**:

   * We use `sentence-transformers/all-MiniLM-L6-v2` to generate vector embeddings for repository README files and metadata.
   * These embeddings are stored in PostgreSQL using the `pgvector` extension.
   * During a query, the user's input is also converted into an embedding vector.
   * We compute **cosine similarity** between the query embedding and repository embeddings to retrieve semantically similar results.
   * This allows us to match relevant repositories even when there are no exact keyword overlaps.

2. **BM25-Based Search**:

   * Traditional keyword-based indexing is performed using PostgreSQL's full-text search engine.
   * We calculate BM25 scores for each document using weighted fields such as repository name, description, and README content.

3. **Hybrid Ranking with RRF (Reciprocal Rank Fusion)**:

   * Both BM25 and semantic search generate separate ranked lists of repository results.

   * We apply **Reciprocal Rank Fusion (RRF)** to combine the two rankings:

     $$
     RRF(r) = \sum_{i=1}^{n} \frac{1}{k + rank_i(r)}
     $$

     * `rank_i(r)` is the rank of repository `r` in the i-th ranked list
     * `k` is a tunable constant (e.g., 60)

   * This method rewards results that appear in both rankings, regardless of exact rank, balancing precision and recall effectively.

This combined approach ensures both exact keyword matches and semantically similar results are surfaced accurately.

## 📊 BM25 Calculation

GitFindr still enhances BM25 ranking by incorporating interaction metrics. Below is the **BM25S formula** used:

$$
BM25S(D, Q) = \sum_{i=1}^{|Q|} IDF(q_i) \cdot \frac{f(q_i, D) \cdot (K+1)}{f(q_i, D) + K \cdot (1 - b + b \cdot \frac{|D|}{avgD})} \cdot alt
$$

Where:

* **📖 Inverse Document Frequency (IDF):**

  $$
  IDF(q_i) = \log \left( \frac{N - df_i + 0.5}{df_i + 0.5} + 1 \right)
  $$

  * `N` = Total number of documents (repositories)
  * `df_i` = Number of documents containing term `i`

* **📈 Term Frequency Weighting:**

  $$
  f(q_i, D) = \sum_{b} v_b \cdot qd_i^b
  $$

  * `v_b` = Frequency weight of field `b`
  * `qd_i^b` = Total occurrences of `q_i` in field `b` of document `D`

* **⚖️ Scaling Factor (`K`):**

  $$
  K = k_1 \cdot \frac{\text{avg term freq in dataset}}{\text{avg term freq in dataset after weighting}}
  $$

  * `k_1` is a tunable parameter (`k_1 ∈ [1.2, 2.0]`)

* **📊 Additional Weighting (`alt`):**

  $$
  alt = (1 + \sum_{i} \alpha_i \log (1 + x_i))
  $$

  * `x_i` represents repository statistics (**⭐ stars, 🍴 forks, 👀 clicks**)
  * `α_i` is a tuning constant

## 🛠️ Installation & Usage

### 📂 Project Structure

GitFindr consists of two main folders:

* **Frontend**: The UI for searching repositories.
* **Backend**: Handles indexing, searching, and processing.

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

1. **Start PostgreSQL with pgvector support using Docker**

   Make sure Docker is installed. From the `backend/` directory, run:

   ```bash
   docker compose up -d
   ```

   This will start a PostgreSQL instance with `pgvector` extension enabled.

2. **Install `uv` for dependency management**

   Follow the installation guide here: [https://docs.astral.sh/uv/](https://docs.astral.sh/uv/)

3. **Install dependencies and start the backend**

   Run the following commands from the `backend/` folder:

   ```bash
   make
   make run
   ```

   This will install Python dependencies and launch the FastAPI server.

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**.
2. **Create a feature branch** (`git checkout -b feature-branch`).
3. **Commit changes** (`git commit -m 'Add new feature'`).
4. **Push to your branch** (`git push origin feature-branch`).
5. **Open a pull request**.

## 🎯 Vision

GitFindr envisions a community where people can discover and submit their ideas, ensuring that no idea gets buried and every project gets a fair chance to be seen. 🚀
