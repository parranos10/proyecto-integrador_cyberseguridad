import os
import bcrypt
from dotenv import load_dotenv

load_dotenv()

PEPPER = os.getenv("PEPPER")

def generate_salt():
    return bcrypt.gensalt().decode("utf-8")

def hash_password(password: str, salt: str):
    password_peppered = password + PEPPER
    hashed = bcrypt.hashpw(
        password_peppered.encode("utf-8"),
        salt.encode("utf-8")
    )
    return hashed.decode("utf-8")

def verify_password(password: str, salt: str, stored_hash: str):
    password_peppered = password + PEPPER
    return bcrypt.checkpw(
        password_peppered.encode("utf-8"),
        stored_hash.encode("utf-8")
    )