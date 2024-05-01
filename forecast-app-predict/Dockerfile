FROM python:3.13-slim

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN adduser -D user

RUN chown user:user ./main.py

USER user

CMD python main.py --http {$HTTP} --database-address {$DATABASE_ADDRESS} --pool {$POOL}

