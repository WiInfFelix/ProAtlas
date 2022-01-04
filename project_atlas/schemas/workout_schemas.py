from pydantic.main import BaseModel

from schemas.exercise_schema import ExerciseOut
from schemas.user_schemas import UserOut


class WorkoutBase(BaseModel):
    exercise_id: int
    sets: int
    repetitions: int
    weight: int

    class Config:
        orm_mode = True


class WorkoutIn(WorkoutBase):
    pass


class WorkoutOut(WorkoutBase):
    id: int
    user_id: int
    exercise: ExerciseOut
    user: UserOut
