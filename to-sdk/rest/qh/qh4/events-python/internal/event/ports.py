from abc import ABC, abstractmethod
from typing import Optional
from api.event.event import Event

class Repository(ABC):
    @abstractmethod
    def save(self, event: Event) -> None:
        pass

    @abstractmethod
    def find_by_name(self, name: str) -> Optional[Event]:
        pass
