<script setup lang="ts">
import Stone from "./Stone.vue";
import { playRound } from "../data/mancalaGame";

const props = defineProps<{
  stones: number;
  index: number;
}>();

const updatePit = () => {
  playRound(props.index);
};
</script>

<template>
  <div class="pit" @click="updatePit">
    <Stone v-for="i in stones" :key="i" />
  </div>
</template>

<style scoped>
.pit {
  width: 10em;
  height: 10em;
  border-radius: 50%;
  border: 1px solid #000;
  margin: 0.5em;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  /* height: 100%; */
  padding: 1rem;
  cursor: pointer;
}
</style>
