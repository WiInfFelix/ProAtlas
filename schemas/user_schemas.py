from pydantic import BaseModel

class UserBase(BaseModel):
    username: str
    email: str


class UserIn(UserBase):
    password: str

class UserOut(UserBase):
    id: int