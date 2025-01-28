SELECT COUNT(*)
FROM Assiste
WHERE Assiste.evento = $1
GROUP BY biglietto
