import argparse

import psycopg2


def connect_to_db(config: argparse.Namespace):
    try:
        with psycopg2.connect(
                host=config.db_address,
                database=config.db_name,
                user=config.db_user,
                password=config.db_password
        ) as conn:
            print('connected to the database')
            return conn
    except psycopg2.DatabaseError as error:
        print(error)
