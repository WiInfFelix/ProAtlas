from pydantic import BaseModel


class ExerciseBase(BaseModel):
    name: str
    description: str
    target_muscle: str


class ExerciseOut(ExerciseBase):
    id: int


class ExerciseIn(ExerciseBase):
    pass
