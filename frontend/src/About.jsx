"use client";

import { useEffect, useState } from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import remarkMath from "remark-math";
import rehypeKatex from "rehype-katex";
import "katex/dist/katex.min.css"; // Import KaTeX styles
import { Loader2 } from "lucide-react";

export default function AboutPage() {
  const [markdown, setMarkdown] = useState("");
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchReadme = async () => {
      try {
        const response = await fetch(
          "https://raw.githubusercontent.com/ItsArnavSh/gitfindr/main/README.md",
        );
        const text = await response.text();
        setMarkdown(text);
      } catch (error) {
        console.error("Error fetching README:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchReadme();
  }, []);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <Loader2 className="w-8 h-8 animate-spin text-blue-700" />
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-white">
      <div className="max-w-4xl mx-auto px-4 py-8">
        <div className="flex mb-8 flex-row items-center justify-center">
          <img
            src="https://hebbkx1anhila5yf.public.blob.vercel-storage.com/gitfindr-uQ5eBXWlUeDOcgz5ubw2hdCa0IBcIb.png"
            alt="GitFindr Logo"
            className="h-12"
          />
          <p className="text-5xl rajdhani-bold p-2 text-blue-700">About</p>
        </div>
        <article className="prose prose-blue max-w-none">
          <ReactMarkdown
            remarkPlugins={[remarkGfm, remarkMath]}
            rehypePlugins={[rehypeKatex]}
            components={{
              h1: ({ node, ...props }) => (
                <h1
                  className="text-4xl font-bold text-blue-700 text-center mb-8"
                  {...props}
                />
              ),
              h2: ({ node, ...props }) => (
                <h2
                  className="text-2xl font-semibold text-blue-600 mt-8 mb-4"
                  {...props}
                />
              ),
              h3: ({ node, ...props }) => (
                <h3
                  className="text-xl font-semibold text-blue-500 mt-6 mb-3"
                  {...props}
                />
              ),
              a: ({ node, ...props }) => (
                <a
                  className="text-blue-700 hover:text-blue-500 underline"
                  {...props}
                />
              ),
              p: ({ node, ...props }) => (
                <p className="text-gray-700 mb-4 leading-relaxed" {...props} />
              ),
              ul: ({ node, ...props }) => (
                <ul
                  className="list-disc list-inside mb-4 text-gray-700"
                  {...props}
                />
              ),
              ol: ({ node, ...props }) => (
                <ol
                  className="list-decimal list-inside mb-4 text-gray-700"
                  {...props}
                />
              ),
              code: ({ node, ...props }) => (
                <code
                  className="bg-gray-100 rounded px-1 py-0.5 text-blue-700"
                  {...props}
                />
              ),
              pre: ({ node, ...props }) => (
                <pre
                  className="bg-gray-100 rounded-lg p-4 overflow-x-auto mb-4"
                  {...props}
                />
              ),
            }}
          >
            {markdown}
          </ReactMarkdown>
        </article>
      </div>
    </div>
  );
}
