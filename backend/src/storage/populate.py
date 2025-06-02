from internal.models.repo import Repository
from internal.crud.repo import create_repository
import requests
from internal.logger import logger

def store_link(fullname:str):
    url = f"https://api.github.com/repos/{fullname}"
    repo_data = load_link(url)
    readme_content = fetch_readme(fullname)

def load_link(url: str) -> dict | None:
    try:
        response = requests.get(url)
        response.raise_for_status()  # Raises HTTPError for bad responses (4xx/5xx)
        return response.json()
    except requests.exceptions.RequestException as e:
        logger.error(f"Failed to load URL {url}: {e}")
        return None
    except ValueError:
        logger.error(f"Response from {url} is not valid JSON.")
        return None
def fetch_readme(full_name: str) -> str | None:
    url = f"https://api.github.com/repos/{full_name}/readme"
    headers = {"Accept": "application/vnd.github.v3.raw"}  # Gets the raw content of the README

    try:
        response = requests.get(url, headers=headers)
        response.raise_for_status()
        return response.text  # The raw README content (e.g., Markdown)
    except requests.exceptions.RequestException as e:
        print(f"Error fetching README for {full_name}: {e}")
        return None
