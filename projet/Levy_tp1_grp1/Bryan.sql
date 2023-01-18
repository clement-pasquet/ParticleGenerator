-- I1A18B monmdpactuel44

CREATE TABLE ville (
	nomville VARCHAR(50) PRIMARY KEY NOT NULL
);

CREATE TABLE conducteur (
	numagr     	VARCHAR(9) PRIMARY KEY NOT NULL,
	dateagr    	DATE NOT NULL,
	nom        	VARCHAR(20) NOT NULL,
	villeadresse   VARCHAR(50) NOT NULL,
	adresse    	VARCHAR(50) NOT NULL,
	FOREIGN KEY ( villeadresse ) REFERENCES ville ( nomville )
);

CREATE TABLE vehicule (
	numero              	VARCHAR(7) PRIMARY KEY,
	type                	VARCHAR(20) NOT NULL,
	marque              	VARCHAR(20) NOT NULL,
	energie             	VARCHAR(20) NOT NULL,
	nbplaces            	NUMERIC(2) NOT NULL,
	anneecircu          	NUMERIC(4) NOT NULL,
	numagrementconducteur   VARCHAR(9),
	FOREIGN KEY ( numagrementconducteur )
    	REFERENCES conducteur ( numagr )
);

CREATE TABLE trajet (
	numero       	INTEGER NOT NULL,
	heuredep     	TIMESTAMP NOT NULL,
	heuredes     	TIMESTAMP NOT NULL,
	longueurtraj 	NUMERIC(4) NOT NULL,
	prix         	NUMERIC(3),
	numerovehicule   VARCHAR(7),
	nomvilledepart   VARCHAR(50) NOT NULL,
	nomvilledest 	VARCHAR(50) NOT NULL,
	PRIMARY KEY ( numero ),
	FOREIGN KEY ( nomvilledepart )REFERENCES ville ( nomville ),
	FOREIGN KEY ( nomvilledest )REFERENCES ville ( nomville ),
	FOREIGN KEY ( numerovehicule )REFERENCES vehicule ( numero )
 
);


create table PASSAGER(
    nomUtilisateur VARCHAR(20) PRIMARY KEY NOT NULL
);

CREATE TABLE NOTATION(
	nomutilisateur VARCHAR(20),
	numero INTEGER,
	avis VARCHAR(300),
	note NUMBER(1) NOT NULL,
	FOREIGN KEY ( nomutilisateur )REFERENCES passager ( nomutilisateur ),
	FOREIGN KEY ( numero )REFERENCES trajet ( numero )
);

INSERT INTO  VILLE VALUES('Angers');
INSERT INTO  VILLE VALUES('Saint-Nazaire');
INSERT INTO  VILLE VALUES('Nantes');
INSERT INTO  VILLE VALUES('Toulouse');
INSERT INTO  VILLE VALUES('Paris');
INSERT INTO  VILLE VALUES('Metz');
INSERT INTO  VILLE VALUES('Marseille');
INSERT INTO  VILLE VALUES('Rennes');
INSERT INTO  VILLE VALUES('Tours');
INSERT INTO  VILLE VALUES('Limoges');
INSERT INTO  VILLE VALUES('Annecy');
INSERT INTO  VILLE VALUES('Ivry-Sur-Seine');
INSERT INTO  VILLE VALUES('Mauves-Sur-Loire');
insert into CONDUCTEUR values('757505603', TO_DATE('13-07-2004', 'DD-MM-YYYY'), 'Basma', 'Paris', 'Champ de Mars, 5 Av. Anatole');
insert into conducteur values('444410902',TO_DATE('20-12-2015', 'DD-MM-YYYY'), 'Raphael', 'Nantes', '3 rue Maréchal Joffres');
insert into conducteur values('696938785', TO_DATE('09-08-2022', 'DD-MM-YYYY'), 'Justine', 'Paris', '17 Rue de France');
insert into conducteur values('444418462', TO_DATE('31-10-2010', 'DD-MM-YYYY'), 'Bryan', 'Saint-Nazaire', '10 Av. Pierre de Coubertin');
insert into conducteur values('854900751', TO_DATE('25-02-2016', 'DD-MM-YYYY'), 'Rachelle', 'Angers', '4 Bd de Lavoisier');
insert into conducteur values('414101856', TO_DATE('14-11-1986', 'DD-MM-YYYY'), 'Marine', 'Ivry-Sur-Seine', '15 Rue de la Chocolaterie');
insert into conducteur values('565626003', TO_DATE('14-05-2007', 'DD-MM-YYYY'), 'Baptiste', 'Mauves-Sur-Loire', '8 Rue Michel de Montaigne');
insert into conducteur values('373710051', TO_DATE('02-07-2013', 'DD-MM-YYYY'), 'Julie', 'Tours', '29 Rue du Pont Volant');
insert into conducteur values('151500089', TO_DATE('22-04-2001', 'DD-MM-YYYY'), 'Floran', 'Nantes', '100 Rue de l Égalité');
insert into conducteur values('747494065', TO_DATE('25-01-2004', 'DD-MM-YYYY'), 'Clement','Annecy', '9 Rue de l Arc en Ciel');

