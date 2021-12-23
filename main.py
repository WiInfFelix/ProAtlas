from fastapi import FastAPI
from database import setup_db
from routes import user_router


setup_db()

app = FastAPI()

app.include_router(user_router)

@app.get("/")
async def root():
    return {"message": "Hello World"}