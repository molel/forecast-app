import argparse
import multiprocessing

from entities import MakePredictRequestEntity
import psycopg2

import pmdarima as pm


class PredictHandler(object):
    def __init__(self, config: argparse.Namespace, queue: multiprocessing.Queue):
        self.db = psycopg2.connect(
            host=config.db_address.split(':')[0],
            database=config.db_name,
            user=config.db_user,
            password=config.db_password,
            port=config.db_address.split(':')[1]
        )
        self.queue = queue
        self.make_prediction()

    def make_prediction(self):
        while True:
            request: MakePredictRequestEntity = self.queue.get(block=True)

            arima_model = pm.auto_arima(
                [item.value for item in request.items],
                start_p=1, start_q=1,
                max_p=5, max_q=5,

                seasonal=True, m=request.period,
                start_P=0, start_Q=0,
                max_P=2, max_Q=2,

                max_D=2, max_d=2,
                alpha=0.05,
                test='kpss',
                seasonal_test='ocsb',

                trace=True,
                error_action='ignore',
                suppress_warnings=True,
                stepwise=True,
                n_fits=10,
                information_criterion='bic',
                out_of_sample_size=7
            )

            x_pred, _ = arima_model.predict(
                n_periods=request.predict_periods,
                return_conf_int=True,
                alpha=0.05
            )

            self.save_prediction(
                username=request.username,
                unit=request.unit,
                name=request.name,
                period=request.period,
                predict_start=len(request.items),
                items=[]
            )

    def save_prediction(
            self,
            username: str,
            unit: str,
            name: str,
            period: int,
            predict_start: int,
            items: list[tuple[int, float]]
    ):

        cursor = self.db.cursor()

        # get user id
        cursor.execute(
            "SELECT id FROM users WHERE username = %s;",
            (username,)
        )
        if cursor.rowcount == 0:
            print(f"no such user {username}")
            cursor.close()
            return
        user_id = cursor.fetchone()[0]

        # get unit id
        cursor.execute(
            "SELECT id FROM measurement_units WHERE name = %s;",
            (unit,)
        )
        if cursor.rowcount == 0:
            cursor.execute(
                "INSERT INTO measurement_units (name) VALUES (%s) RETURNING id;",
                (unit,)
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
            (user_id, unit_id, name, period, predict_start)
        )
        self.db.commit()
        if cursor.rowcount == 0:
            print("error occurred during inserting into time series")
            cursor.close()
            return
        time_series_id = cursor.fetchone()[0]

        cursor.executemany(
            "INSERT INTO records (series_id, ts, value) VALUES (%s, %s, %s);",
            [(time_series_id, item[0], item[1]) for item in items]
        )
        self.db.commit()
        if cursor.rowcount == 0:
            print("error occurred during inserting time series items")
            cursor.close()
            return

        cursor.close()


class GetHandler(object):
    def __init__(self, config: argparse.Namespace):
        self.db = psycopg2.connect(
            host=config.db_address.split(':')[0],
            database=config.db_name,
            user=config.db_user,
            password=config.db_password,
            port=config.db_address.split(':')[1]
        )

    def get_prediction(self, username: str, name: str) -> (str, int, int, list[tuple[int, float]]):
        cursor = self.db.cursor()

        cursor.execute(
            "SELECT id FROM users WHERE username = %s;",
            (username,)
        )
        if cursor.rowcount == 0:
            cursor.close()
            raise Exception(f"no such user {username=}")
        user_id = cursor.fetchone()[0]

        cursor.execute(
            "SELECT id, unit_id, period, prediction_start FROM time_series WHERE name = %s AND user_id = %s;",
            (name, user_id)
        )
        if cursor.rowcount == 0:
            cursor.close()
            raise Exception(f"no such time-series {name=}")
        series_id, unit_id, period, prediction_start = cursor.fetchone()

        cursor.execute(
            "SELECT name FROM measurement_units WHERE id = %s;",
            (unit_id,)
        )
        if cursor.rowcount == 0:
            cursor.close()
            raise Exception(f"no such unit {unit_id}")
        unit = cursor.fetchone()[0]

        cursor.execute(
            "SELECT ts, value FROM records WHERE series_id = %s;",
            (series_id,)
        )
        if cursor.rowcount == 0:
            cursor.close()
            raise Exception(f"no such time-series records {name=}")
        return unit, prediction_start, period, cursor.fetchall()