insert into VEHICULE values('AF894FG', 'Voiture', 'Audi', 'Diesel',5 , 1999,'151500089');
insert into vehicule values('KP754JF', 'Voiture', 'Hyundai', 'Diesel', 5, 2008,'565626003');
insert into vehicule values('XR761TS', 'Fourgon', 'Toyota', 'Hybride', 3, 2016,'444410902');
insert into vehicule values('KC847DS', 'Voiture', 'Volkswagen', 'Essence', 7, 2020,'373710051');
insert into vehicule values('KG123BL', 'SUV', 'BMW', 'Electrique', 5, 2018,'854900751');
insert into vehicule values('FR483TR', 'Voiture', 'Honda', 'Diesel', 5, 2010,'414101856');
insert into vehicule values('BL352LM', 'Voiture', 'Tesla', 'Electrique', 5, 2020,'444418462');
insert into vehicule values('DX891AZ', 'Fourgon', 'Kia', 'Diesel', 3, 1999,'757505603');
insert into vehicule values('EA777AL', 'Micro-Citadine', 'Mini', 'Electrique', 5, 2001,'747494065');
insert into vehicule values('HY912JK', 'Fourgon', 'Nissan', 'Diesel', 5, 2016,'696938785');


insert into PASSAGER values('Raph_01');
insert into PASSAGER values('Julo');
insert into PASSAGER values('Bini');
insert into PASSAGER values('r.dupont');
insert into PASSAGER values('Rea-Renarde');
insert into PASSAGER values('Screeny');
insert into PASSAGER values('Otyla');
insert into PASSAGER values('Basma.01');
insert into PASSAGER values('Floran');
insert into PASSAGER values('Juliaa');





INSERT INTO  TRAJET VALUES(0,TO_TIMESTAMP('15-05-2022 16:20','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('15-05-2022 17:50','DD-MM-YYYY HH24:MI'),111,11,'XR761TS','Nantes','Rennes');
insert into TRAJET VALUES(1, TO_TIMESTAMP('28-03-2019 14:50','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('28-03-2019 15:57','DD-MM-YYYY HH24:MI'),90,9,'FR483TR','Angers','Nantes');
insert into TRAJET VALUES(2, TO_TIMESTAMP('17-11-2022 08:30','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('17-11-2022 13:59','DD-MM-YYYY HH24:MI'),524,52,'DX891AZ','Toulouse','Tours');

insert into TRAJET VALUES(3, TO_TIMESTAMP('24-08-2017 10:00','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('24-08-2017 14:45','DD-MM-YYYY HH24:MI'),439,43,'EA777AL','Saint-Nazaire','Paris');
insert into TRAJET VALUES(4, TO_TIMESTAMP('20-06-2022 14:05','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('20-06-2022 18:12','DD-MM-YYYY HH24:MI'),404,40,'KP754JF','Marseille','Toulouse');
insert into TRAJET VALUES(5, TO_TIMESTAMP('06-01-2023 11:00','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('06-01-2023 16:11','DD-MM-YYYY HH24:MI'),580,58,'KG123BL','Tours','Metz');
insert into TRAJET VALUES(6, TO_TIMESTAMP('29-07-2022 16:10','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('29-07-2022 19:34','DD-MM-YYYY HH24:MI'),331,33,'BL352LM','Metz','Paris');
insert into TRAJET VALUES(7, TO_TIMESTAMP('13-06-2018 17:20','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('13-06-2018 19:03','DD-MM-YYYY HH24:MI'),129,13,'AF894FG','Rennes','Angers');
insert into TRAJET VALUES(8, TO_TIMESTAMP('06-03-2019 12:30','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('06-03-2019  15:04','DD-MM-YYYY HH24:MI'),240,24,'HY912JK','Paris','Tours');
insert into TRAJET VALUES(9, TO_TIMESTAMP('25-10-2022 11:00','DD-MM-YYYY HH24:MI'),TO_TIMESTAMP('25-10-2022 14:51','DD-MM-YYYY HH24:MI'),320,32,'KC847DS','Limoges','Nantes');



insert into NOTATION values('Raph_01', 6, 'Merci pour ce beau trajet', 5);
insert into NOTATION values('Julo', 2, 'Oui', 5);
insert into NOTATION values('Bini', 5, 'Parfait ! Est ce possible de faire abstraction de l insipide autoroute et d arriver, à bon port, en toute sécurité par une conduite prudente, agréable et souple? Absolument! Laissez-vous simplement embarquer par Morgane; si prévoyante,attentionnée et sereine...', 4);
insert into NOTATION values('Rea-Renarde', 1, 'Excellent ! Bonne conduite et très sympa. Je recommande !!', 4);
insert into NOTATION values('Basma.01', 9, 'Super trajet', 3);
insert into NOTATION values('Otyla',3, 'Excellent ! J ai effectué un très bon voyage. Plein de sujets de discussion en commun. Je vous recommande Michel, vous ne serez pas déçu. Conduite très agréable. J espère un jour pouvoir re voyager avec Michel. Yves !', 5);
insert into NOTATION values('r.dupont', 7, 'Trajet très silencieux, mais arrivé à temps', 3);
insert into NOTATION values('Floran', 4, 'Excellent ! Passager ponctuel et sympathique, le trajet avec Julien s est très bien passé. Discussion agréable et très intéressante sur plusieurs sujets., nous avons bien échangé durant le voyage. Je recommande vivement.', 4);
insert into NOTATION values('Screeny', 8, 'PARFAIT, Voyage trop court. Julien et sa chérie sont super, de vrais chouchou tous les 2. A bientôt pour un prochain voyage ou apéro', 1);
insert into NOTATION values('Juliaa', 0, 'Très bon trajet !', 2);