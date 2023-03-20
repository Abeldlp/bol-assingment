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
    <h5 v-if="!bottomPlayer" class="stone-top">{{ stones }}</h5>
    <Stone
      v-for="(_, index) in stones"
      :key="playerId + '-' + index"
      :left="35"
      :top="35"
    />
    <h5 v-if="bottomPlayer" class="stone-bottom">{{ stones }}</h5>
  </div>
</template>

<style scoped>
.pit {
  width: 10em;
  height: 10em;
  border-radius: 50%;
  margin: 0.5em;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  position: relative;
  background: #b37a5f;
  box-shadow: inset 8px 8px 16px #7d5543, inset -8px -8px 16px #e99f7c;
  user-select: none;
  font-family: "Economica", sans-serif;
}

.stone-top {
  position: absolute;
  top: -55%;
  left: 25%;
  right: 25%;
}

.stone-bottom {
  position: absolute;
  bottom: -55%;
  left: 25%;
  right: 25%;
}
</style>
