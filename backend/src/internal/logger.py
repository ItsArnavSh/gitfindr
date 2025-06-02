# internal/logger.py
from loguru import logger

logger.add(
    "logs/gitfindr.log",
    rotation="10 MB",
    retention="7 days",
    compression="zip",
    format="{time} | {level} | {message}",
    level="INFO",              # or "DEBUG" during development
)
