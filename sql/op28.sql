SELECT COUNT(*) AS venduti
FROM Assiste
WHERE Assiste.evento IN (
    SELECT id FROM Evento WHERE Evento.campionato = $1
)
