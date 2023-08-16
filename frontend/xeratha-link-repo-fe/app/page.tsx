'use client'

import { useState } from 'react'
import { apiClient } from './api/axios'

export default function Home() {
  const [ links, setLinks ] = useState([]);
  const handleLoadLinks = async () => {
    const response: [] = await apiClient.get(
      "/links",
      { 
        withCredentials: false
      }
    )
    setLinks(response);
    console.log("links ", links);
  };
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <a onClick={handleLoadLinks}>Add Link</a>
        <div>
          <h1>Daftar Link</h1>
          <ul>
          
          </ul>
        </div>
      </div>
    </main>
  )
}
