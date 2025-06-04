import re
from typing import List
from src.internal.logger import logger
from text_prettifier import TextPrettifier
prettifier = TextPrettifier()
def clean_text(text: str) -> str:
    logger.info(text)

    text = prettifier.remove_urls(text)
    logger.debug("After removing URLs: %s", text)

    text = prettifier.remove_emojis(text)
    logger.debug("After removing emojis: %s", text)

    text = prettifier.remove_html_tags(text)
    logger.debug("After removing HTML tags: %s", text)

    # Uncomment if needed
    # text = prettifier.remove_numbers(text)
    # logger.debug("After removing numbers: %s", text)

    # text = prettifier.remove_stopwords(text)
    # logger.debug("After removing stopwords: %s", text)

    text = remove_code_blocks(text)
    logger.debug("After removing code blocks: %s", text)

    text = prettifier.remove_special_chars(text)
    logger.debug("After removing special characters: %s", text)

    final_text = text.lower()
    logger.info(final_text)

    return final_text

def remove_code_blocks(markdown_text):
    """
    Removes all fenced code blocks from a Markdown string.
    Code blocks start and end with triple backticks ```.
    """
    # Pattern to match code blocks between triple backticks
    code_block_pattern = r"```.*?```"

    # Remove all code blocks, including multiline, using re.DOTALL
    cleaned_text = re.sub(code_block_pattern, "", markdown_text, flags=re.DOTALL)

    return cleaned_text.strip()

def keyword_extraction(text:str)->List[str]:
    text = prettifier.remove_stopwords(text)
    return text.split()
