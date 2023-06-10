# DB Connect

Run the backend first, it will listen on localhost:8000.

Then run the frontend, it will work on localhost:5173.



## Backend

### How to run

```shell
cd backend
pip install -r requirements
python manage.py migrate
python manage.py runserver 0.0.0.0:8000
```



## Frontend

### How to run

```shell
cd frontend
npm install
npm run dev
```

Then access http://localhost:5173 in the browser.
