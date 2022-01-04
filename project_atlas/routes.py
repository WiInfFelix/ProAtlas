from fastapi.security import OAuth2PasswordRequestForm
from sqlalchemy.orm.session import Session
from fastapi import APIRouter, Depends

from auth import auth
from auth.auth import Token
from schemas.user_schemas import UserIn, UserOut, UserLogin
from schemas.exercise_schema import ExerciseIn, ExerciseOut
from typing import List

from schemas.workout_schemas import WorkoutOut, WorkoutIn
from services import user_service, exercise_service, workout_service
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


@user_router.post("/token", response_model=Token)
async def get_user_token(form_data: OAuth2PasswordRequestForm = Depends(), db: Session = Depends(get_db)):
    return user_service.get_user_token(db=db, user=form_data)


@exercise_router.post("/exercises/", response_model=ExerciseOut, tags=["Exercises"])
async def create_new_exercise(exercise: ExerciseIn, db: Session = Depends(get_db)):
    return exercise_service.create_new_exercise(exercise, db)


@exercise_router.get(
    "/exercises/", response_model=List[ExerciseOut], tags=["Exercises"]
)
async def get_all_exercises(db: Session = Depends(get_db)):
    return exercise_service.get_all_exercises(db)


@exercise_router.get(
    "/exercises/{exec_id}", response_model=ExerciseOut, tags=["Exercises"]
)
async def get_exercise_by_id(exec_id: int, db: Session = Depends(get_db)):
    return exercise_service.get_exercise_by_id(exec_id, db)


@workout_router.post("/user/{user_id}/workouts", response_model=WorkoutOut, tags=["Workouts"])
async def create_new_workout(workout: WorkoutIn, db: Session = Depends(get_db), token: str = Depends(auth.oauth2_scheme)):
    return workout_service.create_new_workout(workout, db, token)
