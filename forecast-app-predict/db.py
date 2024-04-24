import argparse
import multiprocessing

from entities import MakePredictRequestEntity
import psycopg2


class DB(object):
    def __init__(self, config: argparse.Namespace, queue: multiprocessing.Queue):
        self.db = psycopg2.connect(
            host=config.db_address.split(':')[0],
            database=config.db_name,
            user=config.db_user,
            password=config.db_password,
            port=config.db_address.split(':')[1]
        )
        if queue is not None:
            self.queue = queue
            self.save_prediction()

    def save_prediction(self):
        while True:
            request: MakePredictRequestEntity = self.queue.get(block=True)

            print(f'got request {MakePredictRequestEntity}')

            cursor = self.db.cursor()

            # get user id
            cursor.execute(
                "SELECT id FROM users WHERE username = %s;",
                (request.username,)
            )
            if cursor.rowcount == 0:
                print(f"no such user {request.username}")
                cursor.close()
                return
            user_id = cursor.fetchone()[0]

            # get unit id
            cursor.execute(
                "SELECT id FROM measurement_units WHERE name = %s;",
                (request.unit,)
            )
            if cursor.rowcount == 0:
                cursor.execute(
                    "INSERT INTO measurement_units (name) VALUES (%s) RETURNING id;",
                    (request.unit,)
                )
                self.db.commit()
                if cursor.rowcount == 0:
                    print("error occurred during inserting into measurement_units")
                    cursor.close()
                    return
            unit_id = cursor.fetchone()[0]

            # insert time series
            cursor.execute(
                "INSERT INTO time_series (user_id, unit_id, name,  period, prediction_start) VALUES (%s, %s, %s, %s, %s) RETURNING id;",
                (user_id, unit_id, request.name, request.period, int(len(request.items) * 0.8))
            )
            self.db.commit()
            if cursor.rowcount == 0:
                print("error occurred during inserting into time series")
                cursor.close()
                return
            time_series_id = cursor.fetchone()[0]

            cursor.executemany(
                "INSERT INTO records (series_id, ts, value) VALUES (%s, %s, %s);",
                [(time_series_id, item.ts, item.value) for item in request.items]
            )
            self.db.commit()
            if cursor.rowcount == 0:
                print("error occurred during inserting time series items")
                cursor.close()
                return

            cursor.close()

    def get_prediction(self, username: str, name: str) -> (str, int, list[tuple[int, float]]):
        cursor = self.db.cursor()

        cursor.execute(
            "SELECT id FROM users WHERE username = %s;",
            (username,)
        )
        if cursor.rowcount == 0:
            print(f"no such user {username}")
            cursor.close()
            return
        user_id = cursor.fetchone()[0]

        cursor.execute(
            "SELECT id, unit_id, prediction_start FROM time_series WHERE name = %s AND user_id = %s;",
            (name, user_id)
        )
        series_id, unit, prediction_start = cursor.fetchone()

        cursor.execute(
            "SELECT ts, value FROM records WHERE series_id = %s;",
            (series_id,)
        )
        return unit, prediction_start, cursor.fetchall()
