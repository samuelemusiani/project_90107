SELECT Carta.nome, SUM(COALESCE(GiocataIn.volte, 0)) as volte
FROM Carta
LEFT JOIN GiocataIn ON Carta.id = GiocataIn.carta 
WHERE GiocataIn.partita = $1
GROUP BY Carta.nome
ORDER BY volte DESC
