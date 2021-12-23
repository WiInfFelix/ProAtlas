from schemas import user_schemas
from models import User
from sqlalchemy.orm import Session

def create_user(user: user_schemas.UserIn, db: Session):
    db_user = User(username=user.username, email=user.email, password=user.password)
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