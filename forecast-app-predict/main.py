import asyncio

import config
import server

if __name__ == '__main__':
    config = config.parse_config()
    asyncio.run(server.serve(config))
