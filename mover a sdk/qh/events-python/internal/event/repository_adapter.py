from typing import Optional
from api.event.event import Event
from api.event.ports import Repository

class InMemoryRepository(Repository):
    def __init__(self):
        self.events = {}

    def save(self, event: Event) -> None:
        self.events[event.name] = event

    def find_by_name(self, name: str) -> Optional[Event]:
        return self.events.get(name)
