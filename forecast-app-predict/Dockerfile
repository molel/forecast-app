FROM python:3.11

WORKDIR /app

COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY . .

RUN adduser --disabled-password newuser

RUN chown newuser:newuser ./main.py

USER newuser

CMD python main.py --http $HTTP --database-address $DATABASE_ADDRESS --pool $POOL

