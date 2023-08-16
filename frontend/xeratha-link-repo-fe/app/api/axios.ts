import axios from "axios";

const URL_API = `https://xerathalinkrepo-azzamgamers12.b4a.run/`;
const URL_LOCAL = `http://localhost:3006/`

export const apiClient = axios.create({
    baseURL: URL_LOCAL,
    headers: {
        'Content-Type': null,
    },
});