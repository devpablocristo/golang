from fastapi import APIRouter, Depends
from api.usecase import UseCase
from api.event.schema import EventCreateSchema, EventSchema
from typing import Optional
from datetime import date

router = APIRouter()

# Dependency
def get_event_service() -> UseCase:
    from api.event.repository_adapter import InMemoryRepository
    repository = InMemoryRepository()
    return UseCase(repository)

@router.post("/events/", response_model=EventSchema)
async def create_event(
    event: EventCreateSchema,
    service: UseCase = Depends(get_event_service)
):
    created_event = service.create_event(
        name=event.name,
        location=event.location,
        event_date=event.event_date,
        description=event.description
    )
    return EventSchema(
        name=created_event.name,
        location=created_event.location,
        event_date=created_event.event_date,
        description=created_event.description,
        attendees=created_event.get_attendees()
    )

@router.get("/events/{name}", response_model=Optional[EventSchema])
async def get_event(name: str, service: UseCase = Depends(get_event_service)):
    event = service.get_event(name)
    if event is None:
        return None
    return EventSchema(
        name=event.name,
        location=event.location,
        event_date=event.event_date,
        description=event.description,
        attendees=event.get_attendees()
    )

@router.post("/events/{event_name}/attendees/", response_model=Optional[EventSchema])
async def add_attendee(event_name: str, attendee: str, service: UseCase = Depends(get_event_service)):
    event = service.add_attendee(event_name, attendee)
    if event is None:
        return None
    return EventSchema(
        name=event.name,
        location=event.location,
        event_date=event.event_date,
        description=event.description,
        attendees=event.get_attendees()
    )

@router.delete("/events/{event_name}/attendees/", response_model=Optional[EventSchema])
async def remove_attendee(event_name: str, attendee: str, service: UseCase = Depends(get_event_service)):
    event = service.remove_attendee(event_name, attendee)
    if event is None:
        return None
    return EventSchema(
        name=event.name,
        location=event.location,
        event_date=event.event_date,
        description=event.description,
        attendees=event.get_attendees()
    )
