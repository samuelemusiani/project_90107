SELECT Carta.nome, SUM(COALESCE(GiocataIn.volte, 0)) as volte
FROM Carta
LEFT JOIN GiocataIn ON Carta.id = GiocataIn.carta 
WHERE GiocataIn.partita IN (
    SELECT Partita.id FROM Partita WHERE Partita.evento = $1
)
GROUP BY Carta.nome
ORDER BY volte DESC
