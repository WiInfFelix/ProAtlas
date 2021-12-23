from sqlalchemy.orm.session import Session
from fastapi import APIRouter, Depends
from schemas.user_schemas import UserIn, UserOut
from schemas.exercise_schema import ExerciseIn, ExerciseOut
from typing import List
from services import user_service
from database import SessionLocal


def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


user_router = APIRouter()
exercise_router = APIRouter()
workout_router = APIRouter()


@user_router.post("/users/", response_model=UserOut, tags=["Users"])
async def create_new_user(user: UserIn, db: Session = Depends(get_db)):
    return user_service.create_user(user, db)


@user_router.get("/users/", response_model=List[UserOut], tags=["Users"])
async def get_all_users(db: Session = Depends(get_db)):
    return user_service.get_all_users(db)


@user_router.get("/users/{user_id}", response_model=UserOut, tags=["Users"])
async def get_user_by_id(user_id: int, db: Session = Depends(get_db)):
    return user_service.get_user_by_id(user_id, db)


@exercise_router.post("/exercises/", response_model=ExerciseOut, tags=["Exercises"])
async def create_new_exercise(exercise: ExerciseIn, db: Session = Depends(get_db)):
    pass


@exercise_router.get(
    "/exercises/", response_model=List[ExerciseOut], tags=["Exercises"]
)
async def get_all_exercises(db: Session = Depends(get_db)):
    pass


@exercise_router.get(
    "/exercises/{exec_id}", response_model=ExerciseOut, tags=["Exercises"]
)
async def get_exercise_by_id(user_id: int, db: Session = Depends(get_db)):
    pass
