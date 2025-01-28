SELECT COUNT(mazzo) as volte, mazzo 
FROM Gioca 
JOIN Partita ON Partita.id = partita 
WHERE Partita IN (
    SELECT id FROM Partita WHERE Partita.evento IN (
        SELECT id FROM Evento WHERE Evento.campionato = $1
    )
) AND Partita.vincitore = giocatore 
GROUP BY mazzo 
ORDER BY volte DESC 
LIMIT 1;
