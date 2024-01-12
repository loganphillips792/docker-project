# Running the backend Flask app

## Installing Python depdencies

1. ```python3 -m venv ~/Desktop/FlaskEnv```
2. ```source ~/Desktop/FlaskEnv/bin/activate``` - this line activates the virtual environment so your Python will use an packages that are installed in it
3. ```which pip``` to verify what is being used (Should point to the one from the virtual environment)
4. ```~/Desktop/FlaskEnv/bin/python3 -m pip install --upgrade pip```
5. ```pip install -r requirements.txt```

## Running application

1. Go to project root
2. Create environment file (.env)
3. Enable virtual environment
4. python3 main.py

## Building in Docker

`docker build  -t mypythonapp:latest .`

`docker run mypythonapp:latest`

## Pushing to Dockerhub

1. Make sure you have a Dockerhub account
2. Log in to Docker Hub from the CLI: `docker login`
3. Tag the image: `docker tag mypythonapp dockedupstream/mypythonapp`
4. Push the image: `docker push dockedupstream/mypythonapp`

Pushing a new tag to the repository (besides the default latest one): `docker push dockedupstream/mypythonapp:tagname`
