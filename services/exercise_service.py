from models import Exercise


def create_new_exercise(exercise, db):
    db_exec = Exercise(
        exercise_name=exercise.exercise_name,
        description=exercise.description,
        target_muscle=exercise.target_muscle,
    )
    db.add(db_exec)
    db.commit()
    db.refresh(db_exec)
    return db_exec


def get_all_exercises(db):
    return db.query(Exercise).all()


def get_exercise_by_id(exec_id, db):
    return db.query(Exercise).filter(Exercise.id == exec_id).first()
