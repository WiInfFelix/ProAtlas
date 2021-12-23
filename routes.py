from fastapi import APIRouter
from schemas.user_schemas import UserIn, UserOut

user_router = APIRouter()
exercise_router = APIRouter()
workout_router = APIRouter()

@user_router.post("/users/")
async def create_new_user(user: UserIn):
    return "Create users"

@user_router.get("/users/")
async def get_all_users():
    return "All users"