# 🧪 Tests Unitaires - Gestion de Livraison

Ce projet inclut une série de tests unitaires pour garantir le bon fonctionnement des différents composants du système de gestion de livraison.

---

## 📂 Structure des Tests

| 📁 Module     | 🧪 Fichier de test               | 🔎 Fonctionnalités testées |
|--------------|--------------------------------|----------------------|
| **Models**   | `pkg/models/drone_test.go`     | ✅ Livraison et statut des drones |
|              | `pkg/models/boat_test.go`      | ✅ Livraison et statut des bateaux |
|              | `pkg/models/truck_test.go`     | ✅ Livraison et statut des camions |
| **Factory**  | `pkg/factory/factory_test.go`  | ✅ Création des moyens de transport |
| **Interfaces** | `pkg/interfaces/interface_test.go` | ✅ Vérification de l’implémentation de `TransportMethod` |
| **Utils**    | `pkg/utils/transport_test.go`  | ✅ Gestion des logs, recharge drone, chargement camions |
| **API / Main** | `cmd/api/main_test.go`        | ✅ Système de tracking et gestion des logs |

---

## 🚀 **Exécution des tests**

### 🔹 Exécuter **tous les tests**
```sh
go test -v ./...
```

### 🔹 Exécuter un test spécifique
```sh
go test -run TestGetTransportMethod ./pkg/factory
```

### 🔹 Vérifier la couverture des tests
```sh
go test -cover ./...
```

---

## ✅ **Bonnes pratiques des tests**
- Les tests sont écrits avec `testing` et `testify/assert` pour des assertions claires.
- Chaque module possède son fichier de test dans le **même dossier** (`*_test.go`).
- Les tests vérifient **les cas normaux et les erreurs possibles** (ex: batterie faible, dépassement de capacité, etc.).
- Le système de logs est testé pour s’assurer qu’il fonctionne correctement.


🔥 **Les tests unitaires garantissent un code fiable et robuste !** 🚀

