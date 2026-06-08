import os
import base64
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
from dotenv import load_dotenv

load_dotenv()

SECRET_KEY = os.getenv("SECRET_KEY", "clavesecretade32caracteresunimayor").encode('utf-8')
SECRET_KEY = SECRET_KEY[:32].ljust(32, b'\0')

def encrypt_password(password: str) -> str:
    """Cifra una contraseña en texto plano usando AES y la devuelve en Base64"""
    cipher = AES.new(SECRET_KEY, AES.MODE_ECB) 
    padded_data = pad(password.encode('utf-8'), AES.block_size)
    encrypted_bytes = cipher.encrypt(padded_data)
    return base64.b64encode(encrypted_bytes).decode('utf-8')

def decrypt_password(encrypted_password: str) -> str:
    """Descifra una contraseña en Base64 usando AES"""
    cipher = AES.new(SECRET_KEY, AES.MODE_ECB)
    encrypted_bytes = base64.b64decode(encrypted_password.encode('utf-8'))
    decrypted_bytes = cipher.decrypt(encrypted_bytes)
    return unpad(decrypted_bytes, AES.block_size).decode('utf-8')