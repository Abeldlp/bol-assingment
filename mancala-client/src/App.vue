<script setup lang="ts">
import MancalaBoard from "./components/MancalaBoard.vue";
import { ref, onMounted, watch } from "vue";
import {
  getAllGames,
  allGames,
  setGame,
  createAndSetGame,
} from "./data/mancalaGame";

const selected = ref();

const selectGame = async () => {
  console.log(selected.value);
};

watch(selected, async () => {
  await setGame(selected.value.value);
});

onMounted(async () => {
  getAllGames();
});
</script>

<template>
  <div class="create-change-game-conatiner">
    <q-btn
      @click="createAndSetGame"
      no-caps
      color="white"
      text-color="black"
      label="Create Game"
    />
    <q-select
      @input="selectGame"
      @update:model-value=""
      v-model="selected"
      :options="allGames"
      width="100px"
      label="Select Game"
      outlined
      color="gray"
      text-color="black"
    >
    </q-select>
  </div>
  <MancalaBoard />
</template>

<style scoped>
.create-change-game-conatiner {
  display: flex;
  justify-content: space-around;
  align-items: center;
  margin: 1rem;
}
</style>
