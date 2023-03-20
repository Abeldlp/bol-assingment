<script setup lang="ts">
import { StyleValue } from "vue";

const props = withDefaults(
  defineProps<{
    left: number;
    top: number;
  }>(),
  {
    left: 40,
    top: 90,
  }
);

const generateRandom = (num: number): number => {
  let result = Math.floor(Math.random() * num);
  result *= Math.round(Math.random()) ? 1 : -1;
  console.log(result);
  return result;
};

const customStyle: StyleValue = {
  width: "2em",
  height: "2em",
  transform:
    "translateX(" +
    generateRandom(props.left) +
    "px) translateY(" +
    generateRandom(props.top) +
    "px)",
  position: "absolute",
};
</script>

<template>
  <div :style="customStyle" class="stone"></div>
</template>

<style scoped>
.stone {
  border-radius: 50%;
  background: radial-gradient(circle at 10px 10px, #945149, #af6057);
  box-shadow: 4px 4px 5px #733f39, -1.5px -1.5px 2px #d57569;
  transition: all 0.2s ease-in-out;
}

.stone:hover {
  background: radial-gradient(circle at 10px 10px, #af6057, #945149);
  box-shadow: 4px 4px 5px #733f39, -1.5px -1.5px 2px #d57569;
}
</style>
