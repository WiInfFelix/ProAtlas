from pydantic import BaseModel


class UserBase(BaseModel):
    username: str
    email: str

    class Config:
        orm_mode = True


class UserIn(UserBase):
    password: str


class UserOut(UserBase):
    id: int
