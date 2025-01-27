CREATE OR REPLACE FUNCTION check_max_carte()
RETURNS TRIGGER AS $$
BEGIN
  IF (SELECT COUNT(*) FROM Formato WHERE mazzo = NEW.mazzo) >= 8 THEN
    RAISE EXCEPTION 'Un mazzo può contenere al massimo 8 carte.';
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_check_max_carte
BEFORE INSERT ON Formato
FOR EACH ROW
EXECUTE FUNCTION check_max_carte();


CREATE OR REPLACE FUNCTION check_players_number()
RETURNS TRIGGER AS $$
BEGIN
  IF (SELECT COUNT(*) FROM Gioca WHERE partita = NEW.partita) >= 2 THEN
    RAISE EXCEPTION 'Una partita può contenere al massimo 2 giocatori.';
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_check_players_number
BEFORE INSERT ON Gioca
FOR EACH ROW
EXECUTE FUNCTION check_players_number();

CREATE OR REPLACE FUNCTION aggiorna_team_stats()
RETURNS TRIGGER AS $$
DECLARE
  giocatore_piu_giovane INT;
  giocatore_piu_vecchio INT;
  eta_media_team REAL;
BEGIN
  -- Il giocatore piu giovane
  SELECT i.giocatore
  INTO giocatore_piu_giovane
  FROM Ingaggio i
  JOIN Giocatore g ON i.giocatore = g.id
  WHERE i.team = NEW.team
  ORDER BY g.dig DESC
  LIMIT 1;

  -- Il giocatore piu vecchio
  SELECT i.giocatore
  INTO giocatore_piu_vecchio
  FROM Ingaggio i
  JOIN Giocatore g ON i.giocatore = g.id
  WHERE i.team = NEW.team
  ORDER BY g.dig ASC
  LIMIT 1;

  -- Eta media
  SELECT AVG(DATE_PART('year', AGE(g.dig)))::REAL
  INTO eta_media_team
  FROM Ingaggio i
  JOIN Giocatore g ON i.giocatore = g.id
  WHERE i.team = NEW.team;

  UPDATE Team
  SET gpg = giocatore_piu_giovane,
      gpv = giocatore_piu_vecchio,
      eta_media = COALESCE(eta_media_team, 0)
  WHERE id = NEW.team;

  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_aggiorna_team_stats
AFTER INSERT OR DELETE OR UPDATE ON Ingaggio
FOR EACH ROW
EXECUTE FUNCTION aggiorna_team_stats();
