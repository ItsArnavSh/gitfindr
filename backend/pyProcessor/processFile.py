import requests
import re
import spacy

from spacy.lang.en.stop_words import STOP_WORDS
nlp = spacy.load("en_core_web_sm")

def processRepo(repo_url):
    print(f"Processing: {repo_url}")
    readme_text = download_readme(repo_url)
    if readme_text:
        cleaned_text = clean_text(readme_text)
        words = extract_words(cleaned_text)
        return words
    else:
        print("README not found for", repo_url)

def download_readme(repo_url):
    """Downloads the README file from a GitHub repository, trying 'main' first and then 'master'."""
    for branch in ["main", "master"]:
        readme_url = repo_url.replace("github.com", "raw.githubusercontent.com") + f"/{branch}/README.md"
        print(f"Trying: {readme_url}")
        try:
            response = requests.get(readme_url, timeout=10)
            if response.status_code == 200:
                return response.text
        except requests.RequestException:
            pass
    return None

def extract_words(text):
    """Extracts relevant words from text using spaCy."""
    doc = nlp(text)
    def is_valid_word(word):
        return (
            re.search(r"[a-zA-Z]", word) and
            word.lower() not in STOP_WORDS and
            not word.startswith("-") and
            word.lower() not in {"git", "install", "run", "add"}
        )
    return [token.lemma_.lower() for token in doc if is_valid_word(token.text)]



def clean_text(text):
    """Cleans and processes text, removing unnecessary noise from GitHub READMEs."""
    # Remove fenced code blocks (both ``` and ~~~)
    text = re.sub(r"```[\s\S]*?```", "", text)
    text = re.sub(r"~~~[\s\S]*?~~~", "", text)

    # Remove inline code (`code`)
    text = re.sub(r"`([^`]+)`", "", text)

    # Remove markdown links but keep visible text: [Text](https://example.com) → Text
    text = re.sub(r"\[([^\]]+)\]\(https?://\S+\)", r"\1", text)

    # Remove standalone URLs
    text = re.sub(r"https?://\S+", "", text)

    # Remove badges and shields (images with markdown or HTML)
    text = re.sub(r"!\[.*?\]\(https?://\S+\)", "", text)  # Markdown images
    text = re.sub(r"<img[^>]+>", "", text)  # HTML images

    # Remove words containing special symbols (e.g., some_code(), #hashtag)
    text = re.sub(r"\b\S*[^\w\s]\S*\b", "", text)

    # Define unwanted headings to remove (and their content)
    unwanted_headings = [
        "installation", "install", "setup", "getting started", "license", "licensing",
        "contributing", "contribution", "usage", "how to use", "about", "disclaimer",
        "faq", "frequently asked questions", "support", "troubleshooting", "table of contents"
    ]
    pattern = re.compile(r"(?m)^#{1,6}\s*(?:" + "|".join(unwanted_headings) + r")\b.*(?:\n(?!#{1,6} ).*)*", re.IGNORECASE)
    text = re.sub(pattern, "", text)

    # Remove table of contents (TOC) sections (Markdown list format)
    text = re.sub(r"(?m)^\s*[-*]\s*\[.*?\]\(#.*?\)\s*$", "", text)

    # Remove HTML tags
    text = re.sub(r"<[^>]+>", "", text)

    # Remove emojis and special characters
    text = re.sub(r"[^\w\s.,!?]", "", text)

    # Remove extra newlines and whitespace
    text = re.sub(r"\n\s*\n+", "\n\n", text).strip()

    return text

def processText(data):
    return extract_words(data)
