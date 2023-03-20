<script setup lang="ts">
import Pit from "./Pit.vue";
import { computed } from "vue";
import { mancalaGame, currentPlayer } from "../data/mancalaGame";
import Bucket from "./Bucket.vue";

const opponentBoard = computed(() => {
  const reversedArray = mancalaGame.value?.Player2.holes;
  return reversedArray
    .map((hole: number, index: number) => {
      return {
        hole: hole,
        index: index,
      };
    })
    .reverse();
});

const winner = computed(() => {
  return mancalaGame.value?.Player1.bucket > mancalaGame.value?.Player2.bucket
    ? "Player 1 wins"
    : "Player 2 wins";
});
</script>

<template>
  <div>
    <h4 v-if="!mancalaGame">Select or Create Game</h4>
  </div>
  <div v-if="mancalaGame">
    <h2 v-if="!mancalaGame?.game_over">Now moves: {{ currentPlayer }}</h2>
    <h2 v-else>{{ winner }}</h2>

    <div class="mancala-board">
      <div class="bucket">
        <Bucket :stones="mancalaGame?.Player2.bucket" />
      </div>
      <div>
        <div class="pit-row">
          <Pit
            :disabled="currentPlayer === 'Player 1'"
            :player-id="mancalaGame?.Player2.ID"
            v-for="pit in opponentBoard"
            :stones="pit.hole"
            :index="pit.index"
            :key="pit.index + '-' + mancalaGame?.Player2.ID"
            :bottom-player="false"
          />
        </div>
        <br />
        <q-separator />
        <br />
        <div class="pit-row">
          <Pit
            :disabled="currentPlayer === 'Player 2'"
            :player-id="mancalaGame?.Player1.ID"
            v-for="(pit, index) in mancalaGame?.Player1.holes"
            :stones="pit"
            :index="index"
            :key="index + '-' + mancalaGame?.Player1.ID"
            :bottom-player="true"
          />
        </div>
      </div>
      <div class="bucket">
        <Bucket :stones="mancalaGame?.Player1.bucket" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.mancala-board {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  padding: 100px;
  border-radius: 16px;
  background-color: #b37a5f;
  box-shadow: 10px 10px 20px #cacaca, -10px -10px 20px #f6f6f6;
}

.pit-row {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 10em;
}
</style>
