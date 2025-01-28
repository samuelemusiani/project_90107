SELECT nome, cognome, data_nascita, luogo_nascita
FROM Persona
JOIN Commenta ON Persona.id = Commenta.persona
WHERE Commenta.partita = $1
