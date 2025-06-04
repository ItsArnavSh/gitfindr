from src.internal.models.repo import Repository
from src.internal.crud.repo import create_repository
import requests
from src.internal.logger import logger
import uuid

def store_link(fullname: str)->Repository:
    url = f"https://api.github.com/repos/{fullname}"
    repo_data = load_link(url)
    readme_content = fetch_readme(fullname)

    if repo_data is None:
        logger.error(f"Repo Data is None for {url}")
        return Repository()

    # Handle topics: some repos may not have them
    topics = repo_data.get("topics", [])
    topics_str = ",".join(topics) if isinstance(topics, list) else ""
    size = 0
    if readme_content:
        size = len(readme_content.split())
    new_repo = Repository(
        url=url,
        readme_content=readme_content,
        name=repo_data.get("name"),
        fullname=repo_data.get("full_name"),
        description=repo_data.get("description", ""),
        topics=topics_str,
        language=repo_data.get("language"),
        stars=repo_data.get("stargazers_count", 0),
        forks=repo_data.get("forks_count", 0),
        issues=repo_data.get("open_issues_count", 0),
        watchers=repo_data.get("subscribers_count", 0),
        archived=repo_data.get("archived", False),
        forked=repo_data.get("fork", False),
        size=size
    )

    # Persist to database
    return create_repository(new_repo)

def load_link(url: str) -> dict | None:
    try:
        response = requests.get(url)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        logger.error(f"Failed to load URL {url}: {e}")
        return None
    except ValueError:
        logger.error(f"Response from {url} is not valid JSON.")
        return None

def fetch_readme(full_name: str) -> str | None:
    url = f"https://raw.githubusercontent.com/{full_name}/main/README.md"
    logger.info(f"Url for readme:{url}")
    try:
        response = requests.get(url)
        response.raise_for_status()
        return response.text
    except requests.exceptions.RequestException as e:
        logger.error(f"Error fetching README for {full_name}: {e}")
        return None
