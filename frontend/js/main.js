import { createApp } from 'https://cdn.jsdelivr.net/npm/vue@3.2.1/dist/vue.esm-browser.js';

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
      fetch("/main/character", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(this.newCharacter),
      })
        .then((response) => response.json())
        .then((data) => {
          this.characters.push(data.data);

          this.newCharacter.name = "";
          this.newCharacter.race = "";
        })
        .catch((error) => console.error("Fehler:", error));
    },
    getAllCharacters() {
      fetch("/main/character/list")
        .then((response) => response.json())
        .then((data) => {
        this.characters = data.data;
        })
        .catch((error) => console.error("Fehler:", error));
    },
    deleteCharacter(characterId) {
        fetch(`/main/character/${characterId}`, {
            method: "DELETE",
        })
            .then(() => {
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

app.mount("#app");
