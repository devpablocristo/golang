from fastapi import FastAPI
from cmd.handlers.rest.router import router

app = FastAPI()

app.include_router(router)
