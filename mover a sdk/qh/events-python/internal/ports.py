from abc import ABC, abstractmethod
from typing import Optional
from datetime import date
from domain.event import Event

class UseCasePort(ABC):
    @abstractmethod
    def create_event(self, name: str, location: str, event_date: date, description: str) -> Event:
        pass

    @abstractmethod
    def get_event(self, name: str) -> Optional[Event]:
        pass

    @abstractmethod
    def add_attendee(self, event_name: str, attendee: str) -> Optional[Event]:
        pass

    @abstractmethod
    def remove_attendee(self, event_name: str, attendee: str) -> Optional[Event]:
        pass
