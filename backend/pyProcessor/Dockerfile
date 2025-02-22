# Use official Python image as base
FROM python:3.10-slim

# Set the working directory in the container
WORKDIR /app/backend/pyProcessor

# Copy only the requirements file first (to leverage Docker cache)
COPY backend/pyProcessor/requirements.txt .

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Download spaCy model
RUN python -m spacy download en_core_web_sm

# Copy the entire application
COPY backend/pyProcessor/ .

# Expose the FastAPI port
EXPOSE 8000

# Run FastAPI application using uvicorn
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
