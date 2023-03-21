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
  <div style="display: inline">
    <q-btn
      @click="createAndSetGame"
      no-caps
      color="white"
      text-color="black"
      label="Create New Game"
    />
    <div class="q-pa-md" style="max-width: 300px; margin: 0 auto">
      <div class="q-gutter-md">
        <q-select
          class="q-mt-md"
          @input="selectGame"
          v-model="selected"
          :options="allGames"
          label="Select Game"
          color="gray"
          text-color="black"
          dense
          outlined
        >
          <template #no-option>
            <q-item>
              <q-item-section class="text-grey">
                No results. Create a Game.
              </q-item-section>
            </q-item>
          </template>
        </q-select>
      </div>
    </div>
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
