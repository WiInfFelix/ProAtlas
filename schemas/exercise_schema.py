from pydantic import BaseModel


class ExerciseBase(BaseModel):
    exercise_name: str
    description: str
    target_muscle: str

    class Config:
        orm_mode = True


class ExerciseOut(ExerciseBase):
    id: int


class ExerciseIn(ExerciseBase):
    pass
