from logging import log
from models import User, Exercise, Workouts
import peewee as p



def setup_db():
    try:
        db = p.SqliteDatabase("database.db", pragmas={"foreign_keys": 1})

        db.connect()
        db.create_tables([User, Exercise, Workouts])
    except Exception as e:
        print(f"There was an error instancing the models: {e}")
        db.close()
        raise SystemExit
