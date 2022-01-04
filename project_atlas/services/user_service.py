from datetime import timedelta

from fastapi import HTTPException, status
from sqlalchemy.orm import Session

from auth import auth
from models import User
from schemas import user_schemas
from schemas.user_schemas import UserLogin


def create_user(user: user_schemas.UserIn, db: Session):
    hash_pass = auth.get_password_hash(user.password)

    db_user = User(username=user.username, email=user.email, password=hash_pass)
    db.add(db_user)
    db.commit()
    db.refresh(db_user)
    return db_user


def get_all_users(db: Session):
    return db.query(User).all()


def get_user_by_id(user_id: int, db: Session):
    return db.query(User).filter(User.id == user_id).first()


def get_user_by_email(email: str, db: Session):
    return db.query(User).filter(User.email == email).first()


def get_user_token(user: UserLogin, db: Session):
    user = auth.authenticate_user(db, user.username, user.password)
    if not user:
        raise HTTPException(
            status_code=status.HTTP_401_UNAUTHORIZED,
            detail="Incorrect username or password",
            headers={"WWW-Authenticate": "Bearer"},
        )
    access_token_expires = timedelta(minutes=auth.ACCESS_TOKEN_EXPIRY_MINUTES)
    access_token = auth.create_access_token(data={"sub": user.username}, expires_delta=access_token_expires)
    return {"access_token": access_token, "token_type": "bearer"}
