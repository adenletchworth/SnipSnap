from fastapi import FastAPI, Request
from sentence_transformers import SentenceTransformer
from pydantic import BaseModel
import uvicorn

app = FastAPI()
model = SentenceTransformer("all-MiniLM-L6-v2")

class EmbedRequest(BaseModel):
    text: str

@app.post("/embed")
async def embed(request: EmbedRequest):
    vec = model.encode(request.text).tolist()
    return {"embedding": vec}

