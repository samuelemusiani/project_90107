SELECT SUM(I.elisir_usato) AS "Elisir Usato", SUM(I.elisir_sprecato) AS "Elisir Sprecato", 
    SUM(I.danni_fatti) as "Danni Fatti", SUM(T.danni_fatti) AS danni_subiti
FROM Gioca I, Gioca T
WHERE I.giocatore = $1 AND I.partita = T.partita
     AND T.giocatore != I.giocatore AND I.partita = $2
