from fastapi import HTTPException

from auth import auth
from models import User, Workouts, Exercise
from schemas.workout_schemas import WorkoutOut


def create_new_workout(workout, db, token):
    user = auth.get_current_user(db=db, token=token)

    workout_db = Workouts(
        user_id=user.id,
        exercise_id=workout.exercise_id,
        sets=workout.sets,
        repetitions=workout.repetitions,
        weight=workout.weight,
    )

    db.add(workout_db)
    db.commit()
    db.refresh(workout_db)

    user = db.query(User).filter(User.id == workout_db.user_id).first()
    exercise = db.query(Exercise).filter(Exercise.id == workout.exercise_id).first()

    if not exercise:
        raise HTTPException(status_code=404, detail="No Exercise found...")

    res = WorkoutOut(
        exercise_id=workout.exercise_id,
        user_id=user.id,
        id=workout_db.id,
        exercise=exercise,
        user=user,
        sets=workout_db.sets,
        repetitions=workout_db.repetitions,
        weight=workout_db.weight,
    )

    return res
