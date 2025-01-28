SELECT tipo_torri, COUNT(*) as countTorri
FROM Gioca
WHERE Gioca.partita IN (
    SELECT id FROM Partita WHERE Partita.evento IN (
        SELECT id FROM Evento WHERE Evento.campionato = $1
    )
)
GROUP BY tipo_torri
ORDER BY countTorri DESC
LIMIT 1
