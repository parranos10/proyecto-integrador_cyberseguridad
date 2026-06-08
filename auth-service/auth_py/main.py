from fastapi import FastAPI, HTTPException
from sqlmodel import SQLModel, Session, create_engine, select
from pydantic import BaseModel

from models import User
from auth import generate_salt, hash_password, verify_password

app = FastAPI(title="API de Autenticación Segura")

sqlite_file_name = "database.db"
sqlite_url = f"sqlite:///{sqlite_file_name}"

engine = create_engine(sqlite_url, echo=True)

class UserRequest(BaseModel):
    username: str
    password: str

@app.on_event("startup")
def on_startup():
    SQLModel.metadata.create_all(engine)

@app.post("/register")
def register(user_data: UserRequest):
    with Session(engine) as session:
        existing_user = session.exec(
            select(User).where(User.username == user_data.username)
        ).first()

        if existing_user:
            raise HTTPException(status_code=400, detail="El usuario ya existe")

        salt = generate_salt()
        password_hash = hash_password(user_data.password, salt)

        user = User(
            username=user_data.username,
            password_hash=password_hash,
            salt=salt
        )

        session.add(user)
        session.commit()
        session.refresh(user)

        return {
            "message": "Usuario registrado correctamente",
            "username": user.username
        }

@app.post("/login")
def login(user_data: UserRequest):
    with Session(engine) as session:
        user = session.exec(
            select(User).where(User.username == user_data.username)
        ).first()

        if not user:
            raise HTTPException(status_code=401, detail="Credenciales incorrectas")

        valid_password = verify_password(
            user_data.password,
            user.salt,
            user.password_hash
        )

        if not valid_password:
            raise HTTPException(status_code=401, detail="Credenciales incorrectas")

        return {
            "message": "Inicio de sesión exitoso",
            "username": user.username
        }