# 📦 Gestion de Livraison

## 🚀 **Description**
Gestion de Livraison est une application permettant de suivre et gérer les livraisons en utilisant différents moyens de transport tels que les **drones**, **camions** et **bateaux**. L'application assure le suivi en temps réel et permet d'optimiser la gestion logistique.

---

## 📌 **Fonctionnalités principales**

✅ Gestion des livraisons via **Drone, Camion et Bateau**.  
✅ Système de **tracking** des livraisons en temps réel.  
✅ Simulation de **livraison avec logs et statut des transports**.  
✅ Gestion de la capacité de charge des camions et de la batterie des drones.  
✅ Logs des statuts des livraisons avec lecture et effacement automatique.  
✅ Interface interactive pour choisir le moyen de transport.  

---

## 🛠 **Stack Technique**

- **Langage** : Go
- **Framework Web** : Gin
- **Gestion des tests** : `testing`, `testify`
- **Gestion des logs** : Fichiers `.log`

---

## 📂 **Structure du projet**

```
📦 gestion-de-livraison
├── cmd/
│   ├── api/
│   │   ├── main.go          # Fichier principal
│   │   ├── main_test.go     # Tests du système de tracking
│
├── pkg/
│   ├── factory/
│   │   ├── factory.go       # Factory pour créer les moyens de transport
│   │   ├── factory_test.go  # Tests de la factory
│   ├── interfaces/
│   │   ├── interface.go     # Interface TransportMethod
│   │   ├── interface_test.go # Tests des interfaces
│   ├── models/
│   │   ├── drone.go         # Modèle Drone
│   │   ├── boat.go          # Modèle Boat
│   │   ├── truck.go         # Modèle Truck
│   │   ├── drone_test.go    # Tests Drone
│   │   ├── boat_test.go     # Tests Boat
│   │   ├── truck_test.go    # Tests Truck
│   ├── utils/
│   │   ├── transport.go     # Fonctions utilitaires (logs, recharges...)
│   │   ├── transport_test.go # Tests des utilitaires
│
├── docs/
│   ├── tests.md             # Documentation des tests
│
├── status.log               # Fichier de log des statuts de livraison
└── README.md                # Documentation du projet
```

---

## 🧪 **Tests Unitaires**

Des tests unitaires couvrent les fonctionnalités principales :

| 📁 Module     | 🧪 Fichier de test               | 🔎 Fonctionnalités testées |
|--------------|--------------------------------|----------------------|
| **Models**   | `pkg/models/drone_test.go`     | ✅ Livraison et statut des drones |
|              | `pkg/models/boat_test.go`      | ✅ Livraison et statut des bateaux |
|              | `pkg/models/truck_test.go`     | ✅ Livraison et statut des camions |
| **Factory**  | `pkg/factory/factory_test.go`  | ✅ Création des moyens de transport |
| **Interfaces** | `pkg/interfaces/interface_test.go` | ✅ Vérification de l’implémentation de `TransportMethod` |
| **Utils**    | `pkg/utils/transport_test.go`  | ✅ Gestion des logs, recharge drone, chargement camions |
| **API / Main** | `cmd/api/main_test.go`        | ✅ Système de tracking et gestion des logs |

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

## 🚀 **Comment utiliser l'application ?**

1️⃣ **Cloner le projet**
```sh
git clone https://github.com/username/gestion-de-livraison.git
cd gestion-de-livraison
```

2️⃣ **Lancer l'application**
```sh
go run cmd/api/main.go
```

3️⃣ **Consulter les logs en direct**
```sh
tail -f status.log
```

---

## 👥 **Auteurs**

🚀 **Thierno Sadou Barry**  
🚀 **El Hadji Mamadou SARR**  
🚀 **Arthur Deumeni Tsako**
🚀 **Arold Ngouani Yapteu**
🚀 **Souleymane SALL**

---

🔥 **Les livraisons automatisées sont l'avenir, et ce projet en est une démonstration !** 🚀