SELECT posizione, team as "team o giocatore"
FROM FormataT
JOIN Team ON Team.id = team
WHERE campionato = $1

UNION ALL

SELECT FormataG.posizione, FormataG.giocatore
FROM FormataG
JOIN Giocatore ON Giocatore.id = giocatore
WHERE campionato = $1

ORDER BY posizione;
