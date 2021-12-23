from database import engine
from fastapi import FastAPI
from routes import user_router, exercise_router
import models

models.Base.metadata.create_all(bind=engine)


app = FastAPI()

app.include_router(user_router)
app.include_router(exercise_router)


@app.get("/")
async def root():
    return {"message": "Hello World"}
