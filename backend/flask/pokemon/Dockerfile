# Use the official Python image
FROM python:3.10-slim

# Set the working directory
WORKDIR /app

# Copy the Flask app and requirements file
COPY *.py requirements.txt /app/
COPY .env requirements.txt /app/

# Install dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Expose port 8888 for Flask
EXPOSE 8888

# Run the Flask app
CMD ["python", "app.py"]
