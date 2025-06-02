import re
from text_prettifier import TextPrettifier
prettifier = TextPrettifier()
def clean_text(text:str)->str:
    text = prettifier.remove_urls(text)
    text = prettifier.remove_emojis(text)
    text = prettifier.remove_html_tags(text)
    text = prettifier.remove_numbers(text)
    #text = prettifier.remove_stopwords(text)
    text = remove_code_blocks(text)
    text = prettifier.remove_special_chars(text)
    return text
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
