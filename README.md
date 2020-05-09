# Installation Requirements
- Python 3.6
- Pip 20.1
- Virtualenv
- Postgres 10.6

# Postgres Setup
- Install postgreSQL
```
$ sudo apt-get update
$ sudo apt-get install python3-pip python3-dev libpq-dev postgresql postgresql-contrib
```
- Run the following commands
```
$ sudo -u postgres psql
$ CREATE DATABASE trello_dev;
$ CREATE USER jatin WITH PASSWORD ' ';
$ ALTER ROLE jatin SET client_encoding TO 'utf8';
$ ALTER ROLE jatin SET default_transaction_isolation TO 'read committed';
$ ALTER ROLE jatin SET timezone TO 'Asia/Kolkata';
```

# Installation Guide
- Switch on your Virtualenv `source bin/activate`
- git clone the repository. `git@github.com:jatingoyal759/sequoia-backend-assignment.git`
- `$ cd sequoia-backend-assignment/trello/`.
- Install all dependencies `$ pip install -r requirements.txt`

# Migrations, SuperUser, Local Server and Admin Panel
- Run migrations
```
$ python manage.py makemigrations
$ python manage.py migrate
```

- Create SuperUser
```
$ python manage.py createsuperuser
```

- Run Server
```
$ python manage.py runserver
```
- The server is running at [127.0.0.1:8000](http://127.0.0.1:8000/)

- Visit [Admin Panel](http://127.0.0.1:8000/admin) and log in with superuser credentials.

# Setting Up Environment Variables
- For the assignment purposes .env file is uploaded to git.