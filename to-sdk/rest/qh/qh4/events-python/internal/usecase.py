from typing import Optional
from datetime import date
from api.event.event import Event
from api.event.ports import Repository

class UseCase:
    def __init__(self, event_repository: Repository):
        self.event_repository = event_repository

    def create_event(self, name: str, location: str, event_date: date, description: str) -> Event:
        event = Event(name, location, event_date, description)
        self.event_repository.save(event)
        return event

    def get_event(self, name: str) -> Optional[Event]:
        return self.event_repository.find_by_name(name)

    def add_attendee(self, event_name: str, attendee: str) -> Optional[Event]:
        event = self.event_repository.find_by_name(event_name)
        if event:
            event.add_attendee(attendee)
            self.event_repository.save(event)
        return event

    def remove_attendee(self, event_name: str, attendee: str) -> Optional[Event]:
        event = self.event_repository.find_by_name(event_name)
        if event:
            event.remove_attendee(attendee)
            self.event_repository.save(event)
        return event
