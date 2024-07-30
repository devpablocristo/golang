from datetime import date

class Event:
    def __init__(self, name: str, location: str, event_date: date, description: str):
        self.name = name
        self.location = location
        self.event_date = event_date
        self.description = description
        self.attendees = []

    def add_attendee(self, attendee: str):
        self.attendees.append(attendee)

    def remove_attendee(self, attendee: str):
        if attendee in self.attendees:
            self.attendees.remove(attendee)

    def get_attendees(self) -> list:
        return self.attendees

    def __str__(self) -> str:
        return (f"Event: {self.name}\nLocation: {self.location}\nDate: {self.event_date}\n"
                f"Description: {self.description}\nAttendees: {', '.join(self.attendees)}")
