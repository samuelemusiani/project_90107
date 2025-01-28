SELECT Team.nome, logo, data_fondazione, stato_geografico, gpg, gpv, eta_media, Persona.nome as coach, Sponsor.nome as sponsor
FROM Team
JOIN Allena ON Team.id = Allena.team
JOIN Persona ON Persona.id = Allena.persona
JOIN Sponsorizza ON Team.id = Sponsorizza.team
JOIN Sponsor ON Sponsorizza.sponsor = Sponsor.id
WHERE Team.id = $1
