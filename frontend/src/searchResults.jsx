import { useState } from "react";
import { motion } from "framer-motion";

export default function SearchResults() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);
  const [repoDetails, setRepoDetails] = useState([]);

  async function searchQuery() {
    try {
      console.log("Searching for:", query);
      const response = await fetch("http://localhost:3000/search", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ query }),
      });

      if (!response.ok) {
        throw new Error("Failed to fetch results");
      }

      const data = await response.json();
      console.log("Search Results:", data.results);
      setResults(data.results);
      fetchRepoDetails(data.results);
    } catch (error) {
      console.error("Error:", error);
      setResults([]);
      setRepoDetails([]);
    }
  }

  async function fetchRepoDetails(links) {
    const details = await Promise.all(
      links.map(async (link) => {
        try {
          const repoPath = link.replace("https://github.com/", "");
          const response = await fetch(
            `https://api.github.com/repos/${repoPath}`,
          );
          if (!response.ok) return null;
          return await response.json();
        } catch (error) {
          console.error("Error fetching repo details:", error);
          return null;
        }
      }),
    );
    setRepoDetails(details.filter(Boolean));
  }

  return (
    <div className="min-h-screen flex flex-col items-center p-6">
      <div className="flex space-x-2 mb-4">
        <input
          type="text"
          placeholder="Search..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          className="border rounded p-2"
        />
        <button
          onClick={searchQuery}
          className="bg-blue-500 text-white px-4 py-2 rounded"
        >
          Search
        </button>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 w-full max-w-4xl">
        {repoDetails.map((repo, index) => (
          <motion.div
            key={index}
            initial={{ opacity: 0, y: 10 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.3 }}
            className="p-4 border rounded shadow-lg bg-white"
          >
            <a
              href={repo.html_url}
              target="_blank"
              rel="noopener noreferrer"
              className="text-blue-500 text-lg font-bold"
            >
              {repo.name}
            </a>
            <p className="text-gray-700 text-sm mt-1">
              {repo.description || "No description available."}
            </p>
            <p className="text-gray-500 text-sm mt-1">
              Made by: {repo.owner?.login}
            </p>
            <p className="text-gray-500 text-sm mt-1">
              ⭐ {repo.stargazers_count} | 🍴 {repo.forks_count}
            </p>
          </motion.div>
        ))}
      </div>
    </div>
  );
}
