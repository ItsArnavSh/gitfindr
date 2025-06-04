from pydantic import BaseModel

class RegisterRequest(BaseModel):
    fullname: str
class QueryRequest(BaseModel):
    query: str
