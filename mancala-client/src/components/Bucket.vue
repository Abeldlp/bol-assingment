<script setup lang="ts">
import Stone from "./Stone.vue";
defineProps<{
  stones: number;
}>();
</script>
<template>
  <div class="bucket-container">
    <h5 class="bucket-count">{{ stones || "0" }}</h5>
    <div class="bucket">
      <Stone v-for="(_, index) in stones" :key="index" />
    </div>
  </div>
</template>

<style scoped>
.bucket-container {
  position: relative;
}
.bucket {
  width: 10em;
  height: 20em;
  border-radius: 10px;
  margin: 0.5em;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  padding: 1rem;
  position: relative;
  background: #b37a5f;
  box-shadow: inset 10px 10px 20px #7d5543, inset -10px -10px 20px #e99f7c;
}

.bucket-count {
  position: absolute;
  font-family: "Economica", sans-serif;
  font-size: 30px;
  top: -35%;
  left: 25%;
  right: 25%;
}
</style>
