import peewee as p

db = p.SqliteDatabase("database.db")

class User(p.Model):
    id = p.BigIntegerField(primary_key=True)
    username = p.CharField(max_length=128)
    email = p.CharField(max_length=256)
    password = p.CharField()

    class Meta:
        database = db


class Exercise(p.Model):
    id = p.BigIntegerField(primary_key=True)
    exercise_name = p.CharField(max_length=256)
    description = p.TextField()
    target_muscle = p.CharField()

    class Meta:
        database = db

class Workouts(p.Model):
    id = p.BigIntegerField(primary_key=True)
    user_id = p.ForeignKeyField(User, backref="workouts")
    exercise_id = p.ForeignKeyField(Exercise, backref="workouts")
    sets = p.SmallIntegerField()
    repetitions = p.SmallIntegerField()
    weight = p.FloatField()

    class Meta:
        database = db
