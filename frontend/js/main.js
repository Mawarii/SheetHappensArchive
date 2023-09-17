import { createApp } from 'https://cdn.jsdelivr.net/npm/vue@3.2.1/dist/vue.esm-browser.js';

// Erstellen Sie eine Vue-App
const app = createApp({
  data() {
    return {
      newCharacter: {
        name: "",
        race: "",
      },
      characters: [],
    };
  },
  methods: {
    createCharacter() {
      // Senden Sie eine POST-Anfrage zum Erstellen eines neuen Charakters
      fetch("/character", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(this.newCharacter),
      })
        .then((response) => response.json())
        .then((data) => {
          // Fügen Sie den Charakter zur Liste hinzu
          this.characters.push(data.data);

          // Formular zurücksetzen
          this.newCharacter.name = "";
          this.newCharacter.race = "";
        })
        .catch((error) => console.error("Fehler:", error));
    },
    getAllCharacters() {
      // Senden Sie eine GET-Anfrage zum Abrufen aller Charaktere
      fetch("/characters")
        .then((response) => response.json())
        .then((data) => {
        this.characters = data.data;
        })
        .catch((error) => console.error("Fehler:", error));
    },
    deleteCharacter(characterId) {
        // Senden Sie eine DELETE-Anfrage, um den Charakter zu löschen
        fetch(`/character/${characterId}`, {
            method: "DELETE",
        })
            .then(() => {
            // Charakter aus der Liste entfernen
            const characterIndex = this.characters.findIndex(
                (character) => character.id === characterId
            );
            if (characterIndex !== -1) {
                this.characters.splice(characterIndex, 1);
            }
            })
            .catch((error) => console.error("Fehler beim Löschen:", error));
    },
  },
});

// Mounten Sie die App auf das DOM
app.mount("#app");
