import { useState } from "react";

export default function RegisterModal({ onClose }) {
  const [repoUrl, setRepoUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState("");

  const handleSubmit = async () => {
    if (!repoUrl.trim()) return;

    setLoading(true);
    setMessage("");

    try {
      const response = await fetch("http://localhost:8000/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ fullname: repoUrl }),
      });

      const data = await response.json();

      setMessage("Indexing complete! ✅");
    } catch (error) {
      setMessage("Failed to connect to the server. ❌");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white p-5 rounded-lg shadow-lg w-80">
        <h2 className="text-lg font-semibold">Write repo fullname</h2>
        <input
          type="text"
          placeholder="https://github.com/user/repo"
          className="border p-2 w-full mt-2 rounded"
          value={repoUrl}
          onChange={(e) => setRepoUrl(e.target.value)}
        />
        <div className="flex justify-between mt-4">
          <button
            onClick={onClose}
            className="bg-gray-500 text-white px-4 py-2 rounded"
          >
            Cancel
          </button>
          <button
            onClick={handleSubmit}
            className="bg-blue-600 text-white px-4 py-2 rounded"
            disabled={loading}
          >
            {loading ? "Loading..." : "Submit"}
          </button>
        </div>
        {message && <p className="text-center mt-2">{message}</p>}
      </div>
    </div>
  );
}
