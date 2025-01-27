CREATE TABLE IF NOT EXISTS Persona(
  id SERIAL PRIMARY KEY,
  nome TEXT NOT NULL,
  cognome TEXT NOT NULL,
  data_nascita DATE NOT NULL,
  luogo_nascita TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Giocatore(
  id SERIAL PRIMARY KEY,
  username TEXT NOT NULL,
  dig DATE NOT NULL,
  FOREIGN KEY (id) REFERENCES Persona(id)
);

CREATE TABLE IF NOT EXISTS Team(
  id SERIAL PRIMARY KEY,
  nome TEXT NOT NULL,
  logo BYTEA,
  data_fondazione DATE NOT NULL,
  stato_geografico TEXT NOT NULL,
  gpg INT NOT NULL,
  gpv INT NOT NULL,
  eta_media REAL NOT NULL,
  FOREIGN KEY (gpg) REFERENCES Giocatore(id),
  FOREIGN KEY (gpv) REFERENCES Giocatore(id)
);

CREATE TABLE IF NOT EXISTS Sponsor(
  id SERIAL PRIMARY KEY,
  nome TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Campionato(
  id SERIAL PRIMARY KEY,
  nome TEXT NOT NULL,
  luogo TEXT NOT NULL,
  data_inizio DATE NOT NULL,
  data_fine DATE NOT NULL,
  tipo TEXT NOT NULL,
  montepremi INT NOT NULL,
);

CREATE TABLE IF NOT EXISTS LeaderboardT(
  id SERIAL PRIMARY KEY,
  campionato INT NOT NULL,
  FOREIGN KEY (campionato) REFERENCES Campionato(id)
);

CREATE TABLE IF NOT EXISTS LeaderboardG(
  id SERIAL PRIMARY KEY,
  campionato INT NOT NULL,
  FOREIGN KEY (campionato) REFERENCES Campionato(id)
);

CREATE TABLE IF NOT EXISTS Evento(
  id SERIAL PRIMARY KEY,
  nome TEXT NOT NULL,
  luogo TEXT NOT NULL,
  data DATE NOT NULL,
  posti_totali INT NOT NULL,
  posti_usati INT NOT NULL,
  campionato INT NOT NULL,
  FOREIGN KEY (campionato) REFERENCES Campionato(id)
);

CREATE TABLE IF NOT EXISTS Biglietto(
  id SERIAL PRIMARY KEY,
  prezzo REAL NOT NULL,
  posto INT NOT NULL,
);

CREATE TABLE IF NOT EXISTS Partita(
  id SERIAL PRIMARY KEY,
  vincitore INT NOT NULL,
  orario TIME NOT NULL,
  tempo INT NOT NULL,
  evento INT NOT NULL,
  FOREIGN KEY (vincitore) REFERENCES Giocatore(id),
  FOREIGN KEY (evento) REFERENCES Evento(id)
);

CREATE TABLE IF NOT EXISTS Mazzo(
  id SERIAL PRIMARY KEY,
);

CREATE TABLE IF NOT EXISTS Carta(
  id SERIAL PRIMARY KEY,
  nome TEXT NOT NULL,
  elisir INT NOT NULL,
  danni INT NOT NULL,
);

CREATE TABLE IF NOT EXISTS Ingaggio(
  team INT NOT NULL,
  giocatore INT NOT NULL,
  salario INT NOT NULL,
  PRIMARY KEY (team, giocatore),
  FOREIGN KEY (team) REFERENCES Team(id),
  FOREIGN KEY (giocatore) REFERENCES Giocatore(id)
);

CREATE TABLE IF NOT EXISTS Allena(
  persona INT NOT NULL,
  team INT NOT NULL,
  salario INT NOT NULL,
  PRIMARY KEY (persona, team),
  FOREIGN KEY (persona) REFERENCES Persona(id),
  FOREIGN KEY (team) REFERENCES Team(id)
);

CREATE TABLE IF NOT EXISTS Sponsorizza(
  sponsor INT NOT NULL,
  team INT NOT NULL,
  budget INT NOT NULL,
  PRIMARY KEY (sponsor, team),
  FOREIGN KEY (sponsor) REFERENCES Sponsor(id),
  FOREIGN KEY (team) REFERENCES Team(id)
);

CREATE TABLE IF NOT EXISTS Assiste(
  persona INT NOT NULL,
  evento INT NOT NULL,
  biglietto INT NOT NULL,
  PRIMARY KEY (persona, evento),
  FOREIGN KEY (persona) REFERENCES Persona(id),
  FOREIGN KEY (evento) REFERENCES Evento(id),
);

CREATE TABLE IF NOT EXISTS Gioca(
  giocatore INT NOT NULL,
  partita INT NOT NULL,
  mazzo INT NOT NULL,
  elisir_usato INT NOT NULL,
  elisir_sprecato INT NOT NULL,
  danni_fatti INT NOT NULL,
  tipo_torri TEXT NOT NULL,
  PRIMARY KEY (giocatore, partita),
  FOREIGN KEY (giocatore) REFERENCES Giocatore(id),
  FOREIGN KEY (partita) REFERENCES Partita(id),
  FOREIGN KEY (mazzo) REFERENCES Mazzo(id)
);

CREATE TABLE IF NOT EXISTS Commenta(
  persona INT NOT NULL,
  partita INT NOT NULL,
  lingua TEXT NOT NULL,
  PRIMARY KEY (persona, partita),
  FOREIGN KEY (persona) REFERENCES Persona(id),
  FOREIGN KEY (partita) REFERENCES Partita(id)
);

CREATE TABLE IF NOT EXISTS Formato(
  mazzo INT NOT NULL,
  carta INT NOT NULL,
  PRIMARY KEY (mazzo, carta),
  FOREIGN KEY (mazzo) REFERENCES Mazzo(id),
  FOREIGN KEY (carta) REFERENCES Carta(id)
);

CREATE TABLE IF NOT EXISTS FormataT(
  campionato INT NOT NULL,
  team INT NOT NULL,
  posizione INT NOT NULL,
  PRIMARY KEY (campionato, team),
  FOREIGN KEY (campionato) REFERENCES Campionato(id),
  FOREIGN KEY (team) REFERENCES Team(id)
);

CREATE TABLE IF NOT EXISTS FormataG(
  campionato INT NOT NULL,
  giocatore INT NOT NULL,
  posizione INT NOT NULL,
  PRIMARY KEY (campionato, giocatore),
  FOREIGN KEY (campionato) REFERENCES Campionato(id),
  FOREIGN KEY (giocatore) REFERENCES Giocatore(id)
);
