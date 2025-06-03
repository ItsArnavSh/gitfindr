from pydantic import BaseModel

class RegisterRequest(BaseModel):
    fullname: str
