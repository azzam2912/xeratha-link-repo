import axios from 'axios';

export const apiClient = axios.create({
    baseURL: 'https://xerathalinkrepo-azzamgamers12.b4a.run/',
    headers: {
        'Content-Type': 'application/json',
    },
});