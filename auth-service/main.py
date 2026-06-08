import os
from fastapi import FastAPI, HTTPException, status, Depends
from fastapi.staticfiles import StaticFiles
from dotenv import load_dotenv
from sqlmodel import SQLModel, Field, create_engine, Session, select

from models.user import User, UserRegister, UserResponse
from services.crypto_service import encrypt_password

load_dotenv()

DATABASE_URL = "sqlite:///./usuarios.db"
engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})

def init_db():
    SQLModel.metadata.create_all(engine)

app = FastAPI(title="Sistema de Autenticación Seguro")

@app.on_event("startup")
def on_startup():
    init_db()

def get_session():
    with Session(engine) as session:
        yield session

@app.get("/")
def home():
    return {"message": "API de Autenticación con SQLModel y AES activa"}

@app.post("/users/register", response_model=UserResponse, status_code=status.HTTP_201_CREATED)
def registrar_usuario(user_data: UserRegister, session: Session = Depends(get_session)):
    statement = select(User).where(User.email == user_data.email)
    usuario_existente = session.exec(statement).first()
    if usuario_existente:
        raise HTTPException(status_code=400, detail="El correo ya está registrado")
    
    password_cifrada = encrypt_password(user_data.password)
    
    nuevo_usuario = User(
        username=user_data.username,
        email=user_data.email,
        password=password_cifrada 
    )
    
    session.add(nuevo_usuario)
    session.commit()
    session.refresh(nuevo_usuario)
    
    return nuevo_usuario

@app.get("/users", response_model=list[UserResponse])
def listar_usuarios(session: Session = Depends(get_session)):
    usuarios = session.exec(select(User)).all()
    return usuarios