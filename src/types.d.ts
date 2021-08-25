interface Meeting {
  id: string;
  date: string;
  note: string;
}

interface Contact {
  id: string;
  name: string;
  location: string;
  now: string;
  background: string;
  meetings: Record<string, Meeting>;
}

interface State {
  contacts: Record<string, Contact>;
}
