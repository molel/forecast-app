import argparse
import os


class EnvDefault(argparse.Action):
    def __init__(self, env_var, required=True, default=None, **kwargs):
        if not default and env_var:
            if env_var[2:] if env_var.startswith('--') else env_var in os.environ:
                default = os.environ[env_var]
        if required and default:
            required = False
        super(EnvDefault, self).__init__(default=default, required=required,
                                         **kwargs)

    def __call__(self, parser, namespace, values, option_string=None):
        setattr(namespace, self.dest, values)


def parse_config() -> argparse.Namespace:
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument(
        '--http',
        action='store',
        default=8080,
        type=int,
        required=True,
        help='http server port',
    )
    arg_parser.add_argument(
        '--database-address',
        action='store',
        type=str,
        required=True,
        help='database address',
        dest='db_address'
    )
    arg_parser.add_argument(
        '--pool',
        action='store',
        type=int,
        required=True,
        help='number of processes for predict handler pool',
        dest='pool'
    )

    DB_USER_KEY = 'DB_USER'
    DB_PASSWORD_KEY = 'DB_PASSWORD'
    DB_NAME_KEY = 'DB_NAME'

    arg_parser.add_argument(
        '--' + DB_USER_KEY,
        nargs=1,
        action=EnvDefault,
        env_var=DB_USER_KEY,
        type=str,
        required=True,
        help='database username',
        dest=DB_USER_KEY.lower()
    )
    arg_parser.add_argument(
        '--' + DB_PASSWORD_KEY,
        nargs=1,
        action=EnvDefault,
        env_var=DB_PASSWORD_KEY,
        type=str,
        required=True,
        help='database user password',
        dest=DB_PASSWORD_KEY.lower()
    )
    arg_parser.add_argument(
        '--' + DB_NAME_KEY,
        nargs=1,
        action=EnvDefault,
        env_var=DB_NAME_KEY,
        type=str,
        required=True,
        help='database name',
        dest=DB_NAME_KEY.lower()
    )

    return arg_parser.parse_args()
