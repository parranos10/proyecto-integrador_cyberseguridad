from pydantic import EmailStr
from sqlmodel import SQLModel, Field
from typing import Optional

class User(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    username: str = Field(index=True)
    email: str = Field(unique=True)
    password: str 

class UserRegister(SQLModel):
    username: str
    email: EmailStr
    password: str

class UserResponse(SQLModel):
    id: int
    username: str
    email: str