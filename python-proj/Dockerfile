# Use an official Python runtime as a parent image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the current directory contents into the container at /usr/src/app
COPY . .

RUN pip install -r requirements.txt

# Make port 5000 available to the world outside this container
EXPOSE 8080

# Command to run the application
CMD ["python", "./main.py"]
