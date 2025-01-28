SELECT SUM(volte) AS sumVolte, nome, Carta.elisir, Carta.danni
FROM GiocataIn
JOIN Carta ON carta = Carta.id
WHERE partita IN (
    SELECT Partita.id
    FROM Partita
    WHERE Partita.evento IN (
        SELECT Evento.id
        FROM Evento
        WHERE Evento.campionato = $1
    )
)
GROUP BY Carta.nome, Carta.elisir, Carta.danni
ORDER BY sumVolte DESC
LIMIT 1;
