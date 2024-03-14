# Fizzbuzzlbc

Développement d'un serveur web HTTP sans paquet open source.

---

## Description

Le projet est un simple serveur web http contenant, actuellement, une seule route.

| Methode | Path      | Header                         |
|---------|-----------|--------------------------------|
| POST    | /fizzbuzz | Content-Type: application/json |

Le body de la requête se compose de cette façon :
```json
{
    "limit": 100,
    "int1": 3,
    "int2": 5,
    "str1": "fizz",
    "str2": "buzz"
}
```

Le but de cette route, est d'exécuter un faiteau, qui consiste à écrire tous les nombres de 1 à `limit`, en remplaçant
tous les multiples de `int1` par `str1`, tous les multiples de `int2` par `str2` et tous les multiples de `int1` et
`int2` par `str1`+`str2`

le résultat de la requête est retourné en JSON de cette façon :
```json
{
  "result": [
    "1",
    "2",
    "fizz",
    "4",
    "buzz",
    "fizz",
    "7",
    "8",
    "fizz",
    "buzz",
    "11",
    "fizz",
    "13",
    "14",
    "fizzbuzz"
  ]
}
```

Si une erreur survient, la réponse aura cette forme :
```json
{
  "error": "erreur retournée"
}
```

une collection postman est disponible à la racine du projet -> `fizzbuzz_collection`

## Lancement du projet

Le projet se lance de la façon suivante :
```shell
# directement à la racine du projet
go run main.go

# Via docker
make docker-run
```

## Test

Les tests s'exécutent de la façon suivante :
```shell
# executer les tests simples
make test

# executer les tests avec le coverage
make coverage
```

La commande `make coverage` génère deux fichiers de sortie, dont le fichier `coverage.html`, qui permettent de voir
en détails dans son navigateur le coverage de chaque fichier.

## Dépendance

Il n'y a aucune dépendance sur le projet.
