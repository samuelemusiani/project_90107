CREATE TRIGGER trigger_check_max_carte
BEFORE INSERT ON Formato
FOR EACH ROW
EXECUTE FUNCTION check_max_carte();

CREATE TRIGGER trigger_check_players_number
BEFORE INSERT ON Gioca
FOR EACH ROW
EXECUTE FUNCTION check_players_number();

CREATE TRIGGER trigger_aggiorna_team_stats
AFTER INSERT OR DELETE OR UPDATE ON Ingaggio
FOR EACH ROW
EXECUTE FUNCTION aggiorna_team_stats();

CREATE TRIGGER trigger_check_used_cards
BEFORE INSERT OR UPDATE ON GiocataIn
FOR EACH ROW
EXECUTE FUNCTION check_used_cards();

