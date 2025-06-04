build:
	uv sync
run:
	uv run alembic revision --autogenerate -m "migrations"
	uv run alembic upgrade head
	uv run uvicorn main:app --reload
