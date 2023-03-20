import axios from "axios";

const API_URL = "http://localhost:5000/v1/api/mancala-game";
// http://localhost:5000/v1/api/mancala-game/2

export const getGame = async (gameId: string) => {
  const response = await axios.get(`${API_URL}/${gameId}`);
  return response.data;
};

export const createGame = async () => {
  const response = await axios.post(`${API_URL}`);
  return response.data;
};

export const makeMove = async (gameId: number, pitId: number) => {
  const response = await axios.put(`${API_URL}/${gameId}`, {
    "selected-pit": pitId,
  });
  return response;
};
