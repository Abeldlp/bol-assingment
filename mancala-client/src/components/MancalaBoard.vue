<script setup lang="ts">
import Pit from "./Pit.vue";
import { onMounted } from "vue";
import { getGame } from "../api/MancalaApi";
import { mancalaGame, currentPlayer } from "../data/mancalaGame";

const mountGame = async () => {
  const res = await getGame("4");
  mancalaGame.value = res.data;
};

onMounted(async () => {
  mountGame();
});
</script>

<template>
  <h2>Now moves: {{ currentPlayer }}</h2>

  <div class="mancala-board">
    <div class="bucket">
      <h2>Player2 Bucket</h2>
      <Pit :stones="mancalaGame?.Player2.bucket" :index="0" />
    </div>
    <div class="pit-row opponent">
      <Pit
        v-for="(pit, index) in mancalaGame?.Player1.holes"
        :stones="pit"
        :index="index"
        :key="pit"
      />
    </div>
    <br />
    <br />
    <div class="pit-row">
      <Pit
        v-for="(pit, index) in mancalaGame?.Player2.holes"
        :stones="pit"
        :index="index"
        :key="pit"
      />
    </div>
    <div class="bucket">
      <h2>Player1 Bucket</h2>
      <Pit :stones="mancalaGame?.Player1.bucket" :index="0" />
    </div>
  </div>
</template>

<style scoped>
.manca-board {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.opponent {
  flex-direction: row-reverse;
}

.pit-row {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 10em;
}
</style>
