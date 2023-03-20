import { computed, ref } from "vue";
import { update, getAll, getGameById, create } from "../api/MancalaApi";

export const allGames = ref([]);
export const mancalaGame = ref();

export const currentPlayer = computed(() => {
  if (mancalaGame.value) {
    return mancalaGame.value.turn % 2 === 0 ? "Player 1" : "Player 2";
  }
});

export const createAndSetGame = async () => {
  const res = await create();
  mancalaGame.value = res.data;
  await getAllGames();
};

export const setGame = async (gameId: string) => {
  const res = await getGameById(gameId);
  mancalaGame.value = res.data;
};

export const getAllGames = async () => {
  const res = await getAll();
  const games = formatSelecterOptions(res.data);
  allGames.value = games;
};

export const playRound = async (pitId: number) => {
  await update(mancalaGame.value.ID, pitId);
  const res = await getGameById(mancalaGame.value.ID);
  mancalaGame.value = res.data;
};

const formatSelecterOptions = (games: any) => {
  return games.map((game: any) => {
    return {
      label: "Game: " + game.ID,
      value: game.ID,
    };
  });
};
