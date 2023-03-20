<script setup lang="ts">
import Stone from "./Stone.vue";
import { playRound } from "../data/mancalaGame";

const props = defineProps<{
  stones: number;
  index: number;
  playerId: number;
  bottomPlayer: boolean;
  disabled: boolean;
}>();

const updatePit = () => {
  if (props.disabled || props.stones === 0) return;
  playRound(props.index);
};
</script>

<template>
  <div class="pit" @click="updatePit">
    <Stone v-for="(_, index) in stones" :key="playerId + '-' + index" />
    <q-tooltip
      v-if="!bottomPlayer"
      anchor="top middle"
      self="bottom middle"
      :offset="[10, 10]"
      >{{ stones }}
    </q-tooltip>
    <q-tooltip v-if="bottomPlayer">{{ stones }} </q-tooltip>
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
  padding: 1rem;
  cursor: pointer;
  position: relative;
}
</style>
