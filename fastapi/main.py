from transformers import AutoTokenizer, AutoModelForSeq2SeqLM
from fastapi.middleware.cors import CORSMiddleware
from typing import Union

from fastapi import FastAPI
from pydantic import BaseModel
app = FastAPI()
origins = ["*"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

model_name = '/home/fanzruskripsi/skripsi/final-project-university/fastapi/model/t5/t5-small-finetuned-xlsum-concat-multi-news'
tokenizer = AutoTokenizer.from_pretrained(model_name)
model = AutoModelForSeq2SeqLM.from_pretrained(model_name)

class Req(BaseModel):
    text: str
    max_length: int

@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}

@app.post("/summarize/")
def summarize(req: Req):
    input_ids = tokenizer.encode(req.text, return_tensors="pt", add_special_tokens=True)
    generated_ids = model.generate(input_ids=input_ids, num_beams=2, max_length=req.max_length,  repetition_penalty=2.5, length_penalty=1.0, early_stopping=True) 
    preds = [tokenizer.decode(g, skip_special_tokens=True, clean_up_tokenization_spaces=True) for g in generated_ids]
    return {"result":preds[0]}
