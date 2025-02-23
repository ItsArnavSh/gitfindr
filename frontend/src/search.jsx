import { BrowserRouter as Router, Route, Routes, Link } from "react-router-dom";
import { motion } from "framer-motion";
import { Search } from "lucide-react";
import RegisterModal from "./registerModal";
import { useState, useEffect } from "react";
import { ArrowUp } from "lucide-react";
export default function Home() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);
  const [repoDetails, setRepoDetails] = useState([]);
  const [showButton, setShowButton] = useState(false);
  const [showModal, setShowModal] = useState(false);
  useEffect(() => {
    const handleScroll = () => {
      setShowButton(window.scrollY > 300);
    };
    window.addEventListener("scroll", handleScroll);
    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

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
    const apiKey = import.meta.env.GITHUB_API_KEY; // Use correct env handling
    console.log(apiKey);
    console.log(apiKey);
    const details = await Promise.all(
      links.map(async (link) => {
        try {
          const repoPath = link.replace("https://github.com/", "");
          const response = await fetch(
            `https://api.github.com/repos/${repoPath}`,
            {
              headers: {
                Authorization: `Bearer ${apiKey}`, // Use Bearer instead of token
                Accept: "application/vnd.github.v3+json",
              },
            },
          );

          if (!response.ok)
            throw new Error(`GitHub API error: ${response.status}`);
          return await response.json();
        } catch (error) {
          console.error("Error fetching repo details:", error);
          return null;
        }
      }),
    );

    setRepoDetails(details.filter(Boolean));
    setTimeout(() => {
      document.getElementById("repos")?.scrollIntoView({ behavior: "smooth" });
    }, 200);
  }

  return (
    <Router>
      <div className="min-h-screen bg-white">
        {/* Navigation */}
        <nav className="flex items-center justify-between p-4 md:p-6">
          <Link to="/" className="flex items-center space-x-2">
            <img
              src="https://hebbkx1anhila5yf.public.blob.vercel-storage.com/gitfindr-removebg-preview-GH0PDgHySVEHMTcagoEKXoJAZVcH74.png"
              alt="GitFindr Logo"
              width={60}
              height={60}
              className="h-8 w-auto md:h-10"
            />
            <span className="text-5xl font-bold rajdhani-bold text-[#1E40AF]">
              GitFindr
            </span>
          </Link>
          <div className="flex items-center space-x-4 md:space-x-6 ">
            <Link
              to="/about"
              className="text-gray-600 text-xl hover:text-[#1E40AF] ibm-plex-sans"
            >
              About
            </Link>
            <button
              to="/register"
              onClick={() => setShowModal(true)}
              className="text-gray-600 hover:text-[#1E40AF] text-xl ibm-plex-sans"
            >
              Register
            </button>
            {showModal && <RegisterModal onClose={() => setShowModal(false)} />}
            <a
              href="https://github.com"
              className=" ibm-plex-sans flex items-center text-xl space-x-2 rounded-full bg-black px-4 py-2 text-sm text-white transition-transform hover:scale-105"
            >
              <span>★</span>
              <span>Github</span>
            </a>
          </div>
        </nav>

        {/* Hero Section */}
        <main className="flex flex-col items-center justify-center px-4 py-12 text-center md:py-20">
          <motion.div
            initial={{ scale: 0.8, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            transition={{ duration: 0.5 }}
            className="mb-8 flex flex-row items-center justify-center"
          >
            <p className="rajdhani-bold text-8xl text-[#1E3A8A]">GitFindr</p>
            <img
              src="https://hebbkx1anhila5yf.public.blob.vercel-storage.com/gitfindr-removebg-preview-GH0PDgHySVEHMTcagoEKXoJAZVcH74.png"
              alt="GitFindr Large Logo"
              width={100}
              height={100}
              className="h-auto w-48 md:w-64"
            />
          </motion.div>

          <motion.h1
            initial={{ y: 20, opacity: 0 }}
            animate={{ y: 0, opacity: 1 }}
            transition={{ delay: 0.2 }}
            className="mb-4 text-3xl font-bold md:text-4xl lexend-deca"
          >
            The Google for GitHub
          </motion.h1>

          <motion.div
            initial={{ y: 20, opacity: 0 }}
            animate={{ y: 0, opacity: 1 }}
            transition={{ delay: 0.3 }}
            className="mb-8 text-xl md:text-2xl lexend-deca"
          >
            Unlock the World of{" "}
            <span className="rounded-lg bg-blue-500 px-2 py-1 text-white">
              Open Source
            </span>
            -ry
          </motion.div>

          {/* Search Bar */}
          <motion.div
            initial={{ y: 20, opacity: 0 }}
            animate={{ y: 0, opacity: 1 }}
            transition={{ delay: 0.4 }}
            className="relative mb-12 w-full max-w-2xl"
          >
            <input
              type="text"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              placeholder="What are we building today?"
              className="ibm-plex-sans w-full rounded-full border-black border-2 px-6 py-6 pr-12 text-lg shadow-sm transition-shadow focus:border-blue-500 focus:outline-none focus:ring-2 focus:ring-blue-200"
            />
            <button
              onClick={searchQuery}
              className="absolute right-3 top-1/2 -translate-y-1/2 rounded-full bg-black px-7 py-3 text-4xl text-white transition-transform hover:scale-110"
            >
              ↵
            </button>
          </motion.div>

          {/* Call to Action */}
          <motion.div
            initial={{ y: 20, opacity: 0 }}
            animate={{ y: 0, opacity: 1 }}
            transition={{ delay: 0.5 }}
            className="text-center flex flex-row"
          >
            <p className="ibm-plex-sans m-4 text-lg text-gray-600">
              Finished A Project?
            </p>
            <button className="ibm-plex-sans rounded-xl bg-black px-6 py-2 text-white text-xl transition-transform hover:scale-105">
              Register Your Repo
            </button>
          </motion.div>
        </main>
        <div className="flex items-center justify-center min-h-screen p-4 bg-white">
          <div
            id="repos"
            className="grid w-full max-w-6xl grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-3"
          >
            {repoDetails.map((repo, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, scale: 0.95 }}
                animate={{ opacity: 1, scale: 1 }}
                transition={{ duration: 0.3 }}
                className="p-6 border border-gray-200 rounded-2xl shadow-xl bg-white hover:shadow-2xl transition-all"
              >
                <a
                  href={repo.html_url}
                  target="_blank"
                  rel="noopener noreferrer"
                  className="text-xl font-bold text-blue-600 hover:underline"
                >
                  {repo.name}
                </a>
                <p className="mt-2 text-gray-700">
                  {repo.description || "No description available."}
                </p>
                <p className="mt-2 text-gray-500">👤 {repo.owner?.login}</p>
                <div className="flex items-center justify-between mt-3 text-gray-600">
                  <span>⭐ {repo.stargazers_count}</span>
                  <span>🍴 {repo.forks_count}</span>
                </div>
              </motion.div>
            ))}
          </div>
        </div>
        {showButton && (
          <button
            onClick={() => window.scrollTo({ top: 0, behavior: "smooth" })}
            className="fixed bottom-6 right-6 z-50 p-3 bg-blue-600 text-white rounded-full shadow-lg hover:bg-blue-700 transition-all flex items-center justify-center"
          >
            <ArrowUp size={24} />
          </button>
        )}
      </div>
    </Router>
  );
}
