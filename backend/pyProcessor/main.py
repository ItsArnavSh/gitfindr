from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import List
from processFile import processRepo, processText

app = FastAPI()

# Define a model for the incoming request
class Item(BaseModel):
    name: str

class Text(BaseModel):
    text: str

# Define the model for the response data
class ResponseData(BaseModel):
    data_list: List[str]

@app.post("/convert", response_model=ResponseData)
async def convert_item(link: Item):
    try:
        listData = processRepo(link.name)
        if not isinstance(listData, list):
            raise HTTPException(status_code=500, detail="Invalid response from processRepo")
        return ResponseData(data_list=listData)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error processing request: {str(e)}")

@app.get("/")
async def read_root():
    return {"message": "Welcome to the service"}

@app.post("/convertText", response_model=ResponseData)
async def convert_text(data: Text):
    try:
        listData = processText(data.text)
        if not isinstance(listData, list):
            raise HTTPException(status_code=500, detail="Invalid response from processText")
        return ResponseData(data_list=listData)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error processing request: {str(e)}")
