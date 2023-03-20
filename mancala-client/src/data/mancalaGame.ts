import { computed, ref } from "vue";
import { makeMove } from "../api/MancalaApi";

export const mancalaGame = ref();

export const currentPlayer = computed(() => {
  if (mancalaGame.value) {
    return mancalaGame.value.turn % 2 === 0 ? "Player 2" : "Player 1";
  }
});

export const playRound = async (pitId: number) => {
  const res = await makeMove(mancalaGame.value.ID, pitId);
  mancalaGame.value = res.data;
};
