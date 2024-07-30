from datetime import date
from pydantic import BaseModel, Field
from typing import List

class EventCreateSchema(BaseModel):
    name: str
    location: str
    event_date: date
    description: str

class EventSchema(BaseModel):
    name: str
    location: str
    event_date: date
    description: str
    attendees: List[str] = Field(default_factory=list)

    class Config:
        arbitrary_types_allowed = True
